package xtypes

import "github.com/jaevor/go-nanoid"

const (
	DefaultTenantName = "Default"
	DefaultTenant     = "default0"
	BprintBlobFolder  = "bprints"
	DydnBlobFolder    = "dyndb"
)

const (
	UserGroupSuperAdmin = "super_admin"
	UserGroupGuest      = "guest"
)

const (
	TEMPHIA_VER_MIN = 3
	TEMPHIA_VER_MAX = 0
)

const (
	EnvDevPageServer = "TEMPHIA_DEV_PAGES_SERVER"

	EnvHost     = "TEMPHIA_HOST"
	EnvPort     = "TEMPHIA_PORT"
	EnvRePath   = "TEMPHIA_RE_PATH"
	EnvReToken  = "TEMPHIA_RE_TOKEN"
	EnvTenantId = "TEMPHIA_TENANT_ID"
	EnvPlugId   = "TEMPHIA_PLUG_ID"
	EnvAgentId  = "TEMPHIA_AGENT_ID"
	EnvBprintId = "TEMPHIA_BPRINT_ID"

	EnvLogdSecret = "TEMPHIA_LOGD_SECRET"
	EnvLogdPort   = "TEMPHIA_LOGD_PORT"

	EnvBdevBprintConfig = "TEMPHIA_BDEV_BPRINT_CONFIG"

	EnvBdevToken   = "TEMPHIA_BDEV_TOKEN"
	EnvBdevApiURL  = "TEMPHIA_BDEV_API_URL"
	EnvBdevPlugId  = "TEMPHIA_BDEV_PLUG_ID"
	EnvBdevAgentId = "TEMPHIA_BDEV_AGENT_ID"

	EnvAppInitSecret = "TEMPHIA_APP_INIT_SECRET"
)

func GetSlugGenerator(length int) func() string {
	gFunc, err := nanoid.CustomASCII("abcdefghijklmnopqrstuvwxyz1234567890", length)
	if err != nil {
		panic(err)
	}

	return func() string {
		return gFunc()
	}

}
