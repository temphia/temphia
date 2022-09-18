package basic

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
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

func (c *Controller) ListRepoSources(uclaim *claim.Session) (any, error) {
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

func (c *Controller) MessageUser(uclaim *claim.Session, userId, message string) (int64, error) {

	return c.coredb.AddUserMessage(&entities.UserMessage{
		Id:           0,
		Title:        "message",
		Read:         false,
		Type:         "user_message",
		Contents:     message,
		UserId:       userId,
		FromUser:     uclaim.UserID,
		FromPlug:     "",
		FromAgent:    "",
		PlugCallback: "",
		WarnLevel:    1,
		Encrypted:    false,
		CreatedAt:    nil,
		TenantId:     uclaim.TenentId,
	})
}

func (c *Controller) GetUserInfo(uclaim *claim.Session, userId string) (*entities.UserInfo, error) {
	usr, err := c.coredb.GetUserByID(uclaim.TenentId, userId)
	if err != nil {
		return nil, err
	}

	fuser := &entities.UserInfo{
		UserId:    uclaim.UserID,
		FullName:  usr.FullName,
		Bio:       "",
		PublicKey: "",
		Email:     "",
		GroupId:   "",
	}

	return fuser, nil
}

func (c *Controller) GetSelfInfo(uclaim *claim.Session) (*entities.UserInfo, error) {
	usr, err := c.coredb.GetUserByID(uclaim.TenentId, uclaim.UserID)
	if err != nil {
		return nil, err
	}

	tenant, err := c.coredb.GetTenant(uclaim.TenentId)
	if err != nil {
		return nil, err
	}

	ugroup, err := c.coredb.GetUserGroup(uclaim.TenentId, uclaim.UserGroup)
	if err != nil {
		return nil, err
	}

	fuser := &entities.UserInfo{
		UserId:     uclaim.UserID,
		FullName:   usr.FullName,
		PublicKey:  usr.PublicKey,
		Bio:        usr.Bio,
		Email:      usr.Email,
		GroupId:    uclaim.UserGroup,
		GroupName:  ugroup.Name,
		TenantName: tenant.Name,
		TenantId:   uclaim.TenentId,
	}

	return fuser, nil
}

func (c *Controller) GetChangeEmail(uclaim *claim.Session) error {

	return nil
}

func (c *Controller) GetChangeAuth(uclaim *claim.Session) error {
	return nil
}

func (c *Controller) ListMessages(uclaim *claim.Session, opts *entities.UserMessageReq) ([]*entities.UserMessage, error) {
	opts.UserId = uclaim.UserID
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
