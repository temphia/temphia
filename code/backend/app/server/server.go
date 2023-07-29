package server

import "net/http"

type Options struct {
	RootDomain   string
	RunnerDomain string
	Port         string
}

type Server struct {
	opts     Options
	duckMode bool
}

func New(opts Options) *Server {
	return &Server{
		opts:     opts,
		duckMode: true,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if s.duckMode {
		return
	}

}

//  /z/extension/<name>
