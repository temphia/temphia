package job

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
)

type RPXJob struct {
	PlugId    string
	AgentId   string
	EventId   string
	Name      string
	Namespace string
	Payload   []byte
	Invoker   invoker.Invoker
	NodeTag   string
}

type RawWebJob struct {
	PlugId         string
	AgentId        string
	EventId        string
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}
