package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type Context struct {
	Rid  int64
	Http *gin.Context
}

type Adapter interface {
	ServeEditorFile(file string) ([]byte, error)
	PreformEditorAction(name string, data []byte) (any, error)

	Handle(ctx Context)
}

type AdapterHandler interface {
	// log
	LogInfo(rid int64) *zerolog.Event
	LogError(rid int64) *zerolog.Event
}

type BuilderOptions struct {
	App            xtypes.App
	TenantId       string
	Domain         *entities.TenantDomain
	AdapterHandler AdapterHandler
}

type Builder func(opts BuilderOptions) (Adapter, error)
