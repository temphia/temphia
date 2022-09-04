package corehub

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *CoreHub) AddUserGroup(ug *entities.UserGroup) error {
	return c.coredb.AddUserGroup(ug)
}

func (c *CoreHub) GetUserGroup(tenantId string, slug string) (*entities.UserGroup, error) {
	return c.coredb.GetUserGroup(tenantId, slug)
}

func (c *CoreHub) ListUserGroups(tenantId string) ([]*entities.UserGroup, error) {
	return c.coredb.ListUserGroups(tenantId)
}

func (c *CoreHub) UpdateUserGroup(tenantId, slug string, data map[string]any) error {
	return c.coredb.UpdateUserGroup(tenantId, slug, data)
}

func (c *CoreHub) RemoveUserGroup(tenantId string, ugslug string) error {
	return c.coredb.RemoveUserGroup(tenantId, ugslug)
}

func (c *CoreHub) AddUser(user *entities.User, data *entities.UserData) error {
	return c.coredb.AddUser(user, data)
}

func (c *CoreHub) UpdateUser(tenantId, user string, data map[string]any) error {
	return c.coredb.UpdateUser(tenantId, user, data)
}

func (c *CoreHub) RemoveUser(tenantId string, username string) error {
	return c.coredb.RemoveUser(tenantId, username)
}

func (c *CoreHub) GetUserByID(tenantId string, username string) (*entities.User, error) {
	return c.coredb.GetUserByID(tenantId, username)
}

func (c *CoreHub) GetUserByEmail(tenantId string, email string) (*entities.User, error) {
	return c.coredb.GetUserByEmail(tenantId, email)
}

func (c *CoreHub) ListUsers(tenantId string) ([]*entities.User, error) {
	users, err := c.coredb.ListUsers(tenantId)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.Password = ""
	}

	return users, nil
}

func (c *CoreHub) ListUsersByGroup(tenantId, group string) ([]*entities.User, error) {
	users, err := c.coredb.ListUsersByGroup(tenantId, group)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.Password = ""
	}

	return users, nil
}

func (c *CoreHub) ListUsersMulti(tenantId string, ids ...string) ([]*entities.User, error) {
	users, err := c.coredb.ListUsersMulti(tenantId, ids...)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.Password = ""
	}
	return users, nil
}

// self user

func (c *CoreHub) AddUserMessage(msg *entities.UserMessage) (int64, error) {
	id, err := c.coredb.AddUserMessage(msg)
	if err != nil {
		pp.Println("Log this", err)
		return 0, err
	}
	msg.Id = int(id)

	err = c.sockdhub.NotifyUser(msg)
	if err != nil {
		// fixme => log this
		pp.Println("sending notification err", err)
		return id, nil
	}

	return id, nil
}

func (c *CoreHub) UserMessageSetRead(tenantId, user string, id int64) error {
	return c.coredb.UserMessageSetRead(tenantId, user, id)
}

func (c *CoreHub) RemoveUserMessage(tenantId string, username string, id int64) error {
	return c.coredb.RemoveUserMessage(tenantId, username, id)
}

func (c *CoreHub) ListUserMessages(tenantId string, data *entities.UserMessageReq) ([]*entities.UserMessage, error) {
	return c.coredb.ListUserMessages(tenantId, data)
}

func (c *CoreHub) ReadUserMessages(tenantId, userId string, id []int64) error {
	return c.coredb.ReadUserMessages(tenantId, userId, id)
}

func (c *CoreHub) DeleteUserMessages(tenantId, userId string, id []int64) error {
	return c.coredb.DeleteUserMessages(tenantId, userId, id)
}

func (c *CoreHub) GetUserData(tenantId string, slug string) (*entities.UserData, error) {
	return c.coredb.GetUserData(tenantId, slug)
}

func (c *CoreHub) UpdateUserData(tenantId, slug string, data map[string]any) error {
	return c.coredb.UpdateUserData(tenantId, slug, data)
}
