package javascript

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/extension"
	"github.com/temphia/temphia/code/executors/re"
	"github.com/temphia/temphia/code/executors/re/rtypes"
)

const EXECUTOR_TYPE = ""

func BuilderFunc(app xtypes.App, handle extension.Handle) (any, error) {

	// fixme => check if python is available

	handle.SetExecutorBuilder(
		EXECUTOR_TYPE,
		re.NewBuilder(EXECUTOR_TYPE, `bash setup.sh`, BootstrapProject),
	)

	return nil, nil
}

func BootstrapProject(ctx rtypes.BootstrapContext) error {

	return nil
}
