package etypes

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type ExecutorBinder interface {
	bindx.Bindings

	GetApp() any
	GetLogger() *zerolog.Logger
	GetModuleInstance(id uint32) any
}
