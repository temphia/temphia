package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) repoAPI(rg *gin.RouterGroup) {
	rg.GET("/", s.X(s.RepoSources))
	rg.GET("/:repo", s.X(s.RepoList))
	rg.GET("/:repo/:group_id/:slug", s.X(s.RepoGet))
	rg.GET("/:repo/:group_id/:slug/:file", s.X(s.RepoGetFile))

}

func (s *Server) RepoSources(ctx httpx.Request) {
	// fixme

}

func (s *Server) RepoList(ctx httpx.Request) {

	sid, err := strconv.ParseInt(ctx.MustParam("repo"), 10, 64)
	if err != nil {
		return
	}

	gid := ctx.Http.Query("group_id")

	resp, err := s.cRepo.RepoSourceList(ctx.Session, gid, sid)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) RepoGet(ctx httpx.Request) {

	sid, err := strconv.ParseInt(ctx.Http.Param("repo"), 10, 64)
	if err != nil {
		return
	}

	resp, err := s.cRepo.RepoSourceGet(
		ctx.Session,
		ctx.Http.Param("group_id"),
		ctx.Http.Param("slug"),
		sid)

	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) RepoGetFile(ctx httpx.Request) {

	sid, err := strconv.ParseInt(ctx.MustParam("repo"), 10, 64)
	if err != nil {
		return
	}

	resp, err := s.cRepo.RepoSourceGetBlob(
		ctx.Session,
		ctx.MustParam("group_id"),
		ctx.MustParam("slug"),
		sid,
		ctx.MustParam("file"),
	)
	httpx.WriteJSON(ctx.Http, resp, err)

}
