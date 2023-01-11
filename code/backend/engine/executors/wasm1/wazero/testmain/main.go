package main

import (
	"github.com/temphia/temphia/code/backend/engine/executors/wasm1/wazero"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
)

func main() {

	builder, err := wazero.NewBuilder(nil)
	handleErr(err)

	executor, err := builder.Instance(etypes.ExecutorOption{
		Binder:   &mb{},
		TenantId: "default1",
		PlugId:   "plug1",
		AgentId:  "agent1",
		File:     "code.wasm",
		ExecType: "wasm1",
		EnvVars:  map[string]any{},
	})
	handleErr(err)

	_, err = executor.Process(&event.Request{
		Id:   "111",
		Name: "action_ping",
		Data: ([]byte(`{}`)),
	})

	handleErr(err)

}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
