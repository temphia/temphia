package scopes

const (
	Cabinet = "cabinet"
	Data    = "data"
	Engine  = "engine"
	Repo    = "repo"
	User    = "user" // basic user listing and stuff (@mention user lookup)
	Admin   = "admin"
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
	AdminDev    = "admin.dev"
	AdminLog    = "admin.log"
	AdminData   = "admin.data"
)
