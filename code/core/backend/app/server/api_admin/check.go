package apiadmin

import "github.com/temphia/temphia/code/core/backend/xtypes/httpx"

func (r *ApiAdmin) CheckBprint(ctx httpx.Request) {
	err := r.cAdmin.CheckBprint(ctx.Session, ctx.Http.Param("bid"))
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) CheckPlug(ctx httpx.Request) {
	err := r.cAdmin.CheckPlug(ctx.Session, ctx.Http.Param("pid"))
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) CheckDataGroup(ctx httpx.Request) {

}

func (r *ApiAdmin) CheckDataTable(ctx httpx.Request) {

}
