package server

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/basic"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (s *Server) selfAPI(rg *gin.RouterGroup) {

	rg.GET("/system/cabinet", s.X(s.cabinetSources))
	rg.GET("/system/datatable", s.X(s.dtableSources))
	rg.GET("/system/adapter", s.X(s.ListRenderers))

	rg.GET("/load", s.X(s.selfGetInfo))
	rg.GET("/session", s.X(s.selfListSession))
	rg.POST("/email/change", s.X(s.selfChangeEmail))
	rg.GET("/message", s.X(s.selfListMessages))
	rg.POST("/message", s.X(s.selfModifyMessages))
	rg.POST("/issue", s.X(s.IssuePlugTkt))

	rg.GET("/user/:id", s.X(s.selfGetUserInfo))
	rg.POST("/user/:id/message", s.X(s.selfMessageUser))

}

func (s *Server) cabinetSources(ctx httpx.Request) {
	sources, err := s.cBasic.ListCabinetSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, sources, err)

}

func (s *Server) dtableSources(ctx httpx.Request) {
	sources, err := s.cBasic.ListDyndbSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, sources, err)
}

func (s *Server) selfMessageUser(ctx httpx.Request) {

	out, err := ioutil.ReadAll(ctx.Http.Request.Body)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	_, err = s.cBasic.MessageUser(ctx.Session, ctx.MustParam("id"), string(out))
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) selfGetUserInfo(ctx httpx.Request) {
	resp, err := s.cBasic.GetUserInfo(ctx.Session, ctx.MustParam("user"))
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) selfGetInfo(ctx httpx.Request) {
	resp, err := s.cBasic.GetSelfInfo(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) selfModifyMessages(ctx httpx.Request) {
	opts := &entities.ModifyMessages{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	err = s.cBasic.ModifyMessages(ctx.Session, opts)
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) selfListMessages(ctx httpx.Request) {
	opts := &entities.UserMessageReq{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := s.cBasic.ListMessages(ctx.Session, opts)
	httpx.WriteJSON(ctx.Http, resp, err)
}

// fixme => impl placeholder

func (s *Server) selfListSession(ctx httpx.Request) {

}

func (s *Server) selfChangeEmail(ctx httpx.Request) {

}

func (s *Server) IssuePlugTkt(ctx httpx.Request) {
	rdata := basic.DevIssueReq{}
	err := ctx.Http.BindJSON(&rdata)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	if rdata.Encoded {
		tkt, err := s.cBasic.DevIssueTktEncoded(ctx.Session, ctx.Http.Request.Host, rdata)
		httpx.WriteJSON(ctx.Http, tkt, err)
		return
	}

	tkt, err := s.cBasic.DevIssueTkt(ctx.Session, ctx.Http.Request.Host, rdata)
	httpx.WriteJSON(ctx.Http, tkt, err)
}

func (s *Server) ListRenderers(ctx httpx.Request) {
	resp := s.notz.ListRenderers()
	httpx.WriteJSON(ctx.Http, resp, nil)
}
