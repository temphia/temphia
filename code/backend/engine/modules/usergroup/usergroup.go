package usergroup

import (
	"fmt"
	"time"

	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type UserGroupModule struct {
	coreHub  store.CoreHub
	group    string
	tenantId string
	bindings bindx.Bindings

	modipc *modipc.ModIPC
}

func (u *UserGroupModule) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {
	return u.modipc.Handle(method, args)
}

func (u *UserGroupModule) Close() error { return nil }

type addUserOpts struct {
	User *entities.User     `json:"user,omitempty"`
	Data *entities.UserData `json:"data,omitempty"`
}

func (u *UserGroupModule) AddUser(opts addUserOpts) error {
	opts.User.GroupID = u.group
	opts.User.TenantID = u.tenantId
	opts.Data.TenantID = u.tenantId

	return u.coreHub.AddUser(opts.User, opts.Data)
}

func (u *UserGroupModule) ListUsersByGroup(tenantId, group string) ([]entities.UserInfo, error) {
	users, err := u.coreHub.ListUsersByGroup(u.tenantId, u.group)

	if err != nil {
		return nil, err
	}

	fusers := make([]entities.UserInfo, 0, len(users))

	for _, user := range users {
		fusers = append(fusers, entities.UserInfo{
			UserId:    user.UserId,
			FullName:  user.FullName,
			Bio:       user.Bio,
			PublicKey: user.PublicKey,
			Email:     user.Email,
			GroupId:   user.GroupID,
		})
	}

	return fusers, nil
}

type messageUserOpts struct {
	Title    string `json:"title,omitempty"`
	Contents string `json:"contents,omitempty"`
	UserId   string `json:"user_id,omitempty"`
	FromUser string `json:"from_user,omitempty"`
}

func (u *UserGroupModule) MessageUserByUser(opts *messageUserOpts) error {
	// fixme => check group
	// u.coreHub.ListUsersMulti(u.tenantId, opts.UserId, opts.FromUser)

	_, err := u.coreHub.AddUserMessage(&entities.UserMessage{
		Title:        fmt.Sprintf("Plug message: %s", opts.Title),
		Type:         "message",
		Contents:     opts.Contents,
		UserId:       opts.UserId,
		FromUser:     opts.FromUser,
		FromPlug:     "",
		FromAgent:    "",
		PlugCallback: "",
		WarnLevel:    1,
		Encrypted:    false,
		CreatedAt: &dbutils.Time{
			Inner: time.Now(),
		},
		TenantId: u.tenantId,
	})

	return err
}

type getUserOpts struct {
	UserId string `json:"title,omitempty"`
}

func (u *UserGroupModule) GetUser(opts *getUserOpts) (*entities.UserInfo, error) {
	user, err := u.coreHub.GetUserByID(u.tenantId, opts.UserId)
	if err != nil {
		return nil, err
	}

	if user.GroupID != u.group {
		return nil, easyerr.NotAuthorized()
	}

	usr := &entities.UserInfo{
		UserId:    user.UserId,
		FullName:  user.FullName,
		Bio:       user.Bio,
		PublicKey: user.PublicKey,
		Email:     user.Email,
		GroupId:   user.GroupID,
	}

	return usr, nil
}
