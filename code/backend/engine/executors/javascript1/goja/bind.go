package goja

import (
	"encoding/json"

	"github.com/dop251/goja"

	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx/ticket"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (g *Goja) qbind(name string, fn any) {
	err := g.runtime.Set(name, fn)
	if err != nil {
		panic(err)
	}

}

func resp(err error) any {
	if err != nil {
		return err.Error()
	}
	return nil
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

	if pbind := g.binder.PlugKVBindingsGet(); pbind != nil {
		g.qbind("_pkv_set", func(txid uint32, key string, value string, opts *store.SetOptions) any {
			return resp(pbind.Set(txid, key, value, opts))
		})

		g.qbind("_pkv_update", func(txid uint32, key string, value string, opts *store.UpdateOptions) any {
			return resp(pbind.Update(txid, key, value, opts))
		})

		g.qbind("_pkv_get", func(txid uint32, key string) (any, any) {
			r, err := pbind.Get(txid, key)
			if err != nil {
				return nil, err.Error()
			}

			return r, nil
		})

		g.qbind("_pkv_del", func(txid uint32, key string) any {
			return resp(pbind.Del(txid, key))
		})
		g.qbind("_pkv_batch_del", func(txid uint32, keys []string) any {
			return resp(pbind.DelBatch(txid, keys))
		})
		g.qbind("_pkv_query", func(txid uint32, query *store.PkvQuery) (any, any) {
			r, err := pbind.Query(txid, query)
			if err != nil {
				return nil, err.Error()
			}

			return r, nil
		})
		g.qbind("_pkv_new_txn", func() (uint32, any) {
			txid, err := pbind.NewTxn()
			if err != nil {
				return 0, err.Error()
			}

			return txid, nil
		})
		g.qbind("_pkv_rollback", func(txid uint32) any {
			return resp(pbind.RollBack(txid))
		})
		g.qbind("_pkv_commit", func(txid uint32) any {
			return resp(pbind.Commit(txid))
		})

		g.qbind("_pkv_ticket", func(opts *ticket.PlugState) (any, any) {
			return pbind.Ticket(opts)
		})

	}

	if sbind := g.binder.SockdBindingsGet(); sbind != nil {

		g.qbind("_sd_send_direct", func(room string, connId int64, payload []byte) any {
			return resp(sbind.SendDirect(room, connId, payload))
		})

		g.qbind("_sd_send_direct_batch", func(room string, connIds []int64, payload []byte) any {
			return resp(sbind.SendDirectBatch(room, connIds, payload))
		})

		g.qbind("_sd_send_broadcast", func(room string, ignores []int64, payload []byte) any {
			return resp(sbind.SendBroadcast(room, ignores, payload))
		})

		g.qbind("_sd_send_tagged", func(room string, tags []string, ignoreConns []int64, payload []byte) any {
			return resp(sbind.SendTagged(room, tags, ignoreConns, payload))
		})

		g.qbind("_sd_ticket", func(room string, opts *ticket.SockdRoom) (any, any) {
			return sbind.Ticket(room, opts)
		})

	}

	if cbind := g.binder.CabinetBindingsGet(); cbind != nil {
		g.qbind("_cab_add_file", func(folder, file string, payload []byte) any {
			return resp(cbind.AddFile(folder, file, payload))
		})

		g.qbind("_cab_list_folder", func(folder, file string, payload []byte) (any, any) {
			fr, err := cbind.ListFolder(folder)
			if err != nil {
				return nil, err.Error()
			}
			return fr, nil
		})

		g.qbind("_cab_get_file_str", func(folder, file string, payload []byte) (any, any) {
			out, err := cbind.GetFile(folder, file)
			if err != nil {
				return nil, err.Error()
			}

			return string(out), nil
		})

		// fixme => actually TEST IT
		g.qbind("_cab_get_file", func(folder, file string, payload []byte) (*goja.ArrayBuffer, any) {
			out, err := cbind.GetFile(folder, file)
			if err != nil {
				return nil, err.Error()
			}
			arr := g.runtime.NewArrayBuffer(out)
			return &arr, nil
		})

		g.qbind("_cab_del_file", func(folder, file string) any {
			return resp(cbind.DeleteFile(folder, file))
		})

		g.qbind("_cab_generate_ticket", func(folder string, opts *ticket.CabinetFolder) (any, any) {
			tok, err := cbind.Ticket(folder, opts)
			if err != nil {
				return nil, err.Error()
			}
			return tok, nil
		})

	}

	if ibind := g.binder.InvokerGet(); ibind != nil {
		g.qbind("_invoker_name", func() any {
			return ibind.Name()
		})

		g.qbind("_invoker_exec_method", func(method, path string, data goja.Value) (any, any) {

			resp, err := ibind.ExecMethod(method, path, lazydata.NewGojaData(g.runtime, data))
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
			return ibind.ContextUser()
		})

		g.qbind("_invoker_context_user_info", func() (any, any) {
			u, err := ibind.ContextUserInfo()
			if err != nil {
				return nil, err.Error()
			}
			return u, nil
		})

		g.qbind("_invoker_context_user_message", func(opts *bindx.UserMessage) any {
			err := ibind.ContextUserMessage(opts)
			if err != nil {
				return err.Error()
			}
			return nil
		})

	}

	if nb := g.binder.NetGet(); nb != nil {

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
			resp, err := self.SelfListResources()
			if err != nil {
				return nil, err.Error()
			}
			return resp, nil
		})

		g.qbind("_self_get_resource", func(name string) (any, any) {
			resp, err := self.SelfGetResource(name)
			if err != nil {
				return nil, err.Error()
			}

			return resp, nil
		})

		g.qbind("_self_inlinks", func() (any, any) {
			resp, err := self.SelfInLinks()
			if err != nil {
				return nil, err.Error()
			}

			return resp, nil
		})

		g.qbind("_self_outlinks", func() (any, any) {
			resp, err := self.SelfOutLinks()
			if err != nil {
				return nil, err.Error()
			}
			return resp, nil
		})

		g.qbind("_self_module_execute", func(name, method, path string, data goja.Value) (any, any) {
			resp, err := self.SelfModuleExec(name, method, path, lazydata.NewGojaData(g.runtime, data))
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
			resp, err := self.SelfLinkExec(name, method, lazydata.NewGojaData(g.runtime, data), async, detached)
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
			err := self.SelfForkExec(method, []byte(data))
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

func http1(nb bindx.Net, runtime *goja.Runtime) func(request *HTTPRequest) HTTPResponse {
	return func(request *HTTPRequest) HTTPResponse {

		var bytes []byte

		switch v := request.Body.(type) {
		case *goja.ArrayBuffer:
			bytes = v.Bytes()
		default:
			bytes, _ = json.Marshal(request.Body)
		}

		resp := nb.HttpRaw(&bindx.HttpRequest{
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
