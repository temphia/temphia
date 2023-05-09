package ahandle

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var (
	_ httpx.AdapterHandle = (*AHandle)(nil)
)

type Options struct {
	Corehub     store.CoreHub
	Logger      *zerolog.Logger
	DomainId    int64
	TenantId    string
	ResetDomain func()
}

type AHandle struct {
	corehub   store.CoreHub
	logger    *zerolog.Logger
	domainId  int64
	tenantId  string
	resetFunc func()
}

func New(opts Options) *AHandle {
	return &AHandle{
		corehub:   opts.Corehub,
		logger:    opts.Logger,
		tenantId:  opts.TenantId,
		resetFunc: opts.ResetDomain,
	}
}

func (ah *AHandle) SelfReset() {
	ah.resetFunc()
}

// log
func (ah *AHandle) GetLogger() *zerolog.Logger {
	return ah.logger
}

func (ah *AHandle) LogInfo(rid int64) *zerolog.Event {
	return ah.logger.Info().Int64("rid", rid)
}

func (ah *AHandle) LogError(rid int64) *zerolog.Event {
	return ah.logger.Error().Int64("rid", rid)
}

func (ah *AHandle) GetKvToken() (string, error) {

	return "", nil
}

func (ah *AHandle) Init() error {
	key := fmt.Sprintf("adapter-%d", ah.domainId)

	bp, _ := ah.corehub.BprintGet(ah.tenantId, key)
	if bp == nil {
		return ah.corehub.BprintNew(ah.tenantId, &entities.BPrint{
			ID:          key,
			Name:        "Domain Adpter Container",
			Slug:        "domain-adapter",
			Type:        "container",
			Description: fmt.Sprintf("This is a bprint controlled by adapter %d", ah.domainId),
			TenantID:    ah.tenantId,
			Tags:        entities.JsonArray{},
			Files:       entities.JsonArray{},
			ExtraMeta:   entities.JsonMap{},
		})
	}

	plug, _ := ah.corehub.PlugGet(ah.tenantId, key)
	if plug == nil {
		ah.corehub.PlugNew(ah.tenantId, &entities.Plug{
			Id:        key,
			Name:      fmt.Sprintf("Domain Adpter plug %d", ah.domainId),
			Live:      true,
			BprintId:  key,
			ExtraMeta: entities.JsonStrMap{},
			TenantId:  ah.tenantId,
		})
	}

	return nil
}
