package etypes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type LaunchOptions struct {
	ApiBaseURL   string            `json:"api_base_url,omitempty"`
	Token        string            `json:"token,omitempty"`
	EntryName    string            `json:"entry,omitempty"`
	ExecLoader   string            `json:"exec_loader,omitempty"`
	JSPlugScript string            `json:"js_plug_script,omitempty"`
	StyleFile    string            `json:"style,omitempty"`
	ExtScripts   map[string]string `json:"ext_scripts,omitempty"`
	Plug         string            `json:"plug,omitempty"`
	Agent        string            `json:"agent,omitempty"`
}

type TargetLaunchData struct {
	TargetId   int64  `json:"target_id,omitempty"`
	TargetType string `json:"target_type,omitempty"`
}

type AdminLaunchData struct {
	PlugId  string `json:"plug_id,omitempty"`
	AgentId string `json:"agent_id,omitempty"`
}

type AuthLaunchData struct{}

type LaunchDomainOptions struct {
	LoaderScript string
	ScriptFile   string
	StyleFile    string
	ExtScripts   map[string]string
	BootData     string
	ApiBaseURL   string
	PlugId       string
	AgentId      string
	Name         string
	ExecLoader   string
}

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

	LaunchTargetDomain(tenantId, host, plugId, agentId string) (*LaunchDomainOptions, error)
	LaunchTarget(uclaim *claim.Session, data TargetLaunchData) (*LaunchOptions, error)
	LaunchAdmin(uclaim *claim.Session, data AdminLaunchData) (*LaunchOptions, error)

	Execute(tenantId, action string, ctx *gin.Context)
	ExecuteDev(dclaim *claim.UserContext, plug, agent, action string, body []byte) ([]byte, error)
	Reset(tenantId, plugId, agentId string) error
	// ExecuteMod(dclaim *claim.UserContext, opts ModExecOptions) (xtypes.LazyData, error)

	ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ListExecutors() ([]string, error)
	ListModules() ([]string, error)

	RunStartupHooks(tenants []string, minwait time.Duration)
	// 	RunDyndbHooks(tenants string, opts map[string]any) error
	// 	RunAdapterHooks(tenants string, opts map[string]any) error
	// 	RunUserHooks(tenants string, opts map[string]any) error

}
