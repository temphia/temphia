package http

type HttpRequest struct {
	Method  string            `json:"method,omitempty"`
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    []byte            `json:"body,omitempty"`
}

type HttpResponse struct {
	SatusCode int                 `json:"status_code,omitempty"`
	Headers   map[string][]string `json:"headers,omitempty"`
	Body      []byte              `json:"body,omitempty"`
}

func New() Http {
	return nil
}

type Http interface {
	HttpRaw(*HttpRequest) *HttpResponse
	HttpRawBatch([]*HttpRequest) []*HttpResponse

	HttpQuickGet(url string, headers map[string]string) ([]byte, error)
	HttpQuickPost(url string, headers map[string]string, data []byte) ([]byte, error)
	HttpFormPost(url string, headers map[string]string, data []byte) ([]byte, error)

	HttpJsonGet(url string, headers map[string]string) ([]byte, error)
	HttpJsonPost(url string, headers map[string]string, data []byte) ([]byte, error)
}
