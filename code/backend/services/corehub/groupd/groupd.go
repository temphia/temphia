package groupd

type Groupd interface {
	HasScope() bool
	HasScopeOrSuper() bool
	InvalidateCache(tenantId string, groupId string)
}

type Tenantd interface {
}
