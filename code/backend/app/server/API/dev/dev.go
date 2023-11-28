package dev

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/controllers/dev"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

type DevAPI struct {
	signer service.Signer
	devc   *dev.Controller
}

func New(signer service.Signer, devc *dev.Controller) *DevAPI {
	return &DevAPI{
		signer: signer,
		devc:   devc,
	}
}

func (d *DevAPI) DevAPI(rg *gin.RouterGroup) {

	rg.GET("/artifact", d.devX(d.artifactList))
	rg.POST("/artifact", d.devX(d.artifactPush))
	rg.GET("/artifact/:file", d.devX(d.artifactGet))
	rg.DELETE("/artifact/:file", d.devX(d.artifactDelete))
	rg.PUT("/artifact", d.devX(d.artifactBulkPush))

	rg.GET("/plug", d.devX(d.plugStatus))
	rg.POST("/plug", d.devX(d.plugUpgrade))

	rg.GET("/agent", d.devX(d.agentList))
	rg.PATCH("/agent/:agent_id", d.devX(d.agentReset))
	rg.GET("/agent/:agent_id/watch", d.devX(d.agentWatch))
	rg.GET("/agent/:agent_id", d.devX(d.agentStatus))
	rg.POST("/agent/:agent_id", d.devX(d.agentUpdate))

	rg.POST("/agent/:agent_id/rpx", d.devX(d.agentRPXExecute))
	rg.POST("/agent/:agent_id/web", d.devX(d.agentWebExecute))

	rg.GET("/resource", d.devX(d.resourceList))
	rg.GET("/resource/agent", d.devX(d.resourceListAgent))
	rg.POST("/resource/link", d.devX(d.resourceLinkAgent))
	rg.POST("/resource/edit/:id", d.devX(d.resourceEdit))
	rg.DELETE("/resource/delete/:id", d.devX(d.resourceDelete))
	rg.DELETE("/resource/link/:id", d.devX(d.resourceDeleteLink))

	rg.GET("/res/cfolder", d.devX(d.resCfolderList))
	rg.POST("/res/cfolder/:folder/:fname", d.devX(d.resCfolderUpload))
	rg.GET("/res/cfolder/:folder/:fname", d.devX(d.resCfolderDownload))
	rg.DELETE("/res/cfolder/:folder", d.devX(d.resCfolderDelete))
	rg.PATCH("/res/cfolder/:folder/:fname", d.devX(d.resCfolderNewFolder))
	rg.PUT("/res/cfolder/:folder/:fname", d.devX(d.resCfolderRenameFolder))
	rg.DELETE("/res/cfolder/:folder/:fname", d.devX(d.resCfolderDeleteFolder))

	rg.GET("/res/dyndb", d.devX(d.resDydbList))
	rg.GET("/res/dyndb/:id", d.devX(d.resDydbStatus))
	rg.POST("/res/dyndb/:id", d.devX(d.resDydbMigrate))
	rg.PATCH("/res/dyndb/:id", d.devX(d.resDydbRollback))
	rg.PUT("/res/dyndb/:id", d.devX(d.resDydbAutoseed))

	rg.GET("/res/socket", d.devX(d.resSocketList))
	rg.GET("/res/socket/:id", d.devX(d.resSocketStatus))
	rg.POST("/res/socket/:id", d.devX(d.resSocketRefresh))
	rg.GET("/res/socket/:id/watch", d.devX(d.resSocketWatch))

}

func (d *DevAPI) devX(fn func(dclaim *claim.PlugDevTkt, ctx *gin.Context)) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

		authtok := ctx.GetHeader("Authorization")
		if authtok == "" {
			authtok = ctx.Query("token")
		}

		tkt, err := d.signer.ParsePlugDevTkt(ctx.Param("tenant_id"), authtok)
		if err != nil {
			httpx.UnAuthorized(ctx)
			pp.Println(err)
			return
		}

		fn(tkt, ctx)
	}

}
