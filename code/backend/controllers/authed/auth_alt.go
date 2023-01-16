package authed

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/thoas/go-funk"
)

func (c *Controller) authGenerate(opts AuthGenerateRequest) (*AuthGenerateResponse, error) {

	sclaim, err := c.signer.ParseSite(opts.SiteToken)
	if err != nil {
		return nil, err
	}

	auth, err := c.coredb.GetUserGroupAuth(sclaim.TenantId, opts.UserGroup, opts.Id)
	if err != nil {
		return nil, err
	}

	switch auth.Type {
	case "oauth":

		otok, err := c.signer.SignOauthState(sclaim.TenantId, &claim.OauthState{
			TenantId:  sclaim.TenantId,
			AuthId:    auth.Id,
			UserGroup: opts.UserGroup,
			DeviceId:  "",
		})
		if err != nil {
			return nil, err
		}

		return &AuthGenerateResponse{
			StateToken: otok,
			AuthURL:    auth.AuthURL(),
			Scopes:     strings.Split(auth.Scopes, ","),
			ClientId:   auth.ClientId(),
		}, nil

	default:
		panic("notimplemented")
	}

}

func (c *Controller) authNextFirst(opts AuthNextFirstRequest) (*AuthNextFirstResponse, error) {
	sclaim, err := c.signer.ParseSite(opts.SiteToken)
	if err != nil {
		return nil, err
	}

	auth, err := c.coredb.GetUserGroupAuth(sclaim.TenantId, opts.UserGroup, opts.Id)
	if err != nil {
		return nil, err
	}

	switch auth.Type {
	case "oauth":
		return c.oauthInNext(sclaim, opts, auth)
	default:
		panic("notimplemented")
	}

}
func (c *Controller) authNextSecond(opts AuthNextSecondRequest) (*AuthNextSecondResponse, error) {

	sclaim, err := c.signer.ParseSite(opts.SiteToken)
	if err != nil {
		return nil, err
	}

	fnclaim, err := c.signer.ParseAutheFirst(sclaim.TenantId, opts.FirstToken)
	if err != nil {
		return nil, err
	}

	if fnclaim.NewUser {
		fnclaim.UserID = opts.SignUpdata.UserId
		err := c.coredb.AddUser(&entities.User{
			UserId:    opts.SignUpdata.UserId,
			FullName:  opts.SignUpdata.FullName,
			Email:     fnclaim.UserEmail,
			Bio:       opts.SignUpdata.Bio,
			GroupID:   fnclaim.UserGroup,
			Password:  "",
			TenantID:  sclaim.TenantId,
			PublicKey: "",
			CreatedAt: time.Now(),
			Active:    true,
		}, &entities.UserData{
			UserId:             opts.SignUpdata.UserId,
			MFAEnabled:         false,
			MFAType:            "",
			MFAData:            "",
			PendingPassChange:  false,
			PendingEmailVerify: false,
			ExtraMeta:          nil,
			TenantID:           sclaim.TenantId,
		})
		if err != nil {
			return nil, err
		}
	}

	ntok, err := c.signer.SignAutheNext(sclaim.TenantId, &claim.AuthNext{
		UserId:      fnclaim.UserID,
		UserGroup:   fnclaim.UserGroup,
		UserEmail:   fnclaim.UserEmail,
		DeviceId:    "",
		EmailVerify: false,
		PassChange:  false,
	})
	if err != nil {
		return nil, err
	}

	return &AuthNextSecondResponse{
		Ok:          true,
		Message:     "",
		NextToken:   ntok,
		UserId:      fnclaim.UserID,
		EmailVerify: false,
	}, nil
}

