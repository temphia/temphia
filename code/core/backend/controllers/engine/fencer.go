package engine

import (
	"sync"
	"time"

	"github.com/antonmedv/expr/vm"
	"github.com/jellydator/ttlcache/v3"
)

type Env struct {
	TenantId string
	PlugId   string
	AgentId  string
	Data     any
}

type fencer struct {
	vmPool *sync.Pool

	dataCache *ttlcache.Cache[string, *vm.Program]
	userCache *ttlcache.Cache[string, *vm.Program]
}

func NewFencer() *fencer {

	return &fencer{
		vmPool: &sync.Pool{
			New: func() any {
				return &vm.VM{}
			},
		},
		dataCache: newCache(),
		userCache: newCache(),
	}

}

func (f *fencer) Get() {

	item := f.dataCache.Get("aa")
	pg := item.Value()

	vm := f.vmPool.Get().(*vm.VM)

	vm.Run(pg, Env{
		TenantId: "",
		PlugId:   "",
		AgentId:  "",
		Data:     nil,
	})

}

func newCache() *ttlcache.Cache[string, *vm.Program] {

	cache := ttlcache.New(
		ttlcache.WithCapacity[string, *vm.Program](10),
		ttlcache.WithTTL[string, *vm.Program](time.Minute*10),
	)

	return cache
}
