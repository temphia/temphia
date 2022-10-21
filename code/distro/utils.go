package distro

import (
	"errors"
	"os"

	"github.com/temphia/temphia/code/core/backend/controllers/operator/opmodels"
	"github.com/temphia/temphia/code/core/backend/controllers/operator/opsutils"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/upper/db/v4"
)

func (da *App) IsDbSchemaApplied() (bool, error) {
	err := da.CoreHub.Ping()
	if err != nil {
		return false, err
	}

	_, err = da.CoreHub.ListTenant()
	return err != nil, nil
}

func (da *App) IsTenantSeeded(tenantId string) (bool, error) {
	_, err := da.CoreHub.GetTenant(tenantId)

	if err != nil {
		if !errors.Is(err, db.ErrNoMoreRows) {
			return false, nil
		}

		return false, err
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
