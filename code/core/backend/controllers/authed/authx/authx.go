package authx

type Options struct {
	Request any
	Handle  Handle
}

type Authenticator interface {
	HandleStage(stage string, opts Options) (any, error)
	HandleFinal(string) (any, error)
}

type Handle interface {
	NewOauthClaim() (string, error)
	RunFencer(env any) error

	DisableUser() error
	GetUser() any
	NewUser() error
}
