package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type BuilderOptions struct {
	App      xtypes.App
	TenantId string
	Domain   *entities.TenantDomain
	Handle   AdapterHandle
}

type Builder func(opts BuilderOptions) (Adapter, error)

type Context struct {
	Rid  int64
	Http *gin.Context
}

type Adapter interface {
	ServeEditorFile(file string) ([]byte, error)

	PreformEditorAction(uclaim *claim.UserContext, name string, data []byte) (any, error)

	Handle(ctx Context)

	Close() error
}

type AdapterHub interface {
	Serve(ctx *gin.Context)

	IsAllowed(tenantId, host string) bool

	ApplyTargetHook(tenantId string, id int64, data *entities.TargetHook)

	ApplyAdapter(tenantId string, id int64, data *entities.TenantDomain)

	PreformEditorAction(uclaim *claim.UserContext, id int64, name string, data []byte) (any, error)

	ListAdapters() []string

	ServeEditorFile(tenantId, file string, did int64, ctx *gin.Context)

	ServePublic(file string, ctx *gin.Context)
}

// handle

type AdapterHandle interface {
	GetLogger() *zerolog.Logger

	LogInfo(rid int64) *zerolog.Event

	LogError(rid int64) *zerolog.Event

	SelfReset()
}
