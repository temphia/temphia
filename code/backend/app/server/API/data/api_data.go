package apidata

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/app/server/middleware"
	"github.com/temphia/temphia/code/backend/controllers/data"
)

type Data struct {
	middleware *middleware.Middleware
	cData      *data.Controller
}

func New(middleware *middleware.Middleware, cData *data.Controller) *Data {
	return &Data{
		middleware: middleware,
		cData:      cData,
	}
}

func (s *Data) API(rg *gin.RouterGroup) {

	rg.GET("/", s.dx(s.loadGroup))

	rg.POST("/table/:tid/row", s.dx(s.newRow))
	rg.GET("/table/:tid/row/:id", s.dx(s.getRow))
	rg.POST("/table/:tid/row/:id", s.dx(s.updateRow))
	rg.DELETE("/table/:tid/row/:id", s.dx(s.deleteRow))

	rg.POST("/table/:tid/load", s.dx(s.loadTable))
	rg.POST("/table/:tid/simple_query", s.dx(s.simpleQuery))
	rg.POST("/table/:tid/fts_query", s.dx(s.FTSQuery))
	rg.POST("/table/:tid/ref_load", s.dx(s.refLoad))
	rg.POST("/table/:tid/ref_resolve", s.dx(s.refResolve))
	rg.POST("/table/:tid/rev_ref_load", s.dx(s.reverseRefLoad))
	rg.GET("/table/:tid/activity/:row_id", s.dx(s.listActivity))
	rg.POST("/table/:tid/activity/:row_id", s.dx(s.commentRow))

	s.dataSheetAPI(rg.Group("/sheet"))

	rg.POST("/utils/user", s.dx(s.listDataUsers))

}

func (s *Data) dataSheetAPI(rg *gin.RouterGroup) {
	rg.POST("/export", s.dx(s.export))
	rg.POST("/list", s.dx(s.listSheetGroup))
	rg.POST("/:id/load", s.dx(s.loadSheet))
	rg.POST("/:id/search", s.dx(s.searchSheet))
	rg.POST("/:id/query", s.dx(s.querySheet))
	rg.POST("/:id/ref_query", s.dx(s.refSheet))

	rg.GET("/", s.dx(s.listSheet))
	rg.GET("/:id", s.dx(s.getSheet))
	rg.POST("/", s.dx(s.newSheet))
	rg.POST("/:id", s.dx(s.updateSheet))
	rg.DELETE("/:id", s.dx(s.deleteSheet))

	rg.GET("/:id/column", s.dx(s.listSheetColumn))
	rg.POST("/:id/column", s.dx(s.newSheetColumn))
	rg.GET("/:id/column/:cid", s.dx(s.getSheetColumn))
	rg.POST("/:id/column/:cid", s.dx(s.updateSheetColumn))
	rg.DELETE("/:id/column/:cid", s.dx(s.deleteSheetColumn))

	rg.POST("/:id/row_cell", s.dx(s.NewRowWithCell))
	rg.POST("/:id/row_cell/:rid", s.dx(s.UpdateRowWithCell))
	rg.GET("/:id/row_cell/:rid", s.dx(s.GetRowWithCell))
	rg.DELETE("/:id/row_cell/:rid", s.dx(s.DeleteRowWithCell))
	rg.GET("/:id/relation/:rid/ref/:refsheet/column/:refcol", s.dx(s.GetRowRelations))

	rg.GET("/:id/history/:rid", s.dx(s.getRowHistory))

}
