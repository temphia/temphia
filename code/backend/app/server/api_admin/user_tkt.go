package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (a *ApiAdmin) userTktAPI(rg *gin.RouterGroup) {
	rg.GET("/", a.UgroupX(a.ugListUser))
	rg.POST("/", a.UgroupX(a.ugAddUser))
	rg.GET("/:user_id", a.UgroupX(a.ugGetUser))
	rg.POST("/:user_id", a.UgroupX(a.ugUpdateUser))
	rg.DELETE("/:user_id", a.UgroupX(a.ugDeleteUser))
}

func (a *ApiAdmin) ugListUser(ctx UgCtx)   {}
func (a *ApiAdmin) ugAddUser(ctx UgCtx)    {}
func (a *ApiAdmin) ugGetUser(ctx UgCtx)    {}
func (a *ApiAdmin) ugUpdateUser(ctx UgCtx) {}
func (a *ApiAdmin) ugDeleteUser(ctx UgCtx) {}

type UgCtx struct {
	uclaim *claim.UserMgmtTkt
	http   *gin.Context
}

func (a *ApiAdmin) UgroupX(fn func(ctx UgCtx)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fn(UgCtx{
			uclaim: nil,
			http:   ctx,
		})
	}
}
