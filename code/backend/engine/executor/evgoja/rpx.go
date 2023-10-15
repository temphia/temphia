package evgoja

import (
	"encoding/json"
	"time"

	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

func (g *EvGoja) rPXecute(ev etypes.Request) (xtypes.BeBytes, error) {

	var val *goja.Promise
	var err error

	g.evLoop.RunOnLoop(func(r *goja.Runtime) {
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

		obj.Set("Invoker", g.invokerObject)

		var entry func(ctx *goja.Object) *goja.Promise
		eval := r.Get(ev.Name)
		if eval == nil {
			err = easyerr.NotFound("js action name")
			return
		}

		val = entry(obj)
	})

	if err != nil {
		return nil, err
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
		ToObject(g.rt).
		MarshalJSON()
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (g *EvGoja) invokerObject(call goja.ConstructorCall) *goja.Object {

	invoker := g.bindx.GetInvoker(call.Arguments[0].String())

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

		p, resolve, reject := g.rt.NewPromise()

		out, err := data.MarshalJSON()
		if err != nil {
			reject(err.Error())
			return p
		}

		go func() {
			resp, err := invoker.ExecMethod(method, lazydata.NewJsonData(out))
			g.evLoop.RunOnLoop(func(r *goja.Runtime) {
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
