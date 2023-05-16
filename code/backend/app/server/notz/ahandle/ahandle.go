package ahandle

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var (
	_ httpx.AdapterHandle = (*AHandle)(nil)
)

type Options struct {
	App         xtypes.App
	Logger      *zerolog.Logger
	DomainId    int64
	TenantId    string
	ResetDomain func()
}

// Ahandle is a common adapter utils handle
type AHandle struct {
	app       xtypes.App
	corehub   store.CoreHub
	logger    *zerolog.Logger
	domainId  int64
	tenantId  string
	resetFunc func()
}

func New(opts Options) *AHandle {

	return &AHandle{
		app:       opts.App,
		domainId:  opts.DomainId,
		corehub:   opts.App.GetDeps().CoreHub().(store.CoreHub),
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
