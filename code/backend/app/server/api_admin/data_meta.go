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
		httpx.WriteErr(ctx.Http, err)
		return
	}
	err = a.cAdmin.NewView(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
		&view)
	httpx.WriteFinal(ctx.Http, err)
}

func (a *ApiAdmin) ModifyView(ctx httpx.Request) {
	view := make(map[string]interface{})

	err := ctx.Http.BindJSON(&view)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	err = a.cAdmin.ModifyView(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
		id, view)
	httpx.WriteFinal(ctx.Http, err)

}

func (a *ApiAdmin) GetView(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	resp, err := a.cAdmin.GetView(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
		id)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) ListView(ctx httpx.Request) {
	resp, err := a.cAdmin.ListView(ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
	)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	httpx.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) DelView(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	err = a.cAdmin.DelView(ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("gid"),
		ctx.MustParam("tid"),
		id)
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
