package evgoja

import (
	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
)

func (g *GojaInstance) temphiaBindings(r *goja.Runtime, module *goja.Object) {

	o := module.Get("exports").(*goja.Object)

	o.Set("Log", g.lastBinding.Log)
	o.Set("LoadFileAsString", func(file string) *goja.Promise {
		p, resolve, reject := r.NewPromise()

		go func() {
			data, _, err := g.lastBinding.GetFileWithMeta(file)
			g.evLoop.RunOnLoop(func(r *goja.Runtime) {
				if err != nil {
					reject(err.Error())
					return
				}

				resolve(string(data))
			})
		}()

		return p
	})

	o.Set("ListResources", func() *goja.Promise {
		return g.withPromise(func() (any, error) {
			return g.lastBinding.ListResources()
		})
	})

	o.Set("GetResource", func(name string) *goja.Promise {
		return g.withPromise(func() (any, error) {
			return g.lastBinding.GetResource(name)
		})
	})

	o.Set("InLinks", func() *goja.Promise {
		return g.withPromise(func() (any, error) {
			return g.lastBinding.InLinks()
		})
	})

	o.Set("OutLinks", func() *goja.Promise {
		return g.withPromise(func() (any, error) {
			return g.lastBinding.InLinks()
		})
	})

	o.Set("LinkExecute", func(name, method string, data *goja.Object) *goja.Promise {

		p, resolve, reject := g.rt.NewPromise()

		out, err := data.MarshalJSON()
		if err != nil {
			g.evLoop.RunOnLoop(func(r *goja.Runtime) {
				reject(err.Error())
			})
			return p
		}

		go func() {
			lazydata, err := g.lastBinding.LinkExec(name, method, lazydata.NewJsonData(out))
			if err != nil {
				g.evLoop.RunOnLoop(func(r *goja.Runtime) {
					reject(err.Error())
				})
				return
			}

			var i any

			err = lazydata.AsObject(&i)
			if err != nil {
				g.evLoop.RunOnLoop(func(r *goja.Runtime) {
					reject(err.Error())
				})
				return
			}

			g.evLoop.RunOnLoop(func(r *goja.Runtime) {
				resolve(i)
			})

		}()

		return p

	})

	o.Set("LinkEmit", func(name, method string, data *goja.Object) any {
		out, err := data.MarshalJSON()
		if err != nil {
			return err
		}

		return g.lastBinding.LinkExecEmit(name, method, lazydata.NewJsonData(out))
	})

	modConstructor := func(call goja.ConstructorCall) *goja.Object {

		name := call.Argument(0).ToString()
		opts := call.Argument(1).ToObject(r)

		optout, err := opts.MarshalJSON()
		if err != nil {
			panic(err)
		}

		mid, err := g.lastBinding.NewModule(name.String(), lazydata.NewJsonData(optout))
		if err != nil {
			panic(err)
		}

		call.This.Set("ExecuteAsync", func(method string, val *goja.Object) *goja.Promise {
			p, resolve, reject := r.NewPromise()

			out, err := val.MarshalJSON()
			if err != nil {
				reject(err.Error())
				return p
			}

			go func() {
				resp, err := g.lastBinding.ModuleExec(mid, method, lazydata.NewJsonData(out))
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

		call.This.Set("Execute", func(method string, val *goja.Object) (any, any) {

			out, err := val.MarshalJSON()
			if err != nil {
				return nil, err
			}

			resp, err := g.lastBinding.ModuleExec(mid, method, lazydata.NewJsonData(out))
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

		return nil
	}

	o.Set("Module", modConstructor)

	o.Set("ModuleTicket", func(name string, opts *goja.Object) (any, any) {
		out, err := opts.MarshalJSON()
		if err != nil {
			return nil, err
		}
		return g.lastBinding.ModuleTicket(name, lazydata.NewJsonData(out))
	})

	o.Set("ForkExecute", func(method string, opts *goja.Object) any {
		out, err := opts.MarshalJSON()
		if err != nil {
			return err
		}
		return g.lastBinding.ForkExec(method, out)
	})

	o.Set("HttpFetch", g.lastBinding.HttpFetch)
}

func (g *GojaInstance) withPromise(fn func() (any, error)) *goja.Promise {
	p, resolve, reject := g.rt.NewPromise()

	go func() {
		resp, err := fn()

		g.evLoop.RunOnLoop(func(r *goja.Runtime) {
			if err != nil {
				reject(err.Error())
				return
			}

			resolve(resp)
		})

	}()

	return p
}
