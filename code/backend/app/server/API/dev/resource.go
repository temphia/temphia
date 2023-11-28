package dev

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (d *DevAPI) resourceList(dclaim *claim.PlugDevTkt, ctx *gin.Context)       {}
func (d *DevAPI) resourceListAgent(dclaim *claim.PlugDevTkt, ctx *gin.Context)  {}
func (d *DevAPI) resourceLinkAgent(dclaim *claim.PlugDevTkt, ctx *gin.Context)  {}
func (d *DevAPI) resourceEdit(dclaim *claim.PlugDevTkt, ctx *gin.Context)       {}
func (d *DevAPI) resourceDelete(dclaim *claim.PlugDevTkt, ctx *gin.Context)     {}
func (d *DevAPI) resourceDeleteLink(dclaim *claim.PlugDevTkt, ctx *gin.Context) {}

func (d *DevAPI) resCfolderList(dclaim *claim.PlugDevTkt, ctx *gin.Context)         {}
func (d *DevAPI) resCfolderUpload(dclaim *claim.PlugDevTkt, ctx *gin.Context)       {}
func (d *DevAPI) resCfolderDownload(dclaim *claim.PlugDevTkt, ctx *gin.Context)     {}
func (d *DevAPI) resCfolderDelete(dclaim *claim.PlugDevTkt, ctx *gin.Context)       {}
func (d *DevAPI) resCfolderNewFolder(dclaim *claim.PlugDevTkt, ctx *gin.Context)    {}
func (d *DevAPI) resCfolderRenameFolder(dclaim *claim.PlugDevTkt, ctx *gin.Context) {}
func (d *DevAPI) resCfolderDeleteFolder(dclaim *claim.PlugDevTkt, ctx *gin.Context) {}

func (d *DevAPI) resDydbList(dclaim *claim.PlugDevTkt, ctx *gin.Context)     {}
func (d *DevAPI) resDydbStatus(dclaim *claim.PlugDevTkt, ctx *gin.Context)   {}
func (d *DevAPI) resDydbMigrate(dclaim *claim.PlugDevTkt, ctx *gin.Context)  {}
func (d *DevAPI) resDydbRollback(dclaim *claim.PlugDevTkt, ctx *gin.Context) {}
func (d *DevAPI) resDydbAutoseed(dclaim *claim.PlugDevTkt, ctx *gin.Context) {}

func (d *DevAPI) resSocketList(dclaim *claim.PlugDevTkt, ctx *gin.Context)    {}
func (d *DevAPI) resSocketStatus(dclaim *claim.PlugDevTkt, ctx *gin.Context)  {}
func (d *DevAPI) resSocketRefresh(dclaim *claim.PlugDevTkt, ctx *gin.Context) {}
func (d *DevAPI) resSocketWatch(dclaim *claim.PlugDevTkt, ctx *gin.Context)   {}
