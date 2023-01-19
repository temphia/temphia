package distro

import (
	"errors"
	"os"
	"time"

	"github.com/temphia/temphia/code/backend/controllers/operator/opmodels"
	"github.com/temphia/temphia/code/backend/controllers/operator/opsutils"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/upper/db/v4"
)

func (da *App) IsDbSchemaApplied() (bool, error) {
	err := da.CoreHub.Ping()
	if err != nil {
		return false, err
	}

	_, err = da.CoreHub.ListTenant()
	return err == nil, nil
}

func (da *App) IsTenantSeeded(tenantId string) (bool, error) {
	_, err := da.CoreHub.GetTenant(tenantId)

	if err != nil {
		if !errors.Is(err, db.ErrNoMoreRows) {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

func (da *App) TenantSeed(tenantId string) error {
	superPass := os.Getenv("TEMPHIA_SUPER_PASSWORD")
	if superPass == "" {
		superPass = "super123"
	}

	superEmail := os.Getenv("TEMPHIA_SUPER_EMAIL")
	if superEmail == "" {
		superEmail = "admin@example.com"
	}

	return opsutils.AddTenant(da.App, &opmodels.NewTenant{
		Name:          "Default",
		Slug:          xtypes.DefaultTenant,
		SuperPassword: superPass,
		SuperEmail:    superEmail,
	})

}

func (da *App) IsDomainSeeded(tenantId string) (bool, error) {
	tds, err := da.CoreHub.ListDomain(tenantId)
	if err != nil {
		return false, err
	}
	return (len(tds) > 0), nil
}

func (da *App) SeedWildcardDomain(tenantId string) error {
	return da.CoreHub.AddDomain(&entities.TenantDomain{
		Name:           "*",
		About:          "Fallback Domain",
		AdapterType:    "easypage",
		AdapterOptions: entities.JsonStrMap{},
		TenantId:       tenantId,
		ExtraMeta:      entities.JsonStrMap{},
	})
}

func (da *App) SeedRepo(tenantId, bprint, user string) error {
	err := da.CoreHub.RepoNew(tenantId, &entities.Repo{
		Id:       0,
		Name:     "Embed",
		Provider: "embed",
		TenantId: tenantId,
	})
	if err != nil {
		return err
	}

	return nil

	// deps := da.App.GetDeps()

	// rhub := deps.RepoHub().(repox.Hub)
	// instancer := rhub.GetInstanceHub()

	// _, err = instancer.AutomaticBundle(repox.InstanceOptions{
	// 	BprintId:       bprint,
	// 	InstancerType:  xbprint.TypeBundle,
	// 	File:           "schema.json",
	// 	UserConfigData: nil,
	// 	Auto:           true,
	// 	UserSession: &claim.Session{
	// 		TenantId:  tenantId,
	// 		UserID:    user,
	// 		UserGroup: xtypes.UserGroupSuperAdmin,
	// 	},
	// })

	// return err
}

func (da *App) SeedWelcomeMessage(tenantId, to string) error {
	deps := da.App.GetDeps()
	ch := deps.CoreHub().(store.CoreHub)

	now := time.Now()

	_, err := ch.AddUserMessage(&entities.UserMessage{
		Title:     "Welcome User",
		Contents:  "this is temphia interglactic information highway system connection portal blah blah",
		TenantId:  tenantId,
		UserId:    to,
		Type:      "system_message",
		CreatedAt: &now,
	})

	return err

}
