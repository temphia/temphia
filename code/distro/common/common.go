package common

import (
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/executors/noop"
	"github.com/temphia/temphia/code/executors/pageform"

	// core executors
	"github.com/temphia/temphia/code/backend/engine/executors/javascript1/goja"
	"github.com/temphia/temphia/code/backend/engine/executors/wasm1/wazero"
	"github.com/temphia/temphia/code/backend/engine/modules/datagroup"
	"github.com/temphia/temphia/code/backend/engine/modules/logproxy"

	// repo providers
	_ "github.com/temphia/temphia/code/backend/services/repohub/rprovider/embed"
	_ "github.com/temphia/temphia/code/backend/services/repohub/rprovider/github"
	_ "github.com/temphia/temphia/code/backend/services/repohub/rprovider/gitlab"
	_ "github.com/temphia/temphia/code/backend/services/repohub/rprovider/local"

	// db providers
	_ "github.com/temphia/temphia/code/backend/stores/upper/vendors/postgres"
	_ "github.com/temphia/temphia/code/backend/stores/upper/vendors/ql"
	_ "github.com/temphia/temphia/code/backend/stores/upper/vendors/sqlite"

	// file providers
	_ "github.com/temphia/temphia/code/backend/stores/cabinet/native"

	// domain adapters
	_ "github.com/temphia/temphia/code/backend/app/server/adapters"
)

func init() {

	xtypes.PrintBuildInfo()

	registry.SetExecutor("javascript1", goja.NewBuilder)
	registry.SetExecutor("goja", goja.NewBuilder)
	registry.SetExecutor("wasm1", wazero.NewBuilder)
	registry.SetExecutor("pageform", pageform.NewBuilder)
	registry.SetExecutor("noop", noop.NewBuilder)

	registry.SetExecModule("data_group", datagroup.NewBuilder)
	registry.SetExecModule("logproxy", logproxy.NewBuilder)

}
