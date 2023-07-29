package api_server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/services/sockd/transports"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

func (s *Server) DataWSAPI(rg *gin.RouterGroup) {
	rg.GET("/", s.sockdDataWS)
	rg.POST("/", s.sockdDataUpdateWS)
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

func (s *Server) sockdDataUpdateWS(ctx *gin.Context) {

}
