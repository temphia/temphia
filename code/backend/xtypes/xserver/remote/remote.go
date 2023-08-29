package remote

import "net/http"

type Request struct {
	Method  string
	Headers http.Header
	URL     string
	Body    []byte
}

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (r *Response) Header() http.Header { return r.Headers }

func (r *Response) Write(data []byte) (int, error) {
	r.Body = append(r.Body, data...)

	return len(data), nil
}

func (r *Response) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
}
