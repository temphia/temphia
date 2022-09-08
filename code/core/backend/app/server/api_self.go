package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/basic"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (s *Server) selfAPI(rg *gin.RouterGroup) {

	rg.GET("/system/cabinet", func(ctx *gin.Context) {})
	rg.GET("/system/datatable", func(ctx *gin.Context) {})

	rg.GET("/load", func(ctx *gin.Context) {})
	rg.GET("/session", func(ctx *gin.Context) {})
	rg.POST("/email/change", func(ctx *gin.Context) {})

	rg.GET("/message", func(ctx *gin.Context) {})
	rg.POST("/message", func(ctx *gin.Context) {})
	rg.POST("/issue", s.X(s.IssuePlugTkt))

	rg.GET("/user/:id", func(ctx *gin.Context) {})
	rg.POST("/user/:id/message", func(ctx *gin.Context) {})

}

type messageUser struct {
	UserId  string `json:"user_id,omitempty"`
	Message string `json:"message,omitempty"`
}

func (s *Server) SelfMessageUser(ctx httpx.Request) {
	data := messageUser{}
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	_, err = s.cBasic.MessageUser(ctx.Session, data.UserId, data.Message)
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) SelfGetUserInfo(ctx httpx.Request) {
	resp, err := s.cBasic.GetUserInfo(ctx.Session, ctx.MustParam("user"))
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) SelfGetInfo(ctx httpx.Request) {
	resp, err := s.cBasic.GetSelfInfo(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) SelfModifyMessages(ctx httpx.Request) {
	opts := &entities.ModifyMessages{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	err = s.cBasic.ModifyMessages(ctx.Session, opts)
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) SelfListMessages(ctx httpx.Request) {
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

func (s *Server) SelfChangeEmail(ctx httpx.Request) {

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
