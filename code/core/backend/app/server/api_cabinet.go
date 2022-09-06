package server

import "github.com/gin-gonic/gin"

func (s *Server) cabinetAPI(rg *gin.RouterGroup) {

	/*

		apiv1.GET("/cabinet_sources", r.Authed(r.ListCabinetSources))
		blobapi.GET("/", r.Authed(r.ListRootFolder))
		blobapi.GET("/:folder", r.Authed(r.ListFolder))
		blobapi.POST("/:folder", r.Authed(r.NewFolder))
		blobapi.GET("/:folder/file/:fname", r.Authed(r.GetFile))
		blobapi.POST("/:folder/file/:fname", r.Authed(r.UploadFile))
		blobapi.DELETE("/:folder/file/:fname", r.Authed(r.DeleteFile))

		blobapi.GET("/:folder/preview/:fname", r.Authed(r.GetFilePreview))
		blobapi.POST("/:folder/ticket", r.Authed(r.GetFolderTicket))

					partCab := apiv1.Group("/ticket_cabinet/:ticket")
			partCab.GET("/", r.TicketCabinetList)
			partCab.GET("/:file", r.TicketCabinetFile)
			partCab.GET("/preview/:file", r.TicketCabinetPreviewFile)
			partCab.POST("/:file", r.TicketCabinetUpload)


			ftkt.GET("/", s.routes.FolderTktList)
		ftkt.GET("/:name", s.routes.FolderTktFile)
		ftkt.GET("/:name/preview", s.routes.FolderTktPreview)
		ftkt.POST("/:name", s.routes.FolderTktUpload)
		ftkt.DELETE("/:name", s.routes.FolderTktDelete)


	*/

}
