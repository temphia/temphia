package httpx

import "github.com/gin-gonic/gin"

type Context struct {
	Http *gin.Context
}

type Adapter interface {
	Handle(ctx Context)
}

type BuilderOptions struct {
	App      any // btypes.App
	TenantId string
	Domain   any // *entities.TenantDomain
}

type Builder func(opts BuilderOptions) (Adapter, error)
