package server

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) userAPI(rg *gin.RouterGroup) {
	rg.GET("/", s.X(s.userList))
	rg.GET("/:id", s.X(s.userGet))
	rg.POST("/:id", s.X(s.userMessage))
}

func (s *Server) userMessage(ctx httpx.Request) {

	out, err := ioutil.ReadAll(ctx.Http.Request.Body)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	_, err = s.cUser.Message(ctx.Session, ctx.MustParam("id"), string(out))
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) userGet(ctx httpx.Request) {
	resp, err := s.cUser.Get(
		ctx.Session,
		ctx.MustParam("user"),
	)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) userList(ctx httpx.Request) {
	resp, err := s.cUser.List(
		ctx.Session,
		nil,
	)

	httpx.WriteJSON(ctx.Http, resp, err)
}
