package evgoja

import (
	"encoding/json"
	"time"

	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
)

var _ etypes.Executor = (*Goja)(nil)

type Goja struct {
	instance *GojaInstance
	binding  bindx.Bindings
}

func NewExecutor(instance *GojaInstance) func(opts etypes.ExecutorOption) (etypes.Executor, error) {
	return func(opts etypes.ExecutorOption) (etypes.Executor, error) {
		return &Goja{
			instance: instance,
			binding:  opts.Binder,
		}, nil
	}
}

type Response struct {
	Payload any `json:"payload,omitempty"`
}

type Request struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Data any    `json:"data,omitempty"`
}

func (g *Goja) Process(ev *event.Request) (resp *event.Response, err error) {

	var val *goja.Promise

	g.instance.evLoop.RunOnLoop(func(r *goja.Runtime) {
		obj := r.NewObject()

		obj.Set("Data", func() any {
			var i any
			err := json.Unmarshal(ev.Data, &i)
			if err != nil {
				panic(err)
			}
			return i
		})

		obj.Set("DataAsBytes", func() goja.ArrayBuffer {
			return r.NewArrayBuffer(ev.Data)
		})

		obj.Set("Invoker", g.InvokerObject)

		var entry func(ctx *goja.Object) *goja.Promise
		eval := r.Get(ev.Name)
		if eval == nil {
			err = easyerr.NotFound("js action name")
			return
		}

		val = entry(obj)
	})

	if err != nil {
		return
	}

	for {
		if val.State() != goja.PromiseStatePending {
			break
		}
		time.Sleep(time.Microsecond * 200)
	}

	if val.State() == goja.PromiseStateRejected {
		return nil, easyerr.Error(val.Result().String())
	}

	out, err := val.
		Result().
		ToObject(g.instance.rt).
		MarshalJSON()
	if err != nil {
		return nil, err
	}

	return &event.Response{
		Payload: out,
	}, nil
}

func (g *Goja) InvokerObject(call goja.ConstructorCall) *goja.Object {

	invoker := g.binding.GetInvoker()

	call.This.Set("Name", invoker.Name)
	call.This.Set("UserContext", invoker.UserContext)
	call.This.Set("UserInfo", func() (any, any) {
		return invoker.UserInfo()
	})

	call.This.Set("UserMessage", func(opts *bindx.UserMessage) any {
		return invoker.UserMessage(opts)
	})

	call.This.Set("ExecMethod", func(method string, data *goja.Object) (any, any) {
		out, err := data.MarshalJSON()
		if err != nil {
			return nil, err
		}

		resp, err := invoker.ExecMethod(method, lazydata.NewJsonData(out))
		if err != nil {
			return nil, err
		}

		var i any
		err = resp.AsObject(&i)
		if err != nil {
			return nil, err
		}

		return i, nil
	})

	call.This.Set("ExecMethodAsync", func(method string, data *goja.Object) *goja.Promise {

		p, resolve, reject := g.instance.rt.NewPromise()

		out, err := data.MarshalJSON()
		if err != nil {
			reject(err.Error())
			return p
		}

		go func() {
			resp, err := invoker.ExecMethod(method, lazydata.NewJsonData(out))
			g.instance.evLoop.RunOnLoop(func(r *goja.Runtime) {
				if err != nil {
					reject(err.Error())
					return
				}
				var i any
				resolve(resp.AsObject(&i))
			})
		}()
		return p
	})

	return nil

}
