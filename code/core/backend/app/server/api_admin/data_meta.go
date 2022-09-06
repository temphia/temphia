package apiadmin

import (
	"strconv"

	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

// view stuff

func (a *ApiAdmin) NewView(ctx httpx.Request) {
	view := entities.DataView{}
	err := ctx.Http.BindJSON(&view)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}
	err = a.cAdmin.NewView(ctx.Session, ctx.Http.Param("table_id"), &view)
	httpx.WriteFinal(ctx.Http, err)
}

func (a *ApiAdmin) ModifyView(ctx httpx.Request) {
	view := make(map[string]interface{})

	err := ctx.Http.BindJSON(&view)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	err = a.cAdmin.ModifyView(ctx.Session, ctx.Http.Param("table_id"), id, view)
	httpx.WriteFinal(ctx.Http, err)

}

func (a *ApiAdmin) GetView(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := a.cAdmin.GetView(ctx.Session, ctx.Http.Param("table_id"), id)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) ListView(ctx httpx.Request) {
	resp, err := a.cAdmin.ListView(ctx.Session, ctx.Http.Param("table_id"))
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	httpx.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) DelView(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	err = a.cAdmin.DelView(ctx.Session, ctx.Http.Param("table_id"), id)
	httpx.WriteFinal(ctx.Http, err)
}

// hooks

func (a *ApiAdmin) NewHook(ctx httpx.Request) {
	hook := entities.DataHook{}
	err := ctx.Http.BindJSON(&hook)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}
	err = a.cAdmin.NewHook(ctx.Session, ctx.Http.Param("table_id"), &hook)
	httpx.WriteFinal(ctx.Http, err)
}

func (a *ApiAdmin) ModifyHook(ctx httpx.Request) {
	data := make(map[string]interface{})

	err := ctx.Http.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	err = a.cAdmin.ModifyHook(ctx.Session, ctx.Http.Param("table_id"), id, data)
	httpx.WriteFinal(ctx.Http, err)

}

func (a *ApiAdmin) GetHook(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := a.cAdmin.GetHook(ctx.Session, ctx.Http.Param("table_id"), id)
	httpx.WriteJSON(ctx.Http, resp, err)

}

func (a *ApiAdmin) ListHook(ctx httpx.Request) {
	resp, err := a.cAdmin.ListHook(ctx.Session, ctx.Http.Param("table_id"))
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	httpx.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) DelHook(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	err = a.cAdmin.DelHook(ctx.Session, ctx.Http.Param("table_id"), id)
	httpx.WriteFinal(ctx.Http, err)
}

func (a *ApiAdmin) AddIndex(ctx httpx.Request) {

}

func (a *ApiAdmin) AddUniqueIndex(ctx httpx.Request) {

}
func (a *ApiAdmin) AddFTSIndex(ctx httpx.Request) {

}
func (a *ApiAdmin) AddColumnFRef(ctx httpx.Request) {

}
func (a *ApiAdmin) ListIndex(ctx httpx.Request) {

}
func (a *ApiAdmin) RemoveIndex(ctx httpx.Request) {

}
