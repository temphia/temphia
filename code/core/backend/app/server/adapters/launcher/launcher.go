package launcher

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

// fixme => x-content-security-policy: frame-ancestors 'self' https://mycourses.w3schools.com;
// Referer: https://example/launcher/<ticket>

type launcher struct {
	app xtypes.App
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	return &launcher{
		app: opts.App,
	}, nil
}

func (d *launcher) ServeEditorFile(ctx *gin.Context, file string) error {
	return nil
}

func (s *launcher) Handle(ctx httpx.Context) {

}
