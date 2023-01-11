package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/operator/opmodels"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
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

	token := c.GetHeader("Authorization")
	opclaim, err := s.signer.ParseOperator(token)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	c.Set("op_claim", opclaim)
	c.Next()
}

func (s *Server) operatorAddTenant(c *gin.Context) {
	data := &opmodels.NewTenant{}
	err := c.BindJSON(data)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	err = s.cOperator.AddTenant(data)
	httpx.WriteFinal(c, err)
}

func (s *Server) operatorUpdateTenant(c *gin.Context) {
	data := make(map[string]interface{})
	err := c.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(c, err)
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
		httpx.WriteErrString(c, "empty tenant slug")
		return
	}

	s.cOperator.DeleteTenant(slug)
}

func (s *Server) operatorStats(c *gin.Context) {

}

func (s *Server) operatorLogin(c *gin.Context) {
	data := &opmodels.OperatorLoginReq{}

	err := c.BindJSON(data)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	resp, err := s.cOperator.Login(data)
	httpx.WriteJSON(c, resp, err)
}
