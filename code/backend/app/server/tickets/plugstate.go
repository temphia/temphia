package tickets

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/tickets"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (a *TicketAPI) PlugState(rg *gin.RouterGroup) {

	rg.POST("/query", a.middleware.PSX(a.listPlugState))
	rg.POST("/key", a.middleware.PSX(a.addPlugState))
	rg.GET("/key/:key", a.middleware.PSX(a.getPlugState))
	rg.POST("/key/:key", a.middleware.PSX(a.updatePlugState))
	rg.DELETE("/key/:key", a.middleware.PSX(a.deletePlugState))

}

func (a *TicketAPI) addPlugState(aclaim *claim.PlugState, ctx *gin.Context) {
	data := tickets.AddPlugStateOptions{}
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = a.cTicket.AddPlugState(aclaim, data)
	httpx.WriteFinal(ctx, err)
}

func (a *TicketAPI) getPlugState(aclaim *claim.PlugState, ctx *gin.Context) {
	resp, err := a.cTicket.GetPlugState(aclaim, ctx.Param("key"))
	httpx.WriteJSON(ctx, resp, err)
}

func (a *TicketAPI) updatePlugState(aclaim *claim.PlugState, ctx *gin.Context) {

	data := tickets.UpdatePlugStateOptions{}
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = a.cTicket.UpdatePlugState(aclaim, data)
	httpx.WriteFinal(ctx, err)
}

func (a *TicketAPI) deletePlugState(aclaim *claim.PlugState, ctx *gin.Context) {
	err := a.cTicket.DeletePlugState(aclaim, ctx.Param("key"))
	httpx.WriteFinal(ctx, err)
}

func (a *TicketAPI) listPlugState(aclaim *claim.PlugState, ctx *gin.Context) {

	query := store.PkvQuery{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := a.cTicket.ListPlugState(aclaim, &query)
	httpx.WriteJSON(ctx, resp, err)
}
