package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (s *Server) selfAPI(rg *gin.RouterGroup) {

	rg.GET("/system/cabinet", s.X(s.cabinetSources))
	rg.GET("/system/datatable", s.X(s.dtableSources))
	rg.GET("/system/adapter", s.X(s.ListRenderers))
	rg.GET("/system/repo", s.X(s.ListRepoSources))

	rg.GET("/load", s.X(s.selfGetInfo))
	rg.GET("/session", s.X(s.selfListSession))
	rg.POST("/email/change", s.X(s.selfChangeEmail))
	rg.GET("/message", s.X(s.selfListMessages))
	rg.POST("/message", s.X(s.selfModifyMessages))
	rg.GET("/self", s.X(s.self))
	rg.POST("/self", s.X(s.selfUpdate))

	rg.POST("/issue/folder", s.X(s.issueFolderTkt))

}

func (s *Server) cabinetSources(ctx httpx.Request) {
	sources, err := s.cBasic.ListCabinetSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, sources, err)

}

func (s *Server) dtableSources(ctx httpx.Request) {
	sources, err := s.cBasic.ListDyndbSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, sources, err)
}

func (s *Server) self(ctx httpx.Request) {
	resp, err := s.cBasic.Self(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) selfUpdate(ctx httpx.Request) {
	err := s.cBasic.SelfUpdate(ctx.Session)
	httpx.WriteJSON(ctx.Http, nil, err)
}

func (s *Server) selfGetInfo(ctx httpx.Request) {
	resp, err := s.cBasic.GetSelfInfo(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) selfModifyMessages(ctx httpx.Request) {
	opts := &entities.ModifyMessages{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	err = s.cBasic.ModifyMessages(ctx.Session, opts)
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) selfListMessages(ctx httpx.Request) {
	cursor, err := strconv.ParseInt(ctx.Http.Query("cursor"), 10, 64)
	if err != nil {
		cursor = 0
	}

	count, err := strconv.ParseInt(ctx.Http.Query("count"), 10, 64)
	if err != nil {
		count = 0
	}

	opts := &entities.UserMessageReq{
		Cursor: cursor,
		Count:  count,
		UserId: ctx.Session.UserID,
	}

	resp, err := s.cBasic.ListMessages(ctx.Session, opts)
	httpx.WriteJSON(ctx.Http, resp, err)
}

// fixme => impl placeholder

func (s *Server) selfListSession(ctx httpx.Request) {

}

func (s *Server) selfChangeEmail(ctx httpx.Request) {

}

func (s *Server) ListRenderers(ctx httpx.Request) {
	resp := s.notz.ListRenderers()
	httpx.WriteJSON(ctx.Http, resp, nil)
}

func (s *Server) ListRepoSources(ctx httpx.Request) {
	resp, err := s.cBasic.ListRepoSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) issueFolderTkt(ctx httpx.Request) {

	resp, err := s.cCabinet.NewFolderTicket(
		ctx.Session,
		ctx.Http.Param("folder"),
	)

	httpx.WriteJSON(ctx.Http, resp, err)
}
