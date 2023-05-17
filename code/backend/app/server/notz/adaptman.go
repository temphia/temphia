package notz

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/app/server/adapters/common/cache"
	"github.com/temphia/temphia/code/backend/app/server/notz/adapter"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store"
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
	activeDomains     map[int64]*adapter.Adapter
	tenantInits       map[string]bool
	domainTenantIndex map[string]int64
	cache             httpx.GlobalCache

	// cloop chan
	cReInstance     chan DomainIdent
	cInstanceTenant chan string
	applogger       *zerolog.Logger
	logService      logx.Service
}

func newAdapterManager(app xtypes.App) AdapterManager {
	deps := app.GetDeps()

	ls := app.GetDeps().LogService().(logx.Service)

	return AdapterManager{
		app:               app,
		activeDomains:     make(map[int64]*adapter.Adapter),
		tenantInits:       make(map[string]bool),
		cReInstance:       make(chan DomainIdent),
		cInstanceTenant:   make(chan string),
		adapterBuilders:   deps.Registry().(*registry.Registry).GetAdapterBuilders(),
		corehub:           deps.CoreHub().(store.CoreHub),
		cabinethub:        deps.Cabinet().(store.CabinetHub),
		domainTenantIndex: make(map[string]int64),
		applogger:         ls.GetAppLogger(),
		logService:        ls,
		cache:             cache.New(),
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
		am.applogger.Error().
			Str("tenant_id", tenantId).
			Int64("domain_id", did).
			Str("handler", "serve_editor_file").
			Str("file", file).
			Msg(logid.NotzAdapterNotFound)

		return nil, easyerr.NotFound("domain instance")
	}

	out, err := instance.ServeEditorFile(file)
	if err != nil {
		am.applogger.Error().
			Str("tenant_id", tenantId).
			Int64("domain_id", did).
			Msg(logid.NotzAdapterEditorFileServeEr)

		return nil, err
	}

	return out, nil
}

func (am *AdapterManager) preformEditorAction(aclaim *claim.UserContext, name string, did int64, data []byte) (any, error) {

	instance := am.get(aclaim.TenantId, did)
	if instance == nil {
		am.applogger.Error().
			Str("tenant_id", aclaim.TenantId).
			Int64("domain_id", did).
			Str("handler", "perform_editor_action").
			Str("action", name).
			Msg(logid.NotzAdapterNotFound)

		return nil, easyerr.NotFound("domain instance")
	}

	resp, err := instance.PreformEditorAction(&claim.UserContext{
		TenantId:  aclaim.TenantId,
		UserID:    aclaim.UserID,
		UserGroup: aclaim.UserGroup,
		SessionID: aclaim.SessionID,
		DeviceId:  aclaim.DeviceId,
	}, name, data)
	if err != nil {
		am.applogger.Error().
			Str("tenant_id", aclaim.TenantId).
			Int64("domain_id", did).
			Msg(logid.NotzAdapterEditorActionErr)

		return nil, err
	}

	return resp, nil
}

// serve not /z

func (am *AdapterManager) Handle(tenantId, host string, ctx *gin.Context) {
	did := am.domainId(tenantId, host)
	instance := am.activeDomains[did]
	if instance == nil {
		am.applogger.Error().
			Str("tenant_id", tenantId).
			Str("host", host).
			Str("handler", "handle").
			Int64("domain_id", did).
			Msg(logid.NotzAdapterNotFound)
		return
	}

	instance.Handle(ctx)
}

// private

func (am *AdapterManager) get(tenantId string, did int64) *adapter.Adapter {
	instance := am.activeDomains[did]

	if instance == nil {
		time.Sleep(300)
		instance = am.activeDomains[did]
		if instance != nil {
			return instance
		}

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