func (c *Controller) authSubmit(opts AuthSubmitRequest) (*AuthSubmitResponse, error) {

	sclaim, err := c.signer.ParseSite(opts.SiteToken)
	if err != nil {
		return nil, err
	}

	nclaim, err := c.signer.ParseAutheNext(sclaim.TenantId, opts.NextToken)
	if err != nil {
		return nil, err
	}

	data, err := c.coredb.GetUserData(sclaim.TenantId, nclaim.UserId)
	if err != nil {
		return nil, err
	}

	if nclaim.EmailVerify {
		if data.PendingEmailVerify {
			return nil, easyerr.Error("Email Not verified")
		}
	}

	if nclaim.PassChange {
		if data.PendingEmailVerify {
			return nil, easyerr.Error("Password not changed")
		}
	}

	patok, err := c.signer.SignPreAuthed(sclaim.TenantId, &claim.PreAuthed{
		UserID:     nclaim.UserId,
		UserGroup:  nclaim.UserGroup,
		UserEmail:  nclaim.UserEmail,
		AuthId:     0,
		NeedsProof: false,
		DeviceId:   "",
	})
	if err != nil {
		return nil, err
	}

	return &AuthSubmitResponse{
		SubmitResponse{Message: "",
			Ok:             false,
			PreAuthedToken: patok,
			HasExecHook:    false,
			HookPlugId:     "",
			HookAgentId:    "",
			HookExecToken:  ""},
	}, nil

}

// oauth

func (c *Controller) oauthInNext(sclaim *claim.Site, opts AuthNextFirstRequest, auth *entities.UserGroupAuth) (*AuthNextFirstResponse, error) {

	state, err := c.signer.ParseOauthState(sclaim.TenantId, opts.AuthState)
	if err != nil {
		return nil, err
	}

	exchanger := &OauthExchanger{
		auth:     auth,
		authCode: opts.AuthCode,
		state:    state,
	}

	data, err := exchanger.Exchange()
	if err != nil {
		return nil, err
	}

	usr, err := c.coredb.GetUserByEmail(sclaim.TenantId, data.email)
	if err != nil {
		// if !auth.NewUserIfNotExist {
		// 	return nil, err
		// }

		errstr := err.Error()
		if !strings.Contains(errstr, "no more rows") && !strings.Contains(errstr, "not found") {
			return nil, err
		}

		tok, err := c.signer.SignAutheFirst(auth.TenantId, &claim.AuthFirst{
			UserGroup: state.UserGroup,
			NewUser:   true,
			UserID:    "",
			UserEmail: data.email,
			AuthId:    auth.Id,
		})

		if err != nil {
			return nil, err
		}

		idhints, err := c.deriveUserIds(sclaim.TenantId, data.email, data.userName)
		if err != nil {
			return nil, err
		}

		return &AuthNextFirstResponse{
			Message:     "",
			Ok:          true,
			NewUser:     true,
			FirstToken:  tok,
			Email:       data.email,
			UserIdHints: idhints,
		}, nil

	}

	tok, err := c.signer.SignAutheFirst(auth.TenantId, &claim.AuthFirst{
		UserGroup: state.UserGroup,
		NewUser:   false,
		UserID:    usr.UserId,
		UserEmail: data.email,
		AuthId:    auth.Id,
	})

	if err != nil {
		return nil, err
	}

	return &AuthNextFirstResponse{
		Message:     "",
		Ok:          false,
		FirstToken:  tok,
		NewUser:     false,
		Email:       data.email,
		UserIdHints: nil,
	}, nil

}

func (c *Controller) deriveUserIds(tenantId, email, username string) ([]string, error) {
	ids := []string{strings.Split(email, "@")[0]}
	if username != "" {
		ids = append(ids, strings.Split(username, "")...)
	}

	dupIds := make([]string, len(ids)*5)

	idx := 0
	for range ids {
		dupIds[idx+0] = fmt.Sprintf("%s%d", ids[idx], rand.Int31n(100))
		dupIds[idx+1] = fmt.Sprintf("%s%d", ids[idx], 100+rand.Int31n(100))
		dupIds[idx+2] = fmt.Sprintf("%s%d", ids[idx], 200+rand.Int31n(100))
		dupIds[idx+3] = fmt.Sprintf("%s%d", ids[idx], 300+rand.Int31n(100))
		dupIds[idx+4] = fmt.Sprintf("%s%d", ids[idx], 400+rand.Int31n(100))
		idx = idx + 5
	}

	usrs, err := c.coredb.ListUsersMulti(tenantId, dupIds...)
	if err != nil {
		return nil, nil
	}

	respIds := make([]string, 0, len(dupIds)-len(usrs))
	ruids := make([]string, len(usrs))

	for _, usr := range usrs {
		ruids = append(ruids, usr.UserId)
	}

	for _, id := range dupIds {
		if funk.ContainsString(ruids, id) {
			continue
		}
		respIds = append(respIds, id)
	}

	return respIds, nil
}
