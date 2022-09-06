package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/authed"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) authAPI(rg *gin.RouterGroup) {

	rg.GET("/", s.AuthListMethods)

	rg.POST("/login/next", s.AuthLoginNext)
	rg.POST("/login/submit", s.AuthLoginSubmit)

	rg.POST("/alt/:id/generate", s.AuthGenerate)
	rg.POST("/alt/:id/next/:stage", s.AuthNext)
	rg.POST("/alt/:id/submit", s.AuthSubmit)

	rg.POST("/signup/next", s.AuthSignupSubmit)
	rg.POST("/signup/submit", s.AuthSignupSubmit)

	rg.POST("/finish", s.AuthedFinish)

	rg.POST("/reset", s.AuthReset)
	rg.POST("/reset/submit", s.AuthResetSubmit)
	rg.POST("/reset/finish", s.AuthResetFinish)

	rg.POST("/refresh", s.AuthRefresh)

	rg.GET("/auth/oauth_redirect", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("<h1> Nice </h1>"))
	})

}

func (s *Server) AuthListMethods(c *gin.Context) {
	resp, err := s.cAuth.AuthListMethods(
		c.GetHeader("Authorization"),
		c.Query("ugroup"),
	)

	httpx.WriteJSON(c, resp, err)
}

func (s *Server) AuthLoginNext(c *gin.Context) {
	opts := authed.LoginNextRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	resp, err := s.cAuth.LoginNext(opts)
	httpx.WriteJSON(c, resp, err)

}

func (s *Server) AuthLoginSubmit(c *gin.Context) {
	opts := authed.LoginSubmitRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	resp, err := s.cAuth.LoginSubmit(opts)
	httpx.WriteJSON(c, resp, err)
}

func (s *Server) AuthedFinish(c *gin.Context) {
	opts := authed.AuthFinishRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	resp, err := s.cAuth.AuthFinish(opts)
	httpx.WriteJSON(c, resp, err)

}

func (s *Server) AuthGenerate(c *gin.Context) {
	opts := authed.AuthGenerateRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	opts.Id = id

	resp, err := s.cAuth.AuthGenerate(opts)
	httpx.WriteJSON(c, resp, err)

}

func (s *Server) AuthNext(c *gin.Context) {
	stage := c.Param("stage")
	switch stage {
	case "first":
		opts := authed.AuthNextFirstRequest{}
		err := c.BindJSON(&opts)
		if err != nil {
			httpx.WriteErr(c, err.Error())
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			httpx.WriteErr(c, err.Error())
			return
		}

		opts.Id = id

		resp, err := s.cAuth.AuthNextFirst(opts)
		httpx.WriteJSON(c, resp, err)
	case "second":
		opts := authed.AuthNextSecondRequest{}
		err := c.BindJSON(&opts)
		if err != nil {
			httpx.WriteErr(c, err.Error())
			return
		}

		resp, err := s.cAuth.AuthNextSecond(opts)
		httpx.WriteJSON(c, resp, err)

	default:
		panic("Stage not found")
	}
}

func (s *Server) AuthSubmit(c *gin.Context) {
	opts := authed.AuthSubmitRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	resp, err := s.cAuth.AuthSubmit(opts)
	httpx.WriteJSON(c, resp, err)
}

func (s *Server) AuthSignupSubmit(c *gin.Context) {

}

func (s *Server) AuthSignupFinish(c *gin.Context) {

}

func (s *Server) AuthReset(c *gin.Context) {

}

func (s *Server) AuthResetSubmit(c *gin.Context) {

}

func (s *Server) AuthResetFinish(c *gin.Context) {
	opts := authed.AuthFinishRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	resp, err := s.cAuth.AuthFinish(opts)
	httpx.WriteJSON(c, resp, err)
}

func (s *Server) AuthRefresh(c *gin.Context) {
	opts := authed.RefreshReq{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	uclaim, err := s.signer.ParseUser(c.Param("tenant_id"), opts.UserToken)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}
	resp := s.cAuth.RefreshService(uclaim, opts)
	httpx.WriteJSON(c, resp, nil)
}
