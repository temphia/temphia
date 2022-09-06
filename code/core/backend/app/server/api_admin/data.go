package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (a *ApiAdmin) dataAPI(rg *gin.RouterGroup) {

	rg.GET("/", a.X(a.ListDtableSources))
	rg.GET("/:source", a.X(a.ListGroup))
	rg.POST("/:source/", a.X(a.NewGroup))
	rg.PATCH("/:source/:group_id", a.X(a.EditGroup))
	rg.GET("/:source/:group_id", a.X(a.GetGroup))
	rg.DELETE("/:source/:group_id", a.X(a.DeleteGroup))

	rg.GET("/", a.X(a.ListTables))
	rg.POST("/", a.X(a.AddTable))
	rg.GET("/:table_id", a.X(a.GetTable))
	rg.PATCH("/:table_id", a.X(a.EditTable))
	rg.DELETE("/:table_id", a.X(a.DeleteTable))
	rg.GET("/:table_id/column", a.X(a.ListColumns))
	rg.POST("/:table_id/column", a.X(a.AddColumn)) // fixme
	rg.PATCH("/:table_id/column/:column_id", a.X(a.EditColumn))
	rg.GET("/:table_id/column/:column_id", a.X(a.GetColumn))
	rg.DELETE("/:table_id/column/:column_id", a.X(a.DeleteColumn))

	rg.GET("/:table_id/view", a.X(a.ListView))
	rg.POST("/:table_id/view", a.X(a.NewView))
	rg.POST("/:table_id/view/:id", a.X(a.ModifyView))
	rg.GET("/:table_id/view/:id", a.X(a.GetView))
	rg.DELETE("/:table_id/view/:id", a.X(a.DelView))

	rg.GET("/:table_id/hook", a.X(a.ListHook))
	rg.POST("/:table_id/hook", a.X(a.NewHook))
	rg.POST("/:table_id/hook/:id", a.X(a.ModifyHook))
	rg.GET("/:table_id/hook/:id", a.X(a.GetHook))
	rg.DELETE("/:table_id/hook/:id", a.X(a.DelHook))

}

func (a *ApiAdmin) ListDtableSources(ctx httpx.Request) {
	// sources, err := r.cBasic.ListDyndbSources(ctx.Session)
	// r.rutil.WriteJSON(ctx.Http, sources, err)
}

// dyn_table_group

func (a *ApiAdmin) NewGroup(ctx httpx.Request) {
	tg := &bprints.NewTableGroup{}

	err := ctx.Http.BindJSON(tg)

	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
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
	err = a.cAdmin.EditGroup(ctx.Session, ctx.Http.Param("source"), ctx.Http.Param("group_id"), tg)
	a.rutil.WriteFinal(ctx.Http, err)
}

func (a *ApiAdmin) GetGroup(ctx httpx.Request) {
	resp, err := a.cAdmin.GetGroup(ctx.Session, ctx.Http.Param("source"), ctx.Http.Param("group_id"))
	a.rutil.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) ListGroup(ctx httpx.Request) {
	gr, err := a.cAdmin.ListGroup(ctx.Session, ctx.Http.Param("source"))
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	a.rutil.WriteJSON(ctx.Http, gr, err)
}

func (a *ApiAdmin) DeleteGroup(ctx httpx.Request) {
	err := a.cAdmin.DeleteGroup(ctx.Session, ctx.Http.Param("source"), ctx.Http.Param("group_id"))
	a.rutil.WriteFinal(ctx.Http, err)

}

// dyn_table

func (a *ApiAdmin) AddTable(ctx httpx.Request) {
	t := &bprints.NewTable{}
	err := ctx.Http.BindJSON(t)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = a.cAdmin.AddTable(ctx.Session, t)
	a.rutil.WriteFinal(ctx.Http, err)

}
func (a *ApiAdmin) EditTable(ctx httpx.Request) {
	tp := &entities.TablePartial{}
	err := ctx.Http.BindJSON(tp)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = a.cAdmin.EditTable(ctx.Session, ctx.Http.Param("table_id"), tp)
	a.rutil.WriteFinal(ctx.Http, err)

}

func (a *ApiAdmin) GetTable(ctx httpx.Request) {
	tbl, err := a.cAdmin.GetTable(ctx.Session, ctx.Http.Param("table_id"))
	a.rutil.WriteJSON(ctx.Http, tbl, err)
}

func (a *ApiAdmin) ListTables(ctx httpx.Request) {
	tbls, err := a.cAdmin.ListTables(ctx.Session)
	a.rutil.WriteJSON(ctx.Http, tbls, err)

}
func (a *ApiAdmin) DeleteTable(ctx httpx.Request) {
	err := a.cAdmin.DeleteTable(ctx.Session, ctx.Http.Param("table_id"))
	a.rutil.WriteFinal(ctx.Http, err)

}

// dyn_table_column

func (a *ApiAdmin) AddColumn(ctx httpx.Request) {
	nc := &bprints.NewColumn{}
	err := ctx.Http.BindJSON(nc)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = a.cAdmin.AddColumn(ctx.Session, ctx.Http.Param("table_id"), nc)
	a.rutil.WriteFinal(ctx.Http, err)

}
func (a *ApiAdmin) EditColumn(ctx httpx.Request) {
	cp := &entities.ColumnPartial{}
	err := ctx.Http.BindJSON(cp)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = a.cAdmin.EditColumn(ctx.Session, ctx.Http.Param("table_id"), ctx.Http.Param("column_id"), cp)
	a.rutil.WriteFinal(ctx.Http, err)

}

func (a *ApiAdmin) GetColumn(ctx httpx.Request) {
	resp, err := a.cAdmin.GetColumn(ctx.Session, ctx.Http.Param("table_id"), ctx.Http.Param("column_id"))
	a.rutil.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) ListColumns(ctx httpx.Request) {
	cols, err := a.cAdmin.ListColumns(ctx.Session, ctx.Http.Param("table_id"))
	a.rutil.WriteJSON(ctx.Http, cols, err)

}
func (a *ApiAdmin) DeleteColumn(ctx httpx.Request) {

	err := a.cAdmin.DeleteColumn(ctx.Session, ctx.Http.Param("table_id"), ctx.Http.Param("column_id"))
	a.rutil.WriteFinal(ctx.Http, err)

}
