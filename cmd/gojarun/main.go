package main

import (
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/k0kubun/pp"
)

func main() {

	evl := eventloop.NewEventLoop()

	evl.Start()

	evl.RunOnLoop(func(r *goja.Runtime) {
		r.Set("NativeModule", NewNativeModule(evl, r))
	})

	var val *goja.Promise

	evl.RunOnLoop(func(r *goja.Runtime) {
		_, err := r.RunString(`

		const dosth = async () => {
			const mod =  new NativeModule({});
			const p = await mod.execute("aa", {"mno": 22})


			return p + 9;
		};
		`)
		if err != nil {
			panic(err)
		}

		var entry func() *goja.Promise

		pp.Println(r.ExportTo(r.Get("dosth"), &entry))

		val = entry()

	})

	time.Sleep(3 * time.Second)

	pp.Println("state |>", val.State())
	pp.Println("result |>", val.Result())

	evl.Stop()

}

func NewNativeModule(evl *eventloop.EventLoop, r *goja.Runtime) func(call goja.ConstructorCall) *goja.Object {

	return func(call goja.ConstructorCall) *goja.Object {

		call.This.Set("execute", func(method string, val *goja.Object) *goja.Promise {
			ej, _ := val.MarshalJSON()

			pp.Println("EXECUTE", method, string(ej))

			p, resolve, _ := r.NewPromise()
			go func() {
				time.Sleep(1 * time.Second)

				evl.RunOnLoop(func(r *goja.Runtime) {
					pp.Println("@inside")
					resolve("HAHAHAHH")
				})
			}()

			return p
		})

		return nil
	}

}
