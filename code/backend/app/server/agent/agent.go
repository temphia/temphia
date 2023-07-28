package agent

import (
	"net/http"
	"sync"
)

type agentState struct {
	webFiles       map[string]string
	spaConfig      any
	ssrConfig      any
	templateConfig any
}

type AgentServer struct {
	agents map[string]any
	aLock  sync.RWMutex
}

type AgentRequest struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	DomainId       int64
}

func (a *AgentServer) Render(ctx AgentRequest) {

}
