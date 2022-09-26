package authx

type Context interface {
	NewOauthClaim() (string, error)
	RunFencer(env any) error

	DisableUser() error
	GetUser() any
	NewUser() error
}
