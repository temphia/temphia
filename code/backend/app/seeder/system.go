package seeder

import (
	"errors"
	"os"

	"github.com/temphia/temphia/code/backend/controllers/operator/opmodels"
	"github.com/temphia/temphia/code/backend/controllers/operator/opsutils"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (a *AppSeeder) IsDbSchemaApplied() (bool, error) {
	err := a.CoreHub.Ping()
	if err != nil {
		return false, err
	}

	_, err = a.CoreHub.ListTenant()
	return err == nil, nil
}

func (a *AppSeeder) IsTenantSeeded() (bool, error) {
	_, err := a.CoreHub.GetTenant(a.TenantSlug)

	if err != nil {
		if !errors.Is(err, db.ErrNoMoreRows) {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

func (a *AppSeeder) TenantSeed() error {
	superPass := os.Getenv("TEMPHIA_SUPER_PASSWORD")
	if superPass == "" {
		superPass = "super123"
	}

	superEmail := os.Getenv("TEMPHIA_SUPER_EMAIL")
	if superEmail == "" {
		superEmail = "admin@example.com"
	}

	return opsutils.AddTenantWithUser(a.CoreHub, &opsutils.TenantWithUserOptions{
		Name:           a.TenantName,
		Slug:           a.TenantSlug,
		SuperPassword:  superPass,
		SuperEmail:     superEmail,
		SuperUserName:  opmodels.DefaultUserName,
		SuperUser:      opmodels.DefaultUser,
		SuperGroupName: opmodels.DefaultGroupName,
		SuperGroup:     opmodels.DefaultGroup,
	})

}

func (a *AppSeeder) IsDomainSeeded() (bool, error) {
	tds, err := a.CoreHub.ListDomain(a.TenantSlug)
	if err != nil {
		return false, err
	}
	return (len(tds) > 0), nil
}

func (a *AppSeeder) SeedWildcardDomain() error {
	return a.CoreHub.AddDomain(&entities.TenantDomain{
		Name:           "*",
		About:          "Fallback Domain",
		AdapterType:    "easypage",
		AdapterOptions: entities.JsonStrMap{},
		TenantId:       a.TenantSlug,
		ExtraMeta:      entities.JsonStrMap{},
	})
}

func (a *AppSeeder) SeedRepo() error {
	err := a.CoreHub.RepoNew(a.TenantSlug, &entities.Repo{
		Id:       0,
		Name:     "Embed",
		Provider: "embed",
		TenantId: a.TenantSlug,
	})
	if err != nil {
		return err
	}

	err = a.CoreHub.RepoNew(a.TenantSlug, &entities.Repo{
		Id:       0,
		Name:     "Official",
		Provider: "official",
		TenantId: a.TenantSlug,
	})

	if err != nil {
		return err
	}

	return nil

}
