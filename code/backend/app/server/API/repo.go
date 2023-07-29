package api_server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

func (s *Server) RepoAPI(rg *gin.RouterGroup) {
	rg.GET("/:repo", s.X(s.repoList))
	rg.GET("/:repo/:group_id/:slug", s.X(s.repoGet))
	rg.GET("/:repo/:group_id/:slug/zip", s.X(s.repoGetZip))

}

func (s *Server) repoList(ctx httpx.Request) {

	sid, err := strconv.ParseInt(ctx.MustParam("repo"), 10, 64)
	if err != nil {
		return
	}

	gid := ctx.Http.Query("group_id")

	resp, err := s.cRepo.RepoSourceList(ctx.Session, gid, sid)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) repoGet(ctx httpx.Request) {

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

func (s *Server) repoGetZip(ctx httpx.Request) {
	sid, err := strconv.ParseInt(ctx.Http.Param("repo"), 10, 64)
	if err != nil {
		return
	}

	version := ctx.Http.Query("version")
	if version == "" {
		version = "current"
	}

	resp, err := s.cRepo.RepoSourceGetZip(
		ctx.Session,
		ctx.Http.Param("group_id"),
		ctx.Http.Param("slug"),
		version,
		sid)
	if err != nil {
		return
	}

	ctx.Http.DataFromReader(http.StatusOK, 0, "application/zip", resp, nil)

}
