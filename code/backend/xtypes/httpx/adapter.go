package httpx

import (
	"io"

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
	Cache    GlobalCache
}

type Builder func(opts BuilderOptions) (Adapter, error)

type Context struct {
	Rid  int64
	Http *gin.Context
}

type AdapterEditorContext struct {
	Id   int64
	User *claim.UserContext
	Name string
	Data []byte
}

type Adapter interface {
	ServeEditorFile(file string) ([]byte, error)

	PreformEditorAction(ctx AdapterEditorContext) (any, error)

	Handle(ctx Context)

	Close() error
}

type AdapterHub interface {
	Serve(ctx *gin.Context)

	IsAllowed(tenantId, host string) bool

	ApplyTargetHook(tenantId string, id int64, data *entities.TargetHook)

	ApplyAdapter(tenantId string, id int64, data *entities.TenantDomain)

	PreformEditorAction(ctx AdapterEditorContext) (any, error)

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

type GlobalCache interface {
	GetSubCache(key string, loader CacheLoader) (SubCache, error)
}

type SubCache interface {
	Has(key string) (bool, error)
	Get(dst io.Writer, key string) error
}

type CacheLoader interface {
	Has(key string) (bool, error)
	Get(dst io.Writer, key string) error
	Put(key string, src io.Reader) error
}
