package authx

import "github.com/temphia/temphia/code/backend/xtypes"

type Context interface {
	GetApp() xtypes.App

	NewOauthClaim() (string, error)
	RunFencer(env any) error

	DisableUser() error
	GetUser() any
	NewUser() error
}
