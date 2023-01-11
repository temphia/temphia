package wasmer2

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"

	"github.com/wasmerio/wasmer-go/wasmer"
)

type wasmer2 struct {
	store        *wasmer.Store
	module       *wasmer.Module
	instance     *wasmer.Instance
	binder       bindx.Bindings
	pendingResp  []byte
	extenFns     map[string]wasmer.IntoExtern
	eventRequest *event.Request
	eventReply   *event.Response
}

func New(store *wasmer.Store, binder bindx.Bindings, code []byte) (*wasmer2, error) {

	module, err := wasmer.NewModule(store, code)
	if err != nil {
		return nil, err
	}

	wexec := &wasmer2{
		store:       store,
		module:      module,
		instance:    nil,
		binder:      binder,
		pendingResp: nil,
		extenFns:    make(map[string]wasmer.IntoExtern, 50),
	}

	return wexec, wexec.init()
}

func (w *wasmer2) init() error {

	importObject := wasmer.NewImportObject()

	memlimit, err := wasmer.NewLimits(100, 100)
	if err != nil {
		return err
	}

	memory := wasmer.NewMemory(
		w.store,
		wasmer.NewMemoryType(memlimit),
	)

	importObject.Register("env", map[string]wasmer.IntoExtern{
		"abort": wasmer.NewFunction(
			w.store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32),
				wasmer.NewValueTypes(),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"memory": memory,
	})
	w.buildBindings()

	if w.binder != nil {
		importObject.Register("temphia", w.extenFns)
	}

	instance, err := wasmer.NewInstance(w.module, importObject)
	if err != nil {
		return err
	}
	w.instance = instance

	return nil
}
