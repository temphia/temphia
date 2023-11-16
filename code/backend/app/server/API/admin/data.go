package apiadmin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/admin"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (a *ApiAdmin) dataAPI(rg *gin.RouterGroup) {

	rg.GET("/:source/group", a.X(a.ListGroup))
	rg.POST("/:source/group", a.X(a.NewGroup))
	rg.PATCH("/:source/group/:gid", a.X(a.EditGroup))
	rg.GET("/:source/group/:gid", a.X(a.GetGroup))
	rg.DELETE("/:source/group/:gid", a.X(a.DeleteGroup))

	rg.GET("/:source/group/:gid/sheet", a.X(a.GetGroupSheet))
	rg.POST("/:source/group/:gid/query", a.X(a.query))

	rg.GET("/:source/group/:gid/table", a.X(a.ListTables))
	rg.GET("/:source/group/:gid/table/:tid", a.X(a.GetTable))
	rg.PATCH("/:source/group/:gid/table/:tid", a.X(a.EditTable))
	rg.DELETE("/:source/group/:gid/table/:tid", a.X(a.DeleteTable))

	rg.GET("/:source/group/:gid/table/:tid/activity", a.X(a.QueryActivity))
	rg.POST("/:source/group/:gid/table/:tid/autoseed", a.X(a.autoSeed))

	rg.GET("/:source/group/:gid/table/:tid/column", a.X(a.ListColumns))
	rg.PATCH("/:source/group/:gid/table/:tid/column/:cid", a.X(a.EditColumn))
	rg.GET("/:source/group/:gid/table/:tid/column/:cid", a.X(a.GetColumn))
	rg.DELETE("/:source/group/:gid/table/:tid/column/:cid", a.X(a.DeleteColumn))

	rg.GET("/:source/group/:gid/table/:tid/view", a.X(a.ListView))
	rg.POST("/:source/group/:gid/table/:tid/view", a.X(a.NewView))
	rg.POST("/:source/group/:gid/table/:tid/view/:id", a.X(a.ModifyView))
	rg.GET("/:source/group/:gid/table/:tid/view/:id", a.X(a.GetView))
	rg.DELETE("/:source/group/:gid/table/:tid/view/:id", a.X(a.DelView))

}

// dyn_table_group

func (a *ApiAdmin) NewGroup(ctx httpx.Request) {
	tg := &xpackage.NewTableGroup{}

	err := ctx.Http.BindJSON(tg)

	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	err = a.cAdmin.NewGroup(ctx.Session, ctx.Http.Param("source"), tg)
	a.rutil.WriteFinal(ctx.Http, err)
}

func (a *ApiAdmin) EditGroup(ctx httpx.Request) {
	tg := &entities.TableGroupPartial{}
	err := ctx.Http.BindJSON(tg)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = a.cAdmin.EditGroup(ctx.Session, ctx.MustParam("source"), ctx.MustParam("gid"), tg)
	a.rutil.WriteFinal(ctx.Http, err)
}

func (a *ApiAdmin) GetGroup(ctx httpx.Request) {
	resp, err := a.cAdmin.GetGroup(ctx.Session, ctx.MustParam("source"), ctx.MustParam("gid"))
	a.rutil.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) GetGroupSheet(ctx httpx.Request) {
	resp, err := a.cAdmin.GetGroupSheets(ctx.Session, ctx.MustParam("source"), ctx.MustParam("gid"))
	a.rutil.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) ListGroup(ctx httpx.Request) {
	gr, err := a.cAdmin.ListGroup(ctx.Session, ctx.MustParam("source"))
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	a.rutil.WriteJSON(ctx.Http, gr, err)
}

func (a *ApiAdmin) DeleteGroup(ctx httpx.Request) {
	err := a.cAdmin.DeleteGroup(ctx.Session, ctx.MustParam("source"), ctx.MustParam("gid"))
	a.rutil.WriteFinal(ctx.Http, err)

}

// dyn_table

func (a *ApiAdmin) EditTable(ctx httpx.Request) {
	tp := &entities.TablePartial{}
	err := ctx.Http.BindJSON(tp)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = a.cAdmin.EditTable(ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"), tp)
	a.rutil.WriteFinal(ctx.Http, err)

}

func (a *ApiAdmin) GetTable(ctx httpx.Request) {
	tbl, err := a.cAdmin.GetTable(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
	)
	a.rutil.WriteJSON(ctx.Http, tbl, err)
}

func (a *ApiAdmin) ListTables(ctx httpx.Request) {
	tbls, err := a.cAdmin.ListTables(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
	)
	a.rutil.WriteJSON(ctx.Http, tbls, err)

}
func (a *ApiAdmin) DeleteTable(ctx httpx.Request) {
	err := a.cAdmin.DeleteTable(ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
	)

	a.rutil.WriteFinal(ctx.Http, err)

}

// dyn_table_column

func (a *ApiAdmin) EditColumn(ctx httpx.Request) {
	cp := &entities.ColumnPartial{}
	err := ctx.Http.BindJSON(cp)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = a.cAdmin.EditColumn(ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
		ctx.MustParam("cid"),
		cp)
	a.rutil.WriteFinal(ctx.Http, err)

}

func (a *ApiAdmin) GetColumn(ctx httpx.Request) {
	resp, err := a.cAdmin.GetColumn(ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
		ctx.MustParam("cid"),
	)
	a.rutil.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) ListColumns(ctx httpx.Request) {
	cols, err := a.cAdmin.ListColumns(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
	)
	a.rutil.WriteJSON(ctx.Http, cols, err)

}

func (a *ApiAdmin) DeleteColumn(ctx httpx.Request) {

	err := a.cAdmin.DeleteColumn(ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
		ctx.MustParam("cid"),
	)

	a.rutil.WriteFinal(ctx.Http, err)

}

func (a *ApiAdmin) QueryActivity(ctx httpx.Request) {

	offset, err := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := a.cAdmin.DataActivityQuery(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
		offset,
	)

	a.rutil.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) query(ctx httpx.Request) {
	opts := admin.DataGroupQuery{}

	err := ctx.Http.BindJSON(&opts)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := a.cAdmin.QueryDataGroup(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		opts,
	)

	a.rutil.WriteJSON(ctx.Http, resp, err)
}

type autoSeedRequest struct {
	Max int `json:"max"`
}

func (a *ApiAdmin) autoSeed(ctx httpx.Request) {

	r := &autoSeedRequest{}
	err := ctx.Http.BindJSON(r)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	a.cAdmin.AutoSeed(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
		r.Max,
	)

}
