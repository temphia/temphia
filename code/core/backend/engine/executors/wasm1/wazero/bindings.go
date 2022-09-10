package wazero

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func BuildTemphiaModule(runtime wazero.Runtime) (api.Module, error) {
	return runtime.NewModuleBuilder("temphia").
		ExportFunction("log", log).
		ExportFunction("sleep", sleep).
		ExportFunction("get_file_with_meta", getFileWithMeta).
		Instantiate(context.TODO(), runtime)
}
