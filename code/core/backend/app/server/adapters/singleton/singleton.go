package singleton

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

type singleton struct {
	app xtypes.App
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	return &singleton{
		app: opts.App,
	}, nil
}

func (s *singleton) Handle(ctx httpx.Context) {

}
