package net

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
)

func httpRaw(client *http.Client, request *bindx.HttpRequest) *bindx.HttpResponse {
	resp, err := httpRequest(
		client,
		request.Path,
		request.Method,
		request.Headers,
		bytes.NewReader(request.Body),
	)

	if err != nil {
		return &bindx.HttpResponse{
			SatusCode: 500,
			Headers:   map[string][]string{"binding_intercepted_err": {"true"}},
			Body:      []byte(err.Error()),
		}
	}
	return resp
}

func httpRequest(client *http.Client, u, method string, header map[string]string, body io.Reader) (*bindx.HttpResponse, error) {
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	for headkey, headval := range header {
		req.Header.Set(headkey, headval)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &bindx.HttpResponse{
		Headers:   res.Header,
		Body:      bytes,
		SatusCode: res.StatusCode,
	}
	return r, nil
}
