package server

import "github.com/gin-gonic/gin"

func (s *Server) buildRoutes() {

}

func (s *Server) zRoutes(apiv1 *gin.RouterGroup) {

	// assets_api

	// auth_api
	// self_api
	// cabinet_api
	// engine_api
	// dev_api
	// admin_api

	// operator_api

}

func (s *Server) adminAPI(apiv1 *gin.RouterGroup) {
	// admin_api
	// bprint_api
	// user_id
	// repo_api

}

/*
func (s *Server) API(apiv1 *gin.RouterGroup) {
	s.adminTenantAPI(apiv1)
	s.authAPI(apiv1)
	s.bprintAPI(apiv1)
	s.resourceAPI(apiv1)
	s.userAPI(apiv1)
	s.userSelfAPI(apiv1)
	s.plugAPI(apiv1)
	s.repoAPI(apiv1)
	s.cabinetAPI(apiv1)
	s.dtableAPI(apiv1)
	s.engineAPI(apiv1)
	s.dev(apiv1)
	s.sysAssets(apiv1)
}

*/
