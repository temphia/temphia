package notz

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type DomainIdent struct {
	tenantId string
	domainId int64
}

type AdapterManager struct {
	app               xtypes.App
	corehub           store.CoreHub
	cabinethub        store.CabinetHub
	adapterBuilders   map[string]httpx.Builder
	activeDomains     map[int64]*DomainInstance
	tenantInits       map[string]bool
	domainTenantIndex map[string]int64

	// cloop chan
	cReInstance     chan DomainIdent
	cInstanceTenant chan string
}

func newAdapterManager(app xtypes.App) AdapterManager {
	deps := app.GetDeps()

	return AdapterManager{
		app:               app,
		activeDomains:     make(map[int64]*DomainInstance),
		tenantInits:       make(map[string]bool),
		cReInstance:       make(chan DomainIdent),
		cInstanceTenant:   make(chan string),
		adapterBuilders:   deps.Registry().(*registry.Registry).GetAdapterBuilders(),
		corehub:           deps.CoreHub().(store.CoreHub),
		cabinethub:        deps.Cabinet().(store.CabinetHub),
		domainTenantIndex: make(map[string]int64),
	}

}

func (am *AdapterManager) ListAdapters() []string {

	keys := make([]string, 0, len(am.adapterBuilders))
	for k := range am.adapterBuilders {
		keys = append(keys, k)
	}

	return keys
}

func (am *AdapterManager) serveEditorFile(tenantId string, did int64, file string) ([]byte, error) {

	instance := am.get(tenantId, did)
	if instance == nil {
		return nil, easyerr.NotFound()
	}

	return instance.serveEditorFile(file)
}

func (am *AdapterManager) preformEditorAction(tenantId, name string, did int64, data []byte) (any, error) {

	instance := am.get(tenantId, did)
	if instance != nil {
		return nil, easyerr.NotFound()
	}

	return instance.preformEditorAction(name, data)
}

// serve not /z

func (am *AdapterManager) Handle(tenantId, host string, ctx *gin.Context) {
	did := am.domainId(tenantId, host)

	pp.Println("@serve did error", did)

	instance := am.activeDomains[did]

	pp.Println("@serve instance", instance)

	instance.handle(ctx)
}

// private

func (am *AdapterManager) get(tenantId string, did int64) *DomainInstance {
	instance := am.activeDomains[did]

	if instance == nil {
		if am.tenantInits[tenantId] {
			am.cInstanceTenant <- tenantId
		} else {
			am.cReInstance <- DomainIdent{
				tenantId: tenantId,
				domainId: did,
			}
		}

		for i := 0; i < 4; i = i + 1 {
			time.Sleep(200)
			instance := am.activeDomains[did]
			if instance != nil {
				break
			}
		}
	}

	return instance
}

func (am *AdapterManager) domainId(tenantId, domain string) int64 {

	did, ok := am.domainTenantIndex[domain+"|"+tenantId]
	if ok {
		return did
	}

	dparts := strings.Split(domain, ".")
	dparts[0] = "*"

	did, ok = am.domainTenantIndex[strings.Join(dparts, ".")+"|"+tenantId]
	if ok {
		return did
	}

	return am.domainTenantIndex["*|"+tenantId]

}
