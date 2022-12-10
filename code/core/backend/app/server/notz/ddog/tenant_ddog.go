package ddog

type TenantDdog struct {
	// domains map[string]*entities.TenantDomain

	AnyServe     bool
	AnyNotZServe bool
	AnyAPIServe  bool
}

type Tddog interface {
	CorsCheck(domain string) (bool, error)
	CanServe(domain string) (bool, error)
}
