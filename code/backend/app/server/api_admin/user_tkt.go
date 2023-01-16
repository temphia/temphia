package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (a *ApiAdmin) userTktAPI(rg *gin.RouterGroup) {
	rg.GET("/", a.UgroupX(a.ugListUser))
	rg.POST("/", a.UgroupX(a.ugAddUser))
	rg.GET("/:user_id", a.UgroupX(a.ugGetUser))
	rg.POST("/:user_id", a.UgroupX(a.ugUpdateUser))
	rg.DELETE("/:user_id", a.UgroupX(a.ugDeleteUser))
}

func (a *ApiAdmin) ugListUser(ctx UgCtx) {
	resp, err := a.cAdmin.UgroupListUsersByGroup(ctx.uclaim)
	a.rutil.WriteJSON(ctx.http, resp, err)
}

func (a *ApiAdmin) ugAddUser(ctx UgCtx) {
	user := &entities.User{}
	err := ctx.http.BindJSON(user)
	if err != nil {
		a.rutil.WriteErr(ctx.http, err.Error())
		return
	}

	err = a.cAdmin.UgroupAddUser(ctx.uclaim, user)
	a.rutil.WriteFinal(ctx.http, err)
}

func (a *ApiAdmin) ugGetUser(ctx UgCtx) {
	resp, err := a.cAdmin.UgroupGetUserByID(ctx.uclaim, ctx.http.Param("user_id"))
	a.rutil.WriteJSON(ctx.http, resp, err)
}

func (a *ApiAdmin) ugUpdateUser(ctx UgCtx) {
	data := make(map[string]any)
	err := ctx.http.BindJSON(&data)
	if err != nil {
		a.rutil.WriteErr(ctx.http, err.Error())
		return
	}

	err = a.cAdmin.UgroupUpdateUser(ctx.uclaim, ctx.http.Param("user_id"), data)
	a.rutil.WriteFinal(ctx.http, err)
}

func (a *ApiAdmin) ugDeleteUser(ctx UgCtx) {
	err := a.cAdmin.UgroupDeleteUser(ctx.uclaim, ctx.http.Param("user_id"))
	a.rutil.WriteFinal(ctx.http, err)
}

type UgCtx struct {
	uclaim *claim.UserMgmtTkt
	http   *gin.Context
}

func (a *ApiAdmin) UgroupX(fn func(ctx UgCtx)) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

		umClaim, err := a.signer.ParseUserMgmtTkt(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			httpx.UnAuthorized(ctx)
			return
		}

		fn(UgCtx{
			uclaim: umClaim,
			http:   ctx,
		})
	}
}
