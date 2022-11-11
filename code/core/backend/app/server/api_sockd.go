package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/sockd"
	"github.com/temphia/temphia/code/core/backend/services/sockdhub/transports"
)

func (s *Server) sockdAPI(rg *gin.RouterGroup) {

	rg.GET("/user/ws", s.sockdUserWS)
	rg.GET("/data/ws", s.sockdDataWS)
	rg.GET("/data/update", s.sockdDataUpdateWS)
	rg.GET("/room/ws", s.sockdRoomWS)
	rg.GET("/room/update", s.sockdRoomUpdateWS)
	rg.GET("/dev/room/ws", s.sockdDevWS)

}

func (s *Server) sockdUserWS(ctx *gin.Context) {

	sclaim, err := s.signer.ParseSession(ctx.Param("tenant_id"), ctx.Query("token"))
	if err != nil {
		return
	}
	conn, err := transports.NewConnWS(ctx, 0)
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

func (s *Server) sockdDataWS(ctx *gin.Context) {

}

func (s *Server) sockdDataUpdateWS(ctx *gin.Context) {

}

func (s *Server) sockdRoomWS(ctx *gin.Context) {

}

func (s *Server) sockdRoomUpdateWS(ctx *gin.Context) {

}

func (s *Server) sockdDevWS(ctx *gin.Context) {

}
