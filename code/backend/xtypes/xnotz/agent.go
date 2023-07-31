package xnotz

import (
	"net/http"
)

type Context struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	TenantId  string
	PlugId    string
	AgentId   string
	DomainId  int64
	RequestId int64
}
