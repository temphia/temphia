package agent

import (
	"net/http"
	"sync"
)

type AgentServer struct {
	agents map[string]*agentState
	aLock  sync.RWMutex
}

type AgentRequest struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	TenantId       string
	PlugId         string
	AgentId        string
	DomainId       int64
}

func (a *AgentServer) Render(ctx AgentRequest) {

}

// /z/agent_auth
