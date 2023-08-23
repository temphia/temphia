package scopes

import "github.com/temphia/temphia/code/backend/libx/easyerr"

var (
	ErrNoAdminUserScope   = easyerr.Error("scope not found [admin.user]")
	ErrNoAdminRepoScope   = easyerr.Error("scope not found [admin.repo]")
	ErrNoAdminTenantScope = easyerr.Error("scope not found [admin.tenant]")
	ErrNoAdminEngineScope = easyerr.Error("scope not found [admin.engine]")
	ErrNoAdminLensScope   = easyerr.Error("scope not found [admin.lens]")
	ErrNoAdminTargetScope = easyerr.Error("scope not found [admin.target]")
	ErrNoAdminSystemScope = easyerr.Error("scope not found [admin.system]")
	ErrNoAdminDataScope   = easyerr.Error("scope not found [admin.data]")

	ErrNoCabinetScope = easyerr.Error("scope not found [cabinet]")
	ErrNoDataScope    = easyerr.Error("scope not found [data]")
	ErrNoEngineScope  = easyerr.Error("scope not found [engine]")
	ErrNoRepoScope    = easyerr.Error("scope not found [repo]")
	ErrNoUserScope    = easyerr.Error("scope not found [user]")
)

const (
	Cabinet = "cabinet"
	Data    = "data"
	Engine  = "engine"
	Repo    = "repo"
	User    = "user" // basic user listing and stuff (@mention user lookup)
)

var (
	SuperScopes = []string{
		Cabinet, Data, Engine, Repo, User, AdminUser, AdminRepo, AdminTenant, AdminEngine, AdminDev, AdminLog, AdminData,
	}
)

const (
	AdminUser   = "admin.user"
	AdminRepo   = "admin.repo"
	AdminTenant = "admin.tenant"
	AdminEngine = "admin.engine"
	AdminTarget = "admin.target"
	AdminDev    = "admin.dev"
	AdminLog    = "admin.log"
	AdminData   = "admin.data"
)
