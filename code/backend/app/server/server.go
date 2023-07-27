package server

import "net/http"

type Options struct {
	RootDomain   string
	RunnerDomain string
	Port         string
}

type Server struct {
	opts Options
}

func New(opts Options) *Server {
	return &Server{
		opts: opts,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}
