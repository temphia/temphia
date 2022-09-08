package server

import (
	"github.com/gin-gonic/gin"
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
