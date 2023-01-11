package wasmer2

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func NewBuilder() func(opts etypes.ExecutorOption) (etypes.Executor, error) {
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	return func(opts etypes.ExecutorOption) (etypes.Executor, error) {
		out, _, err := opts.Binder.GetFileWithMeta("server.wasm")
		if err != nil {
			return nil, err
		}

		exec, err := New(store, opts.Binder, out)
		if err != nil {
			return nil, err
		}
		err = exec.init()
		if err != nil {
			return nil, err
		}
		return exec, nil
	}

}

func (w *wasmer2) Process(ereq *event.Request) (*event.Response, error) {

	entry, err := w.instance.Exports.GetFunction(ereq.Name)
	if err != nil {
		return nil, err
	}

	w.eventRequest = ereq

	_, err = entry()
	if err != nil {
		return nil, err
	}

	if w.eventReply == nil {
		return nil, easyerr.Error("Empty Reply")
	}

	return w.eventReply, nil
}
