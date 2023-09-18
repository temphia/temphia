package api_server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/services/sockd/transports"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (s *Server) DataWSAPI(rg *gin.RouterGroup) {
	rg.GET("/", s.sockdDataWS)

}

func (s *Server) WSAPI(rg *gin.RouterGroup) {
	rg.GET("/data", s.sockdDataWS)
	rg.GET("/plug", s.enginePlugWS)
	rg.GET("/self", s.selfUserWS)
	rg.GET("/mqtt", func(ctx *gin.Context) {})
}

func (s *Server) sockdDataWS(ctx *gin.Context) {
	if !ctx.IsWebsocket() {
		return
	}

	dclaim, err := s.signer.ParseData(ctx.Param("tenant_id"), ctx.Query("token"))
	if err != nil {
		httpx.UnAuthorized(ctx)
		return
	}

	conn, err := transports.NewConnWS(ctx, s.idNode.Generate().Int64())
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cSockd.AddData(sockd.DataConnOptions{
		TenantId:  dclaim.TenantId,
		UserId:    dclaim.UserID,
		DynSource: dclaim.DataSource,
		DynGroup:  dclaim.DataGroup,
		Conn:      conn,
	})

	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

}

func (s *Server) enginePlugWS(ctx *gin.Context) {

	if !ctx.IsWebsocket() {
		return
	}

	tenantId := ctx.Param("tenant_id")

	dclaim, err := s.signer.ParseSockdTkt(tenantId, ctx.Query("ticket"))
	if err != nil {
		httpx.UnAuthorized(ctx)
		return
	}

	conn, err := transports.NewConnWS(ctx, s.idNode.Generate().Int64())
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cSockd.AddPlugConn(sockd.PlugConnOptions{
		TenantId: tenantId,
		UserId:   dclaim.UserId,
		GroupId:  "",
		DeviceId: dclaim.DeviceId,
		Plug:     "",
		Conn:     conn,
		Room:     dclaim.Room,
	})

	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

}

func (s *Server) selfUserWS(ctx *gin.Context) {

	sclaim, err := s.signer.ParseSession(ctx.Param("tenant_id"), ctx.Query("token"))
	if err != nil {
		httpx.UnAuthorized(ctx)
		return
	}

	conn, err := transports.NewConnWS(ctx, s.idNode.Generate().Int64())
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
