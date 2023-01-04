package notz

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type DomainInstance struct {
	adapter httpx.Adapter
	model   *entities.TenantDomain
}

func (d *DomainInstance) serveEditorFile(file string) ([]byte, error) {
	return d.adapter.ServeEditorFile(file)
}

func (d *DomainInstance) preformEditorAction(name string, data []byte) (any, error) {
	return d.adapter.PreformEditorAction(name, data)
}

func (d *DomainInstance) middleWare(ctx *gin.Context) error {
	return nil
}

func (d *DomainInstance) handle(ctx *gin.Context) {
	pp.Println("@serve domain instance")

	d.adapter.Handle(httpx.Context{
		Rid:  d.model.Id,
		Http: ctx,
	})
}
