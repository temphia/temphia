package net

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

var (
	StaticJsonHeader = map[string]string{ContentType: ApplicationJson}
)

var _ bindx.Net = (*Binding)(nil)

type Binding struct {
	hclient http.Client
}

func New() Binding {
	return Binding{
		hclient: http.Client{},
	}
}

func (n *Binding) HttpRaw(req *bindx.HttpRequest) *bindx.HttpResponse {
	return httpRaw(&n.hclient, req)
}

func (n *Binding) HttpRawBatch(reqs []*bindx.HttpRequest) []*bindx.HttpResponse {
	return n.httpRawBatch(reqs)
}

func (n *Binding) HttpQuickGet(url string, headers map[string]string) ([]byte, error) {
	return n.httpQuickGet(url, headers)
}

func (n *Binding) HttpQuickPost(url string, headers map[string]string, data []byte) ([]byte, error) {
	return n.httpQuickPost(url, headers, data)
}

func (n *Binding) HttpFormPost(url string, headers map[string]string, data []byte) ([]byte, error) {
	return n.httpFormPost(url, headers, data)
}

func (n *Binding) HttpJsonGet(url string, headers map[string]string) ([]byte, error) {
	return n.httpJsonGet(url, headers)
}

func (n *Binding) HttpJsonPost(url string, headers map[string]string, data []byte) ([]byte, error) {
	return n.httpJsonPost(url, headers, data)
}
