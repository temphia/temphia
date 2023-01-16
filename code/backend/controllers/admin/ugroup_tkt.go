package admin

import (
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) UgroupGetUserByID(uclaim *claim.UserMgmtTkt, username string) (*entities.User, error) {
	return c.coredb.GetUserByID(uclaim.TenantId, username)
}

func (c *Controller) UgroupListUsersByGroup(uclaim *claim.UserMgmtTkt) ([]*entities.User, error) {
	return c.coredb.ListUsersByGroup(uclaim.TenantId, uclaim.Group)
}

func (c *Controller) UgroupAddUser(uclaim *claim.UserMgmtTkt, data *entities.User) error {

	data.TenantID = uclaim.TenantId
	data.CreatedAt = time.Now()

	return c.coredb.AddUser(data, &entities.UserData{
		UserId:             data.UserId,
		MFAEnabled:         false,
		MFAType:            "",
		MFAData:            "",
		PendingPassChange:  true,
		PendingEmailVerify: false,
		ExtraMeta:          nil,
		TenantID:           uclaim.TenantId,
	})
}

func (c *Controller) UgroupUpdateUser(uclaim *claim.UserMgmtTkt, userid string, data map[string]any) error {
	return c.coredb.UpdateUser(uclaim.TenantId, userid, data)
}

func (c *Controller) UgroupDeleteUser(uclaim *claim.UserMgmtTkt, userId string) error {
	return c.coredb.RemoveUser(uclaim.TenantId, userId)
}
