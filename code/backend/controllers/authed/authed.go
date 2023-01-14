package authed

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type Controller struct {
	coredb      store.CoreHub
	signer      service.Signer
	sessionNode *snowflake.Node
	deviceNode  *snowflake.Node
}

func New(coredb store.CoreHub, signer service.Signer, seq xplane.IDService) *Controller {
	return &Controller{
		coredb:      coredb,
		signer:      signer,
		sessionNode: seq.SessionNode(),
		deviceNode:  seq.DeviceNode(),
	}
}

func (c *Controller) AuthListMethods(sitetoken, ugroup string) (*ListAuthResponse, error) {
	site, err := c.signer.ParseSite(sitetoken)
	if err != nil {
		return nil, err
	}

	if ugroup == "" && site.PinnedUserGroup == "" {
		return nil, easyerr.Error("user group not found")
	}

	if ugroup != "" && site.PinnedUserGroup != "" && ugroup != site.PinnedUserGroup {
		return nil, easyerr.Error("user group not allowed")
	}

	if ugroup == "" {
		ugroup = site.PinnedUserGroup
	}

	if ugroup == xtypes.UserGroupSuperAdmin {
		return &ListAuthResponse{
			PasswordAuth:   true,
			OpenSignUp:     false,
			AltAuthMethods: nil,
		}, nil
	}

	auths, err := c.coredb.ListUserGroupAuth(site.TenentId, ugroup)
	if err != nil {
		return nil, err
	}

	alts := make([]AltAuthMethod, len(auths))

	for _, uga := range auths {
		alts = append(alts, AltAuthMethod{
			Id:       uga.Id,
			Name:     uga.Name,
			Type:     uga.Type,
			Provider: uga.Provider,
		})
	}

	resp := ListAuthResponse{
		PasswordAuth:   true,
		OpenSignUp:     true,
		AltAuthMethods: alts,
	}

	return &resp, nil
}

/*
	// fixme  => device scope properly
		scope.Derive(sclaim.Scopes, strings.Split(ugroup.Scopes, ",")),
		ugroup.Scopes


*/

func (c *Controller) AuthFinish(opts AuthFinishRequest, deviceName, addr string) (*AuthFinishResponse, error) {
	sclaim, err := c.signer.ParseSite(opts.SiteToken)
	if err != nil {
		return nil, err
	}

	paclaim, err := c.signer.ParsePreAuthed(sclaim.TenentId, opts.PreAuthedToken)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	ugroup, err := c.coredb.GetUserGroup(sclaim.TenentId, paclaim.UserGroup)
	if err != nil {
		return nil, err
	}

	did := c.deviceNode.Generate().Int64()

	device := &entities.UserDevice{
		Id:         did,
		Name:       deviceName,
		UserId:     paclaim.UserID,
		DeviceType: "login",
		LastData: entities.JsonStrMap{
			"address": addr,
		},
		APNToken:    "",
		Scopes:      ugroup.Scopes,
		ExtraMeta:   entities.JsonStrMap{},
		TenantID:    sclaim.TenentId,
		PairOptions: entities.JsonStrMap{},
		ExpiresOn:   time.Now().Add(time.Hour * 24 * 60),
	}

	err = c.coredb.AddUserDevice(sclaim.TenentId, paclaim.UserID, device)
	if err != nil {
		return nil, err
	}

	utok, err := c.signer.SignUser(sclaim.TenentId, device.Derive(paclaim.UserGroup))
	if err != nil {
		return nil, err
	}

	return &AuthFinishResponse{
		UserToken: utok,
	}, nil
}

func (c *Controller) AuthGenerate(opts AuthGenerateRequest) (*AuthGenerateResponse, error) {
	return c.authGenerate(opts)
}

func (c *Controller) AuthNextFirst(opts AuthNextFirstRequest) (*AuthNextFirstResponse, error) {
	return c.authNextFirst(opts)
}

func (c *Controller) AuthNextSecond(opts AuthNextSecondRequest) (*AuthNextSecondResponse, error) {
	return c.authNextSecond(opts)
}

func (c *Controller) AuthSubmit(opts AuthSubmitRequest) (*AuthSubmitResponse, error) {
	return c.authSubmit(opts)
}

func (c *Controller) LoginNext(opts LoginNextRequest) (*LoginNextResponse, error) {
	return c.loginNext(opts)
}

func (c *Controller) LoginSubmit(opts LoginSubmitRequest) (*LoginSubmitResponse, error) {
	return c.loginSubmit(opts)
}
