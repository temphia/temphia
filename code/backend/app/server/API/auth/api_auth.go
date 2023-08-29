package authserver

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/authed"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz/httpx"
)

/*

	# next
	- normal auth
	- device auth
	- app auth



*/

type Auth struct {
	cAuth    *authed.Controller
	signer   service.Signer
	tenantId string
}

func New(cAuth *authed.Controller, signer service.Signer, tenantId string) *Auth {
	return &Auth{
		cAuth:    cAuth,
		signer:   signer,
		tenantId: tenantId,
	}
}

func (s *Auth) API(rg *gin.RouterGroup) {

	rg.GET("/", s.authListMethods)

	rg.POST("/login/next", s.authLoginNext)
	rg.POST("/login/submit", s.authLoginSubmit)

	rg.POST("/alt/:id/generate", s.authGenerate)
	rg.POST("/alt/:id/next/:stage", s.authNext)
	rg.POST("/alt/:id/submit", s.authSubmit)

	rg.POST("/signup/next", s.authSignupSubmit)
	rg.POST("/signup/submit", s.authSignupSubmit)

	rg.POST("/finish", s.authedFinish)

	rg.POST("/reset", s.authReset)
	rg.POST("/reset/submit", s.authResetSubmit)
	rg.POST("/reset/finish", s.authResetFinish)

	rg.POST("/refresh", s.authRefresh)
	rg.GET("/about", s.authAbout)

	rg.GET("/auth/oauth_redirect", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("<h1> Nice </h1>"))
	})

}

func (s *Auth) authListMethods(c *gin.Context) {
	resp, err := s.cAuth.AuthListMethods(
		s.tenantId,
		c.Query("ugroup"),
	)

	httpx.WriteJSON(c, resp, err)
}

func (s *Auth) authedFinish(c *gin.Context) {
	opts := authed.AuthFinishRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	resp, err := s.cAuth.AuthFinish(opts, "FIXME device name", c.ClientIP())
	httpx.WriteJSON(c, resp, err)

}

func (s *Auth) authGenerate(c *gin.Context) {
	opts := authed.AuthGenerateRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	opts.Id = id

	resp, err := s.cAuth.AuthGenerate(opts)
	httpx.WriteJSON(c, resp, err)

}

func (s *Auth) authNext(c *gin.Context) {
	stage := c.Param("stage")
	switch stage {
	case "first":
		opts := authed.AuthNextFirstRequest{}
		err := c.BindJSON(&opts)
		if err != nil {
			httpx.WriteErr(c, err)
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			httpx.WriteErr(c, err)
			return
		}

		opts.Id = id

		resp, err := s.cAuth.AuthNextFirst(opts)
		httpx.WriteJSON(c, resp, err)
	case "second":
		opts := authed.AuthNextSecondRequest{}
		err := c.BindJSON(&opts)
		if err != nil {
			httpx.WriteErr(c, err)
			return
		}

		resp, err := s.cAuth.AuthNextSecond(opts)
		httpx.WriteJSON(c, resp, err)

	default:
		panic("Stage not found")
	}
}

func (s *Auth) authSubmit(c *gin.Context) {
	opts := authed.AuthSubmitRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	resp, err := s.cAuth.AuthSubmit(opts)
	httpx.WriteJSON(c, resp, err)
}

func (s *Auth) authSignupSubmit(c *gin.Context) {

}

func (s *Auth) authSignupFinish(c *gin.Context) {

}

func (s *Auth) authReset(c *gin.Context) {

}

func (s *Auth) authResetSubmit(c *gin.Context) {

}

func (s *Auth) authResetFinish(c *gin.Context) {
	opts := authed.AuthFinishRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	resp, err := s.cAuth.AuthFinish(opts, c.GetHeader("User-Agent"), c.ClientIP())
	httpx.WriteJSON(c, resp, err)
}

func (s *Auth) authRefresh(c *gin.Context) {
	opts := authed.RefreshReq{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	uclaim, err := s.signer.ParseUser(c.Param("tenant_id"), opts.UserToken)
	if err != nil {
		httpx.UnAuthorized(c)
		return
	}
	resp := s.cAuth.RefreshService(uclaim, opts)
	httpx.WriteJSON(c, resp, nil)
}

func (s *Auth) authAbout(c *gin.Context) {
	uclaim, err := s.signer.ParseUser(c.Param("tenant_id"), c.GetHeader("Authorization"))
	if err != nil {
		httpx.UnAuthorized(c)
		return
	}

	resp, err := s.cAuth.About(uclaim)
	httpx.WriteJSON(c, resp, err)
}
