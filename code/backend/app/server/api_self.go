package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/services/sockdhub/transports"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (s *Server) selfAPI(rg *gin.RouterGroup) {

	rg.GET("/load", s.X(s.selfGetInfo))
	rg.POST("/email/change", s.X(s.selfChangeEmail))
	rg.GET("/message", s.X(s.selfListMessages))
	rg.POST("/message", s.X(s.selfModifyMessages))
	rg.GET("/self", s.X(s.self))
	rg.POST("/self", s.X(s.selfUpdate))
	rg.GET("/user/:user_id", s.X(s.selfUserGet))
	rg.POST("/user/:user_id", s.X(s.selfUserMessage))
	rg.POST("/issue/folder", s.X(s.issueFolderTkt))
	rg.POST("/issue/data", s.X(s.issueDataTkt))
	rg.GET("/self/ws", s.sockdUserWS)

	s.selfSysAPI(rg.Group("/system"))
	s.selfDeviceAPI(rg.Group("/device"))

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

func (s *Server) selfChangeEmail(ctx httpx.Request) {

}

func (s *Server) issueFolderTkt(ctx httpx.Request) {

	req := &FolderIssueRequest{}
	err := ctx.Http.BindJSON(req)
	if err != nil {
		return
	}

	resp, err := s.cCabinet.NewFolderTicket(
		ctx.Session,
		req.Source,
		ctx.Http.Param("folder"),
	)

	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) issueDataTkt(ctx httpx.Request) {

	req := &DataIssueRequest{}
	err := ctx.Http.BindJSON(req)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	resp, err := s.cData.IssueDataClaim(
		ctx.Session,
		req.Source,
		req.Group,
	)

	httpx.WriteJSON(ctx.Http, gin.H{"data_token": resp}, err)
}

func (s *Server) selfUserGet(ctx httpx.Request) {
	resp, err := s.cUser.Get(
		ctx.Session,
		ctx.MustParam("user_id"),
	)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) selfUserMessage(ctx httpx.Request) {
	var out string

	err := ctx.Http.BindJSON(&out)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	_, err = s.cUser.Message(ctx.Session, ctx.MustParam("user_id"), string(out))
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) sockdUserWS(ctx *gin.Context) {

	sclaim, err := s.signer.ParseSession(ctx.Param("tenant_id"), ctx.Query("token"))
	if err != nil {
		return
	}

	conn, err := transports.NewConnWS(ctx, s.sockdConnIdGenerator.Generate().Int64())
	if err != nil {
		return
	}

	s.cSockd.AddUserConn(sockd.UserConnOptions{
		TenantId: ctx.Param("tenant_id"),
		UserId:   sclaim.UserID,
		GroupId:  sclaim.UserGroup,
		DeviceId: sclaim.DeviceId,
		Conn:     conn,
	})

}

// models

type DataIssueRequest struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
}

type FolderIssueRequest struct {
	Source string `json:"source,omitempty"`
	Folder string `json:"folder,omitempty"`
}