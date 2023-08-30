package authed

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
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

func (c *Controller) AuthListMethods(tenantId, ugroup string) (*ListAuthResponse, error) {

	if ugroup == "" {
		ugroup = xtypes.UserGroupSuperAdmin
	}

	if ugroup == xtypes.UserGroupSuperAdmin {
		return &ListAuthResponse{
			PasswordAuth:   true,
			OpenSignUp:     false,
			AltAuthMethods: nil,
		}, nil
	}

	auths, err := c.coredb.ListUserGroupAuth(tenantId, ugroup)
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

func (c *Controller) AuthFinish(opts AuthFinishRequest, tenantId, deviceName, addr string) (*AuthFinishResponse, error) {

	paclaim, err := c.signer.ParsePreAuthed(tenantId, opts.PreAuthedToken)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	ugroup, err := c.coredb.GetUserGroup(tenantId, paclaim.UserGroup)
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
		TenantID:    tenantId,
		PairOptions: entities.JsonStrMap{},
		ExpiresOn: dbutils.Time{
			Inner: time.Now().Add(time.Hour * 24 * 60),
		},
	}

	err = c.coredb.AddUserDevice(tenantId, paclaim.UserID, device)
	if err != nil {
		return nil, err
	}

	utok, err := c.signer.SignUser(tenantId, device.Derive(paclaim.UserGroup))
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

func (c *Controller) LoginNext(tenantId string, opts LoginNextRequest) (*LoginNextResponse, error) {
	return c.loginNext(tenantId, opts)
}

func (c *Controller) LoginSubmit(tenantId string, opts LoginSubmitRequest) (*LoginSubmitResponse, error) {
	return c.loginSubmit(tenantId, opts)
}
