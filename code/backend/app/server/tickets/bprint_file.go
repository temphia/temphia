package tickets

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (t *TicketAPI) BprintFile(rg *gin.RouterGroup) {

	rg.GET("/", t.middleware.BprintX((t.bprintTktList)))

}

func (t *TicketAPI) bprintTktList(uclaim *claim.BprintTkt, ctx *gin.Context) {

}
