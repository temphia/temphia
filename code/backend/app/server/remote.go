package server

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/temphia/temphia/code/backend/xtypes/remote"
)

func (s *Server) HandleRemote(req *remote.Request) (*remote.Response, error) {
	var body io.Reader
	if req.Body != nil {
		body = bytes.NewReader(req.Body)
	} else {
		body = http.NoBody
	}

	u, err := url.Parse(req.URL)
	if err != nil {
		return nil, err
	}

	u.Host = "localhost"

	hreq, err := http.NewRequest(req.Method, u.String(), body)
	if err != nil {
		return nil, err
	}

	hreq.Header.Add("Is-Remote", "true")

	resp := &remote.Response{}

	s.ginEngine.ServeHTTP(resp, hreq)

	return resp, nil
}
