package api_server

import "github.com/gin-gonic/gin"

func (s *Server) EngineAPI(rg *gin.RouterGroup) {
	s.engineAPI.EngineAPI(rg)
}
