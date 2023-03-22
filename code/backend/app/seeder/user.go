package seeder

import (
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (a *AppSeeder) AddNewUserGroup(gname, gslug string) error {
	return a.CoreHub.AddUserGroup(&entities.UserGroup{
		Name:     gname,
		Slug:     gslug,
		TenantID: a.TenantSlug,
	})
}

func (a *AppSeeder) AddNewUser(userId, fname, email, pass string, ugroup string) error {

	return a.CoreHub.AddUser(&entities.User{
		UserId:    userId,
		FullName:  fname,
		Email:     email,
		GroupID:   ugroup,
		Password:  pass,
		TenantID:  a.TenantSlug,
		PublicKey: "",
		CreatedAt: time.Now(),
		Active:    true,
	}, &entities.UserData{
		UserId:             userId,
		MFAEnabled:         false,
		MFAType:            "",
		MFAData:            "",
		PendingPassChange:  false,
		PendingEmailVerify: false,
		ExtraMeta:          nil,
		TenantID:           a.TenantSlug})

}

func (a *AppSeeder) SendWelcome() error {
	return a.SendUserWelcome(a.DefaultUser, "")
}

func (a *AppSeeder) SendUserWelcome(user, msg string) error {
	now := time.Now()

	if msg == "" {
		msg = "this is temphia interglactic information highway system connection portal blah blah"
	}

	_, err := a.CoreHub.AddUserMessage(&entities.UserMessage{
		Title:     "Welcome User",
		Contents:  msg,
		TenantId:  a.TenantSlug,
		UserId:    user,
		Type:      "system_message",
		CreatedAt: &now,
	})
	return err
}
