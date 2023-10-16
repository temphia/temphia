package evgoja

import (
	"io"
	"net/http"

	"github.com/dop251/goja"
)

func (g *EvGoja) webRawXecute(rw http.ResponseWriter, req *http.Request) {

	g.evLoop.RunOnLoop(func(r *goja.Runtime) {

		obj := g.buildRequestContext(rw, req, r)

		var entry func(ctx *goja.Object)
		eval := r.Get("http_handle")
		if eval == nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		entry(obj)
	})

}

func (g *EvGoja) buildRequestContext(rw http.ResponseWriter, req *http.Request, r *goja.Runtime) *goja.Object {
	obj := r.NewObject()

	obj.Set("dataAsBytes", func() goja.ArrayBuffer {
		out, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		return r.NewArrayBuffer(out)
	})

	obj.Set("dataAsObject", func() *goja.Object {
		out, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		dobj := r.NewObject()
		err = dobj.UnmarshalJSON(out)
		if err != nil {
			panic(err)
		}

		return dobj
	})

	obj.Set("setStatus", func(status int) {
		rw.WriteHeader(status)
	})

	obj.Set("setHeader", func(key, value string) {
		rw.Header().Set(key, value)
	})

	obj.Set("write", func(data []byte) error {
		_, err := rw.Write(data)
		return err
	})

	obj.Set("getQueryParam", func(name string) string {
		return req.URL.Query().Get(name)
	})

	obj.Set("getRequestHeader", func(name string) string {
		return req.Header.Get(name)
	})

	obj.Set("getRequestHeaders", func(name string) any {
		return req.Header
	})

	obj.Set("path", req.URL.String())

	return obj

}
