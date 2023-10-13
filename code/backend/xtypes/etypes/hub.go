package etypes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/launch"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type BootData struct {
	TenantId   string `json:"tenant_id,omitempty"`
	ApiBaseURL string `json:"api_base_url,omitempty"`
	PlugId     string `json:"plug_id,omitempty"`
	AgentId    string `json:"agent_id,omitempty"`
	EntryName  string `json:"entry_name,omitempty"`
	ExecLoader string `json:"exec_loader,omitempty"`
}

type ModExecOptions struct {
	PlugId  string
	AgentId string
	Mod     string
	Method  string
	Data    xtypes.LazyData
}

// EngineHub is sits on top of Engine, Launch are related to TargetApps, Run are related to TargetHooks.
type EngineHub interface {
	GetEngine() Engine
	Start() error

	GetCache() Ecache

	LaunchAgent(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error)
	LaunchTarget(uclaim *claim.Session, targetId int64) (*launch.Response, error)
	LaunchEditor(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error)

	Execute(tenantId, action string, ctx *gin.Context)
	ExecuteDev(dclaim *claim.UserContext, plug, agent, action string, body []byte) ([]byte, error)
	Reset(tenantId, plugId, agentId string) error
	// ExecuteMod(dclaim *claim.UserContext, opts ModExecOptions) (xtypes.LazyData, error)

	ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ListExecutors() ([]string, error)
	ListModules() ([]string, error)

	GetExecutorBuilder(name string) ExecutorBuilder

	RunStartupHooks(tenants []string, minwait time.Duration)
	// 	RunDyndbHooks(tenants string, opts map[string]any) error
	// 	RunAdapterHooks(tenants string, opts map[string]any) error
	// 	RunUserHooks(tenants string, opts map[string]any) error

}

/*

	Events
	- Call(Link)  => sync/single => rpx(name, data)
	- Emit(Signal)  => async/multi  => rpx(handle, {cname, name, plugid, agentid, data})

	AddOn
	- Hijacker
	- Extension
	- Native

*/
