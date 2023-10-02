package engine2

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type Engine struct {
}

func (e *Engine) GetCache() etypes.Ecache {
	return nil
}
func (e *Engine) RPXecute(options etypes.Execution) ([]byte, error) {
	return nil, nil
}
func (e *Engine) WebRawXecute(rw http.ResponseWriter, req *http.Request) {

}

func (e *Engine) SetRemoteOption(opt any) {

}

func (e *Engine) ResetAgent(tenantId, plugId, agentId string) error { return nil }
func (e *Engine) ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return nil, nil
}

func (e *Engine) ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return nil, nil
}

func (e *Engine) RemotePerform(opt etypes.Remote) ([]byte, error) {
	return nil, nil
}
