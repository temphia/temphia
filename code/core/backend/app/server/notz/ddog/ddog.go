package ddog

import "sync"

// ddog is short for domain dog
type Ddog struct {
	tenantdogs map[string]*TenantDdog
	errorsMap  map[string][2]int64 // [<int_counter, int_last_error_timestamp>]
	mlock      sync.RWMutex
}

func New() *Ddog {
	return &Ddog{
		tenantdogs: make(map[string]*TenantDdog),
		mlock:      sync.RWMutex{},
		errorsMap:  make(map[string][2]int64),
	}
}

func (d *Ddog) Renew(tenantId string) {

}

func (d *Ddog) Get(tenantId string) *TenantDdog {
	d.mlock.RLock()
	defer d.mlock.RUnlock()

	return d.tenantdogs[tenantId]
}
