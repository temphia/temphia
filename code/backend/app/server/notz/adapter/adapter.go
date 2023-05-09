package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Adapter struct {
	inner httpx.Adapter
	model *entities.TenantDomain
}

func New(adptr httpx.Adapter, model *entities.TenantDomain) *Adapter {
	return &Adapter{
		inner: adptr,
		model: model,
	}
}

func (d *Adapter) ServeEditorFile(file string) ([]byte, error) {
	return d.inner.ServeEditorFile(file)
}

func (d *Adapter) PreformEditorAction(name string, data []byte) (any, error) {
	return d.inner.PreformEditorAction(name, data)
}

func (d *Adapter) Handle(ctx *gin.Context) {

	d.inner.Handle(httpx.Context{
		Rid:  0,
		Http: ctx,
	})
}
