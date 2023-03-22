package seeder

import (
	"github.com/temphia/temphia/code/backend/controllers/operator/opmodels"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

// this is app seeder not data table  seeder

type AppSeeder struct {
	App     xtypes.App
	CoreHub store.CoreHub

	TenantName string
	TenantSlug string

	DefaultUserName  string
	DefaultUser      string
	DefaultGroupName string
	DefaultGroup     string
}

func NewAppSeeder(app xtypes.App) *AppSeeder {
	deps := app.GetDeps()

	return &AppSeeder{
		App:     app,
		CoreHub: deps.CoreHub().(store.CoreHub),

		TenantName: xtypes.DefaultTenantName,
		TenantSlug: xtypes.DefaultTenant,

		DefaultUserName:  opmodels.DefaultUserName,
		DefaultUser:      opmodels.DefaultUser,
		DefaultGroupName: opmodels.DefaultGroupName,
		DefaultGroup:     opmodels.DefaultGroup,
	}
}
