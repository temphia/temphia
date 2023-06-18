package goja

import (
	"encoding/json"

	"github.com/dop251/goja"

	"github.com/temphia/temphia/code/backend/engine/modules/http"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

func (g *Goja) qbind(name string, fn any) {
	err := g.runtime.Set(name, fn)
	if err != nil {
		panic(err)
	}

}

func (g *Goja) bind() {

	g.runtime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	{
		// core
		g.qbind("_log", g.binder.Log)
		g.qbind("_lazy_log", g.binder.LazyLog)
		g.qbind("_sleep", g.binder.Sleep)
		g.qbind("_get_self_file_as_str", func(file string) (any, any) {
			data, _, err := g.binder.GetFileWithMeta(file)
			if err != nil {
				return nil, err.Error()
			}
			return string(data), nil
		})
	}

	if ibind := g.binder.InvokerGet(); ibind != nil {
		g.qbind("_invoker_name", func() any {
			return ibind.Name()
		})

		g.qbind("_invoker_exec_method", func(method string, data goja.Value) (any, any) {

			resp, err := ibind.ExecMethod(method, lazydata.NewGojaData(g.runtime, data))
			if err != nil {
				return nil, err.Error()
			}
			var i any

			err = resp.AsObject(&i)
			if err != nil {
				return nil, err.Error()
			}

			return i, nil
		})

		g.qbind("_invoker_context_user", func() any {
			return ibind.UserContext()
		})

		g.qbind("_invoker_context_user_info", func() (any, any) {
			u, err := ibind.UserInfo()
			if err != nil {
				return nil, err.Error()
			}
			return u, nil
		})

		g.qbind("_invoker_context_user_message", func(opts *bindx.UserMessage) any {
			err := ibind.UserMessage(opts)
			if err != nil {
				return err.Error()
			}
			return nil
		})

	}

	if nb := http.New(); nb != nil {

		hfunc := http1(nb, g.runtime)
		g.qbind("_http1", func(method string, url string, headers map[string]string, body any) (int, any, any) {

			resp := hfunc(&HTTPRequest{
				Method:  method,
				Path:    url,
				Headers: headers,
				Body:    body,
			})

			return resp.SatusCode, resp.Headers, resp.Body
		})

	}

	if self := g.binder.SelfBindingsGet(); self != nil {
		g.qbind("_self_list_resource", func() (any, any) {
			resp, err := self.ListResources()
			if err != nil {
				return nil, err.Error()
			}
			return resp, nil
		})

		g.qbind("_self_get_resource", func(name string) (any, any) {
			resp, err := self.GetResource(name)
			if err != nil {
				return nil, err.Error()
			}

			return resp, nil
		})

		g.qbind("_self_inlinks", func() (any, any) {
			resp, err := self.InLinks()
			if err != nil {
				return nil, err.Error()
			}

			return resp, nil
		})

		g.qbind("_self_outlinks", func() (any, any) {
			resp, err := self.OutLinks()
			if err != nil {
				return nil, err.Error()
			}
			return resp, nil
		})

		g.qbind("_self_new_module", func(name string, data goja.Value) (any, any) {
			return self.NewModule(name, lazydata.NewGojaData(g.runtime, data))
		})

		g.qbind("_self_module_ticket", func(name string, data goja.Value) (any, any) {
			return self.ModuleTicket(name, lazydata.NewGojaData(g.runtime, data))
		})

		g.qbind("_self_module_exec", func(mid int32, name string, data goja.Value) (any, any) {
			resp, err := self.ModuleExec(mid, name, lazydata.NewGojaData(g.runtime, data))
			if err != nil {
				return nil, err.Error()
			}

			var i any

			err = resp.AsObject(&i)
			if err != nil {
				return nil, err.Error()
			}

			return i, nil
		})

		g.qbind("_self_link_execute", func(name, method, path string, data goja.Value, async, detached bool) (any, any) {
			resp, err := self.LinkExec(name, method, lazydata.NewGojaData(g.runtime, data))
			if err != nil {
				return nil, err.Error()
			}

			var i any

			err = resp.AsObject(&i)
			if err != nil {
				return nil, err.Error()
			}

			return i, nil
		})

		g.qbind("_self_fork_execute", func(method string, data string) any {
			err := self.ForkExec(method, []byte(data))
			if err != nil {
				return err.Error()
			}
			return nil
		})

	}

}

type HTTPRequest struct {
	Method  string            `json:"method,omitempty"`
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    any               `json:"body,omitempty"`
}

type HTTPResponse struct {
	SatusCode int                 `json:"status_code,omitempty"`
	Headers   map[string][]string `json:"headers,omitempty"`
	Body      any                 `json:"body,omitempty"`
}

func http1(nb http.Http, runtime *goja.Runtime) func(request *HTTPRequest) HTTPResponse {
	return func(request *HTTPRequest) HTTPResponse {

		var bytes []byte

		switch v := request.Body.(type) {
		case *goja.ArrayBuffer:
			bytes = v.Bytes()
		default:
			bytes, _ = json.Marshal(request.Body)
		}

		resp := nb.HttpRaw(&http.HttpRequest{
			Method:  request.Method,
			Path:    request.Path,
			Headers: request.Headers,
			Body:    bytes,
		})

		return HTTPResponse{
			SatusCode: resp.SatusCode,
			Headers:   resp.Headers,
			Body:      runtime.NewArrayBuffer(resp.Body),
		}
	}
}
