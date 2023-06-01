package tickets

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (t *TicketAPI) ExecMod(rg *gin.RouterGroup) {

	rg.GET("/", t.middleware.BprintX((t.execModList)))

}

func (t *TicketAPI) execModList(uclaim *claim.BprintTkt, ctx *gin.Context) {

}
