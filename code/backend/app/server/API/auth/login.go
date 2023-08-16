package authserver

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/authed"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz/httpx"
)

func (s *Auth) authLoginNext(c *gin.Context) {
	opts := authed.LoginNextRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	resp, err := s.cAuth.LoginNext(opts)
	httpx.WriteJSON(c, resp, err)

}

func (s *Auth) authLoginSubmit(c *gin.Context) {
	opts := authed.LoginSubmitRequest{}
	err := c.BindJSON(&opts)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	resp, err := s.cAuth.LoginSubmit(opts)
	httpx.WriteJSON(c, resp, err)
}
