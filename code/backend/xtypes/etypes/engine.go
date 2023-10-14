package etypes

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type RPXecuteOptions struct {
	TenantId string
	PlugId   string
	AgentId  string
	Action   string
	Payload  []byte
	Invoker  invoker.Invoker
}

type WebRawXecuteOptions struct {
	TenantId string
	PlugId   string
	AgentId  string
	Writer   http.ResponseWriter
	Request  *http.Request
}

type Remote struct {
	TenantId string
	PlugId   string
	AgentId  string
	Eid      string
	Action   string
	Data     xtypes.BeBytes
}

type Request struct {
	Id      string
	Name    string
	Data    xtypes.BeBytes
	Invoker invoker.Invoker
}

type RemoteOptions struct {
	TenantId      string `json:"tenant_id,omitempty"`
	PlugId        string `json:"plug_id,omitempty"`
	AgentId       string `json:"agent_id,omitempty"`
	Addr          string `json:"addr,omitempty"`
	RPXPrefix     string `json:"rpx_prefix,omitempty"`
	ControlPrefix string `json:"control_prefix,omitempty"`
	ReplyToken    string `json:"reply_token,omitempty"`
}

type Engine interface {
	Run() error

	GetCache() Ecache

	RPXecute(opts RPXecuteOptions) ([]byte, error)
	WebRawXecute(opts WebRawXecuteOptions)

	SetRemoteOption(opt RemoteOptions)
	ResetAgent(tenantId, plugId, agentId string) error

	ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error)

	ListExecutors() []string
	ListModules() []string

	RemotePerform(opt Remote) ([]byte, error)
	GetRemoteHandler() any
}

type Ecache interface {
	GetAgent(tenantId, plug, agent string) *entities.Agent
	GetPlug(tenantId, plugId string) *entities.Plug
}

/*

type RunningExec struct {
	EventId  string
	BprintId string
	PlugId   string
	AgentId  string
}

ListRunning(tenantId string) ([]RunningExec, error)

*/
