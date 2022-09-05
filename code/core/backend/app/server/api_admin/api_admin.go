package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

type ApiAdmin struct {
	rutil  httpx.Rutil
	cAdmin *admin.Controller
}

func New(cAdmin *admin.Controller) ApiAdmin {
	return ApiAdmin{
		rutil:  httpx.Rutil{},
		cAdmin: cAdmin,
	}
}

func (a *ApiAdmin) API(apiv1 *gin.RouterGroup) {
	// bprint_api
	// user_id
	// repo_api

}
