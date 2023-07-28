package xtypes

import "github.com/jaevor/go-nanoid"

const (
	DefaultTenantName = "Default"
	DefaultTenant     = "default0"
	BprintBlobFolder  = "bprints"
)

const (
	UserGroupSuperAdmin = "super_admin"
	UserGroupGuest      = "guest"
)

const (
	TEMPHIA_VER_MIN = 3
	TEMPHIA_VER_MAX = 0
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
