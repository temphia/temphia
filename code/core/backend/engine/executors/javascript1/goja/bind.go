package goja

import (
	"github.com/dop251/goja"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func (g *Goja) qbind(name string, fn interface{}) {
	err := g.runtime.Set(name, fn)
	if err != nil {
		panic(err)
	}

}

func resp(err error) interface{} {
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
		g.qbind("_get_self_file_as_str", func(file string) (interface{}, interface{}) {
			data, _, err := g.binder.GetFileWithMeta(file)
			if err != nil {
				return nil, err.Error()
			}
			return string(data), nil
		})
	}

	if pbind := g.binder.PlugKVBindingsGet(); pbind != nil {
		g.qbind("_pkv_set", func(txid uint32, key string, value string, opts *store.SetOptions) interface{} {
			return resp(pbind.Set(txid, key, value, opts))
		})

		g.qbind("_pkv_update", func(txid uint32, key string, value string, opts *store.UpdateOptions) interface{} {
			return resp(pbind.Update(txid, key, value, opts))
		})

		g.qbind("_pkv_get", func(txid uint32, key string) (interface{}, interface{}) {
			r, err := pbind.Get(txid, key)
			if err != nil {
				return nil, err.Error()
			}

			return r, nil
		})

		g.qbind("_pkv_del", func(txid uint32, key string) interface{} {
			return resp(pbind.Del(txid, key))
		})
		g.qbind("_pkv_batch_del", func(txid uint32, keys []string) interface{} {
			return resp(pbind.DelBatch(txid, keys))
		})
		g.qbind("_pkv_query", func(txid uint32, query *store.PkvQuery) (interface{}, interface{}) {
			r, err := pbind.Query(txid, query)
			if err != nil {
				return nil, err.Error()
			}

			return r, nil
		})
		g.qbind("_pkv_new_txn", func() (uint32, interface{}) {
			txid, err := pbind.NewTxn()
			if err != nil {
				return 0, err.Error()
			}

			return txid, nil
		})
		g.qbind("_pkv_rollback", func(txid uint32) interface{} {
			return resp(pbind.RollBack(txid))
		})
		g.qbind("_pkv_commit", func(txid uint32) interface{} {
			return resp(pbind.Commit(txid))
		})
	}

	/*

		if sbind := g.binder.GetSockdBindings(); sbind != nil {
			g.qbind("_sd_send_direct", func(room string, connId []string, payload []byte) interface{} {
				return resp(sbind.SendDirect(room, connId, payload))
			})

			g.qbind("_sd_send_broadcast", func(room string, payload []byte) interface{} {
				return resp(sbind.SendBroadcast(room, payload))
			})

			g.qbind("_sd_send_tagged", func(room string, tags []string, ignoreConns []string, payload []byte) interface{} {
				return resp(sbind.SendTagged(room, tags, ignoreConns, payload))
			})

			g.qbind("_sd_add_to_room", func(room string, connId string, tags []string) interface{} {
				return resp(sbind.AddToRoom(room, connId, tags))
			})

			g.qbind("_sd_kick_from_room", func(room string, connId string) interface{} {
				return resp(sbind.KickFromRoom(room, connId))
			})

			g.qbind("_sd_list_room_conns", func(room string) (interface{}, interface{}) {
				r, err := sbind.ListRoomConns(room)
				if err != nil {
					return nil, err.Error()
				}

				return r, nil
			})

			g.qbind("_sd_bann_conn", func(connId string) interface{} {
				return resp(sbind.BannConn(connId))
			})
		}

	*/

	if cbind := g.binder.CabinetBindingsGet(); cbind != nil {
		g.qbind("_cab_add_file", func(folder, file string, payload []byte) interface{} {
			return resp(cbind.AddFile(folder, file, payload))
		})

		g.qbind("_cab_list_folder", func(folder, file string, payload []byte) (interface{}, interface{}) {
			fr, err := cbind.ListFolder(folder)
			if err != nil {
				return nil, err.Error()
			}
			return fr, nil
		})

		g.qbind("_cab_get_file_str", func(folder, file string, payload []byte) (interface{}, interface{}) {
			out, err := cbind.GetFile(folder, file)
			if err != nil {
				return nil, err.Error()
			}

			return string(out), nil
		})

		// fixme => actually TEST IT
		g.qbind("_cab_get_file", func(folder, file string, payload []byte) (*goja.ArrayBuffer, interface{}) {
			out, err := cbind.GetFile(folder, file)
			if err != nil {
				return nil, err.Error()
			}
			arr := g.runtime.NewArrayBuffer(out)
			return &arr, nil
		})

		g.qbind("_cab_del_file", func(folder, file string) interface{} {
			return resp(cbind.DeleteFile(folder, file))
		})

		g.qbind("_cab_generate_ticket", func(folder string, opts *bindx.CabTicket) (interface{}, interface{}) {
			tok, err := cbind.GenerateTicket(folder, opts)
			if err != nil {
				return nil, err.Error()
			}
			return tok, nil
		})

	}

	g.qbind("_http1", func(method string, url string, headers map[string]string, body []byte) (int, interface{}, interface{}) {

		return 0, nil, nil
	})

	// if rbind := g.binder.GetResourceBindings(); rbind != nil {
	// 	g.qbind("_GetResource", func(name string) interface{} {
	// 		resp, err := rbind.GetResource(name)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return resp
	// 	})

	// 	g.qbind("_ListResource", func() interface{} {
	// 		resp, err := rbind.ListResource()
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return resp
	// 	})
	// }

	// if ibind := g.binder.GetIPCBindings(); ibind != nil {
	// 	g.qbind("_CallSlot", func(slot string, payload interface{}) interface{} {
	// 		resp, err := ibind.CallSlot(slot, payload)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return resp
	// 	})

	// 	g.qbind("_ListIncommin", func() interface{} {
	// 		resp, err := ibind.ListIncommin()
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return resp
	// 	})
	// 	g.qbind("_ListOutgoing", func() interface{} {
	// 		resp, err := ibind.ListOutgoing()
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return resp
	// 	})
	// }

	// if bind := g.binder.GetModuleBindings(); bind != nil {

	// }

	// if bind := g.binder.GetCryptoBindings(); bind != nil {

	// }

	// if bind := g.binder.GetUserBindings(); bind != nil {

	// }

}

/*

type HTTPRequest struct {
	Method  string            `json:"method,omitempty"`
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    interface{}       `json:"body,omitempty"`
}

type HTTPResponse struct {
	SatusCode int                 `json:"status_code,omitempty"`
	Headers   map[string][]string `json:"headers,omitempty"`
	Body      interface{}         `json:"body,omitempty"`
}

func http(bindings etypes.Bindings, runtime *goja.Runtime) func(request *HTTPRequest) *HTTPResponse {
	return func(request *HTTPRequest) *HTTPResponse {

		var bytes []byte

		switch v := request.Body.(type) {
		case *goja.ArrayBuffer:
			bytes = v.Bytes()
		default:
			bytes, _ = json.Marshal(request.Body)
		}

		resp := bindings.HTTPCall(etypes.HTTPRequest{
			Method:  request.Method,
			Path:    request.Path,
			Headers: request.Headers,
			Body:    bytes,
		})

		var out interface{}

		if resp.Json {
			json.Unmarshal(resp.Body, &out)
		} else {
			out = runtime.NewArrayBuffer(resp.Body)
		}

		return &HTTPResponse{
			SatusCode: resp.SatusCode,
			Headers:   resp.Headers,
			Body:      out,
		}
	}
}


*/
