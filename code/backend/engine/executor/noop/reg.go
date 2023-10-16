package noop

import (
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

func init() {

	registry.SetExecutor("noop", func(app any) (etypes.ExecutorBuilder, error) {

		b := &Builder{}

		return b, nil
	})

}
