package launcher

import (
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

func (s *launcher) Handle(ctx httpx.Context) {

}
