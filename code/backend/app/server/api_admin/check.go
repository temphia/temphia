package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

func (r *ApiAdmin) checkAPI(rg *gin.RouterGroup) {
	rg.GET("/bprint/:bid", r.X(r.checkBprint))
	rg.GET("/plug/:bid", r.X(r.checkPlug))
	rg.GET("/dgroup/:bid", r.X(r.checkDataGroup))
	rg.GET("/dtable/:bid", r.X(r.checkDataTable))
}

func (r *ApiAdmin) checkBprint(ctx httpx.Request) {
	err := r.cAdmin.CheckBprint(ctx.Session, ctx.Http.Param("bid"))
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) checkPlug(ctx httpx.Request) {
	err := r.cAdmin.CheckPlug(ctx.Session, ctx.Http.Param("pid"))
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) checkDataGroup(ctx httpx.Request) {

}

func (r *ApiAdmin) checkDataTable(ctx httpx.Request) {

}
