package store

type AuthZ interface {
	CheckScope(scope string) bool
	CheckAction(user, action string) (bool, error)
}
