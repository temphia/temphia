package distro

import (
	"time"

	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (a *DistroApp) SeedSuperUser() error {

	err := a.AddNewUserGroup("Super Admins", xtypes.UserGroupSuperAdmin)
	if err != nil {
		return err
	}

	err = a.AddNewUser("superuser", "Super User", "admin@example.com", "super123", xtypes.UserGroupSuperAdmin)
	if err != nil {
		return err
	}

	return a.SendUserWelcome("superuser", "this is temphia interglactic information highway system connection portal blah blah")
}

func (a *DistroApp) AddNewUserGroup(gname, gslug string) error {

	return a.corehub().AddUserGroup(&entities.UserGroup{
		Name:     gname,
		Slug:     gslug,
		TenantID: a.app.TenantId(),
	})
}

func (a *DistroApp) AddNewUser(userId, fname, email, pass string, ugroup string) error {

	return a.corehub().AddUser(&entities.User{
		UserId:    userId,
		FullName:  fname,
		Email:     email,
		GroupID:   ugroup,
		Password:  pass,
		TenantID:  a.app.TenantId(),
		PublicKey: "",
		CreatedAt: dbutils.Time{
			Inner: time.Now(),
		},
		Active: true,
	}, &entities.UserData{
		UserId:             userId,
		MFAEnabled:         false,
		MFAType:            "",
		MFAData:            "",
		PendingPassChange:  false,
		PendingEmailVerify: false,
		ExtraMeta:          nil,
		TenantID:           a.app.TenantId()})

}

func (a *DistroApp) SendUserWelcome(user, msg string) error {
	now := time.Now()

	if msg == "" {
		msg = "this is temphia interglactic information highway system connection portal blah blah"
	}

	_, err := a.corehub().AddUserMessage(&entities.UserMessage{
		Title:    "Welcome User",
		Contents: msg,
		TenantId: a.app.TenantId(),
		UserId:   user,
		Type:     "system_message",
		CreatedAt: &dbutils.Time{
			Inner: now,
		},
	})
	return err
}

func (a *DistroApp) corehub() store.CoreHub {
	return a.app.GetDeps().CoreHub().(store.CoreHub)
}
