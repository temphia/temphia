package job

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
)

type RPXJob struct {
	EventId   string
	Name      string
	Namespace string
	Payload   []byte
	Invoker   invoker.Invoker
	NodeTag   string
}

type RawWebJob struct {
	EventId string
	Writer  http.ResponseWriter
	Request *http.Request
}
