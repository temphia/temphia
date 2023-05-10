package authz

type AuthZ struct {
	scoper scoper
}

func (a *AuthZ) CheckScope(scope string) bool {
	return a.scoper.check(scope)
}

func (a *AuthZ) CheckAction(user, action string) (bool, error) {

	return false, nil
}
