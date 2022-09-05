package apiadmin

import (
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
