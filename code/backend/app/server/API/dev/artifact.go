package dev

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (d *DevAPI) artifactList(dclaim *claim.PlugDevTkt, ctx *gin.Context)     {}
func (d *DevAPI) artifactPush(dclaim *claim.PlugDevTkt, ctx *gin.Context)     {}
func (d *DevAPI) artifactGet(dclaim *claim.PlugDevTkt, ctx *gin.Context)      {}
func (d *DevAPI) artifactDelete(dclaim *claim.PlugDevTkt, ctx *gin.Context)   {}
func (d *DevAPI) artifactBulkPush(dclaim *claim.PlugDevTkt, ctx *gin.Context) {}
