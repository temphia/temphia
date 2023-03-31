package server

import (
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/data"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/services/sockdhub/transports"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/tidwall/gjson"
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
	rg.POST("/issue/ugroup", s.X(s.issueUgroup))
	rg.GET("/self/ws", s.sockdUserWS)

	rg.GET("/sheet/template", s.X(s.listSheetTemplates))
	rg.POST("/sheet/template", s.X(s.instanceSheetTemplate))

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

func (s *Server) instanceSheetTemplate(ctx httpx.Request) {
	req := data.QuickSheetInstance{}

	err := ctx.Http.BindJSON(&req)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	resp, err := s.cData.InstanceSheet(ctx.Session, req)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) listSheetTemplates(ctx httpx.Request) {
	resp, err := s.cData.ListSheetTemplates(ctx.Session)
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
		req.Folder,
	)

	httpx.WriteJSON(ctx.Http, gin.H{"folder_token": resp}, err)
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

func (s *Server) issueUgroup(ctx httpx.Request) {

	out, err := io.ReadAll(ctx.Http.Request.Body)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	tok, err := s.cBasic.IssueUgroup(ctx.Session, gjson.GetBytes(out, "ugroup").String())

	httpx.WriteJSON(ctx.Http, gin.H{"ugroup_token": tok}, err)
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
		httpx.UnAuthorized(ctx)
		return
	}

	conn, err := transports.NewConnWS(ctx, s.sockdConnIdGenerator.Generate().Int64())
	if err != nil {
		httpx.WriteErr(ctx, err)
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
