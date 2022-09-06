package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/vmodels"
)

func (s *Server) operatorAPI(rg *gin.RouterGroup) {

	rg.Use(s.opsX)

	rg.POST("/login", s.operatorLogin)
	rg.GET("/stats", s.operatorStats)

	rg.GET("/tenant", s.operatorListTenant)
	rg.POST("/tenant", s.operatorAddTenant)
	rg.PATCH("/tenant/:slug", s.operatorUpdateTenant)
	rg.DELETE("/tenant:slug", s.operatorDeleteTenant)
}

func (s *Server) opsX(c *gin.Context) {
	if strings.HasSuffix(c.Request.URL.Path, "/login") {
		c.Next()
	}

	token := c.Request.Header.Get("Authorization")
	opclaim, err := s.signer.ParseOperator(token)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	c.Set("op_claim", opclaim)
	c.Next()
}

func (s *Server) operatorAddTenant(c *gin.Context) {
	data := &vmodels.NewTenant{}
	err := c.BindJSON(data)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	err = s.cOperator.AddTenant(data)
	httpx.WriteFinal(c, err)
}

func (s *Server) operatorUpdateTenant(c *gin.Context) {
	data := make(map[string]interface{})
	err := c.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	delete(data, "slug")

	err = s.cOperator.UpdateTenant(c.Param("slug"), data)
	httpx.WriteFinal(c, err)

}

func (s *Server) operatorListTenant(c *gin.Context) {
	data, err := s.cOperator.ListTenant()
	httpx.WriteJSON(c, data, err)
}

func (s *Server) operatorDeleteTenant(c *gin.Context) {
	slug := c.Param("slug")

	if slug == "" {
		httpx.WriteErr(c, "empty tenant slug")
		return
	}

	s.cOperator.DeleteTenant(slug)
}

func (s *Server) operatorStats(c *gin.Context) {

}

func (s *Server) operatorLogin(c *gin.Context) {
	data := &vmodels.OperatorLoginReq{}

	err := c.BindJSON(data)
	if err != nil {
		httpx.WriteErr(c, err.Error())
		return
	}

	resp, err := s.cOperator.Login(data)
	httpx.WriteJSON(c, resp, err)
}
