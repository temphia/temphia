package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/basic"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) selfDeviceAPI(rg *gin.RouterGroup) {
	rg.GET("/", s.X(s.selfListDevices))
	rg.POST("/", s.X(s.selfAddDevices))
	rg.GET("/:id", s.X(s.selfGetDevice))
	rg.DELETE("/:id", s.X(s.selfRemoveDevices))

}

func (s *Server) selfListDevices(ctx httpx.Request) {
	resp, err := s.cBasic.ListUserDevice(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) selfGetDevice(ctx httpx.Request) {
	resp, err := s.cBasic.GetUserDevice(ctx.Session, ctx.MustParamInt("id"))
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) selfRemoveDevices(ctx httpx.Request) {
	err := s.cBasic.RemoveUserDevice(ctx.Session, ctx.MustParamInt("id"))
	httpx.WriteJSON(ctx.Http, nil, err)
}

func (s *Server) selfAddDevices(ctx httpx.Request) {
	data := basic.NewUserDevice{}
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	err = s.cBasic.AddUserDevice(ctx.Session, &data)
	httpx.WriteJSON(ctx.Http, nil, err)
}
