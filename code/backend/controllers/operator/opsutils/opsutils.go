package opsutils

import (
	"time"

	"github.com/temphia/temphia/code/backend/controllers/operator/opmodels"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

const (
	DefaultUserName = "Super User"
	DefaultUser     = "superuser"

	DefaultGroupName = "Super Admin"
	DefaultGroup     = "super_admin"
)

func AddTenant(app xtypes.App, data *opmodels.NewTenant) error {

	if app.SingleTenant() {
		if app.StaticTenants()[0] != data.Slug {
			return easyerr.Error("Server is in single tenant mode, you cannot add this tenant")
		}
	}

	coredb := app.GetDeps().CoreHub().(store.CoreDB)

	err := coredb.AddTenant(&entities.Tenant{
		Name: data.Name,
		Slug: data.Slug,
	})

	if err != nil {
		return err
	}

	err = coredb.AddUserGroup(&entities.UserGroup{
		Name:     DefaultGroupName,
		Slug:     DefaultGroup,
		TenantID: data.Slug,
	})

	if err != nil {
		return err
	}

	return coredb.AddUser(&entities.User{
		UserId:    DefaultUser,
		FullName:  DefaultUserName,
		Email:     data.SuperEmail,
		GroupID:   DefaultGroup,
		Password:  data.SuperPassword,
		TenantID:  data.Slug,
		PublicKey: "",
		CreatedAt: time.Now(),
		Active:    true,
	}, &entities.UserData{
		UserId:             DefaultUser,
		MFAEnabled:         false,
		MFAType:            "",
		MFAData:            "",
		PendingPassChange:  false,
		PendingEmailVerify: false,
		ExtraMeta:          nil,
		TenantID:           data.Slug})

}
