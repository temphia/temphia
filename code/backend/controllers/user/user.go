package user

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Controller struct {
	corehub store.CoreHub
}

func New(corehub store.CoreHub) *Controller {
	return &Controller{
		corehub: corehub,
	}
}

func (c *Controller) Message(uclaim *claim.Session, userId, message string) (int64, error) {

	return c.corehub.AddUserMessage(&entities.UserMessage{
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

func (c *Controller) Get(uclaim *claim.Session, userId string) (*entities.UserInfo, error) {
	usr, err := c.corehub.GetUserByID(uclaim.TenentId, userId)
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

func (c *Controller) List(uclaim *claim.Session, opts any) ([]entities.UserInfo, error) {

	return nil, nil
}
