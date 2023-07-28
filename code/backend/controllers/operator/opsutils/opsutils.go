package opsutils

import (
	"time"

	"github.com/temphia/temphia/code/backend/controllers/operator/opmodels"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func AddTenant(coredb store.CoreDB, data *opmodels.NewTenant) error {

	return AddTenantWithUser(coredb, &TenantWithUserOptions{
		Name:          data.Name,
		Slug:          data.Slug,
		SuperPassword: data.SuperPassword,
		SuperEmail:    data.SuperEmail,
		SuperUserName: opmodels.DefaultUserName,
		SuperUser:     opmodels.DefaultUser,
		SuperGroup:    opmodels.DefaultGroup,
	})

}

type TenantWithUserOptions struct {
	Name           string
	Slug           string
	SuperPassword  string
	SuperEmail     string
	SuperUserName  string
	SuperUser      string
	SuperGroup     string
	SuperGroupName string
}

func AddTenantWithUser(coredb store.CoreDB, opts *TenantWithUserOptions) error {

	err := coredb.AddTenant(&entities.Tenant{
		Name: opts.Name,
		Slug: opts.Slug,
	})

	if err != nil {
		return err
	}

	err = coredb.AddUserGroup(&entities.UserGroup{
		Name:     opts.SuperGroupName,
		Slug:     opts.SuperGroup,
		TenantID: opts.Slug,
	})

	if err != nil {
		return err
	}

	return coredb.AddUser(&entities.User{
		UserId:    opts.SuperUser,
		FullName:  opts.SuperUserName,
		Email:     opts.SuperEmail,
		GroupID:   opts.SuperGroup,
		Password:  opts.SuperPassword,
		TenantID:  opts.Slug,
		PublicKey: "",
		CreatedAt: dbutils.Time{
			Inner: time.Now(),
		},
		Active: true,
	}, &entities.UserData{
		UserId:             opmodels.DefaultUser,
		MFAEnabled:         false,
		MFAType:            "",
		MFAData:            "",
		PendingPassChange:  false,
		PendingEmailVerify: false,
		ExtraMeta:          nil,
		TenantID:           opts.Slug})

}
