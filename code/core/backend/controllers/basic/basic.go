package basic

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/scopes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Controller struct {
	coredb  store.CoreHub
	cabinet store.CabinetHub
	dynHub  store.DataHub
	pacman  repox.Hub
	signer  service.Signer
}

func New(coredb store.CoreHub, cabinet store.CabinetHub, dynHub store.DataHub, pacman repox.Hub) *Controller {
	ctrl := &Controller{
		coredb:  coredb,
		cabinet: cabinet,
		dynHub:  dynHub,
		pacman:  pacman,
	}

	return ctrl
}

func (c *Controller) ListRepoSources(uclaim *claim.Session) (map[int64]string, error) {
	return c.pacman.RepoSources(uclaim.TenentId)
}

func (c *Controller) ListCabinetSources(uclaim *claim.Session) ([]string, error) {
	return c.cabinet.ListSources(uclaim.TenentId)
}

func (c *Controller) ListDyndbSources(uclaim *claim.Session) ([]string, error) {
	return c.dynHub.ListSources((uclaim.TenentId))
}

func (c *Controller) JoinNotification() error {
	return nil
}

func (c *Controller) Self(uclaim *claim.Session) (*entities.User, error) {
	return c.coredb.GetUserByID(uclaim.TenentId, uclaim.UserID)
}

func (c *Controller) SelfUpdate(uclaim *claim.Session) error {
	// fixme =>
	return nil
}

func (c *Controller) GetSelfInfo(uclaim *claim.Session) (*entities.SelfLoad, error) {
	usr, err := c.coredb.GetUserByID(uclaim.TenentId, uclaim.UserID)
	if err != nil {
		pp.Println("uclaim", uclaim)
		pp.Println("@USER BY ID", err)
		return nil, err
	}

	tenant, err := c.coredb.GetTenant(uclaim.TenentId)
	if err != nil {
		pp.Println("@TENANT", err)
		return nil, err
	}

	ugroup, err := c.coredb.GetUserGroup(uclaim.TenentId, uclaim.UserGroup)
	if err != nil {
		pp.Println("@USER_GROUP ID", err)
		return nil, err
	}

	var scs []string
	if uclaim.UserGroup == xtypes.UserGroupSuperAdmin {
		scs = scopes.SuperScopes
	}

	apps, err := c.coredb.ListTargetAppByUgroup(uclaim.TenentId, uclaim.UserGroup)
	papps := make([]entities.PlugApp, 0, len(apps))
	if err == nil {
		for _, app := range apps {
			papps = append(papps, entities.PlugApp{
				Id:      app.Id,
				Name:    app.Name,
				PlugId:  app.PlugId,
				AgentId: app.AgentId,
				Icon:    app.Icon,
			})
		}
	}

	return &entities.SelfLoad{
		UserInfo: entities.UserInfo{
			UserId:    uclaim.UserID,
			FullName:  usr.FullName,
			PublicKey: usr.PublicKey,
			Bio:       usr.Bio,
			Email:     usr.Email,
			GroupId:   uclaim.UserGroup,
			GroupName: ugroup.Name},
		TenantName: tenant.Name,
		TenantId:   uclaim.TenentId,
		Scopes:     scs,
		PlugApps:   papps,
	}, nil

}

func (c *Controller) GetChangeEmail(uclaim *claim.Session) error {

	return nil
}

func (c *Controller) GetChangeAuth(uclaim *claim.Session) error {
	return nil
}

func (c *Controller) ListMessages(uclaim *claim.Session, opts *entities.UserMessageReq) ([]*entities.UserMessage, error) {
	return c.coredb.ListUserMessages(uclaim.TenentId, opts)
}

func (c *Controller) ModifyMessages(uclaim *claim.Session, opts *entities.ModifyMessages) error {
	switch opts.Operation {
	case "delete":
		return c.coredb.DeleteUserMessages(uclaim.TenentId, uclaim.UserID, opts.Ids)
	case "read":
		return c.coredb.ReadUserMessages(uclaim.TenentId, uclaim.UserID, opts.Ids)
	default:
		panic("not impl")
	}
}
