package static

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type static struct {
	source store.CabinetSourced
	folder string
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {
	src := opts.App.GetDeps().Cabinet().(store.CabinetHub).GetSource(opts.Domain.ServeSource, opts.TenantId)
	if src == nil {
		return nil, easyerr.Error("serve source not found ")
	}

	return &static{
		source: src,
		folder: opts.Domain.ServeFolder,
	}, nil
}

func (s *static) Render(ctx *gin.Context) {

}

func (s *static) Handle(ctx httpx.Context) {
	out, err := s.source.GetBlob(ctx.Http.Request.Context(), s.folder, ctx.Http.Request.URL.Path)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	ctx.Http.Writer.Write(out)
}
