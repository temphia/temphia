package etypes

import (
	"github.com/gin-gonic/gin"
)

type Engine interface {
	Run() error
	GetRuntime() Runtime

	ServerLaunchExec(tenantId, plugId, agentId, mode string, arg any, resp any) error
	ClientLaunchExec(tenantId, plugId, agentId, mode string, ctx *gin.Context)

	ExecAction(tenantId, plugId, agentId, action string, ctx *gin.Context)
	ServePlugFile(tenantId, plugId, agentId, file string, ctx *gin.Context)
	ServeExecutorFile(tenantId, plugId, agentId, loader string, ctx *gin.Context)
}
