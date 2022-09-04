package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type Context struct {
	Rid  int64
	Http *gin.Context
}

type Adapter interface {
	Handle(ctx Context)
}

type BuilderOptions struct {
	App      xtypes.App
	TenantId string
	Domain   *entities.TenantDomain
}

type Builder func(opts BuilderOptions) (Adapter, error)
