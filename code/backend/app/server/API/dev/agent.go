package dev

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (d *DevAPI) agentList(dclaim *claim.PlugDevTkt, ctx *gin.Context)       {}
func (d *DevAPI) agentReset(dclaim *claim.PlugDevTkt, ctx *gin.Context)      {}
func (d *DevAPI) agentWatch(dclaim *claim.PlugDevTkt, ctx *gin.Context)      {}
func (d *DevAPI) agentStatus(dclaim *claim.PlugDevTkt, ctx *gin.Context)     {}
func (d *DevAPI) agentUpdate(dclaim *claim.PlugDevTkt, ctx *gin.Context)     {}
func (d *DevAPI) agentRPXExecute(dclaim *claim.PlugDevTkt, ctx *gin.Context) {}
func (d *DevAPI) agentWebExecute(dclaim *claim.PlugDevTkt, ctx *gin.Context) {}
