package common

import (
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/engine/executors/javascript1/goja"
	"github.com/temphia/temphia/code/core/backend/engine/executors/wasm1/wazero"

	// repo providers
	_ "github.com/temphia/temphia/code/core/backend/services/repohub/rprovider/embed"
	_ "github.com/temphia/temphia/code/core/backend/services/repohub/rprovider/gitlab"
	_ "github.com/temphia/temphia/code/core/backend/services/repohub/rprovider/local"

	// db providers
	_ "github.com/temphia/temphia/code/core/backend/stores/upper/vendors/postgres"
	_ "github.com/temphia/temphia/code/core/backend/stores/upper/vendors/ql"
	_ "github.com/temphia/temphia/code/core/backend/stores/upper/vendors/sqlite"

	// file providers
	_ "github.com/temphia/temphia/code/core/backend/stores/cabinet/native"
)

func init() {
	registry.SetExecutor("goja", (goja.NewBuilder))
	registry.SetExecutor("wasm1", (wazero.NewBuilder))
}
