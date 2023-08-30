package authed

import (
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) loginNext(tenantId string, opts LoginNextRequest) (*LoginNextResponse, error) {
	const msg = "User or Password incorrect"

	var user *entities.User
	var err error

	if strings.Contains(opts.UserIdent, "@") {
		user, err = c.coredb.GetUserByEmail(tenantId, opts.UserIdent)
	} else {
		user, err = c.coredb.GetUserByID(tenantId, opts.UserIdent)
	}
	if err != nil {
		return &LoginNextResponse{
			Message: msg,
		}, nil
	}

	data, err := c.coredb.GetUserData(tenantId, user.UserId)
	if err != nil {
		return nil, err
	}

	ugroup, err := c.coredb.GetUserGroup(tenantId, user.GroupID)
	if err != nil {
		return nil, err
	}

	if !ugroup.HasFeature("pass_auth") && ugroup.Slug != "super_admin" {
		return &LoginNextResponse{
			Message: "auth method not allowed",
		}, nil
	}

	if user.Password != opts.Password {
		return &LoginNextResponse{
			Message: msg,
		}, nil
	}

	tok, err := c.signer.SignAutheNext(tenantId, &claim.AuthNext{
		UserId:      user.UserId,
		UserGroup:   user.GroupID,
		UserEmail:   user.Email,
		EmailVerify: data.PendingEmailVerify,
		PassChange:  data.PendingPassChange,
		DeviceId:    "",
	})

	if err != nil {
		return nil, err
	}

	return &LoginNextResponse{
		Message:     "",
		Ok:          true,
		NextToken:   tok,
		EmailVerify: data.PendingEmailVerify,
	}, nil

}

func (c *Controller) loginSubmit(tenantId string, opts LoginSubmitRequest) (*LoginSubmitResponse, error) {

	nclaim, err := c.signer.ParseAutheNext(tenantId, opts.NextToken)
	if err != nil {
		return nil, err
	}

	udata, err := c.coredb.GetUserData(tenantId, nclaim.UserId)
	if err != nil {
		return nil, err
	}

	if nclaim.EmailVerify {
		if udata.PendingEmailVerify {
			return nil, easyerr.Error("Email not verified")
		}
	}

	if nclaim.PassChange {
		if !udata.PendingPassChange {
			return nil, easyerr.Error("Password not changed")
		}
	}

	patok, err := c.signer.SignPreAuthed(tenantId, &claim.PreAuthed{
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

	return &LoginSubmitResponse{
		SubmitResponse: SubmitResponse{
			Ok:             true,
			Message:        "",
			PreAuthedToken: patok,
		},
	}, nil
}
