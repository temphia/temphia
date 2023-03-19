package web2agent

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
)

type WATarget struct {
	adapter *Web2Agent
	rid     int64
	http    *gin.Context
}

func (w *WATarget) handle() {
	if w.adapter.mainHook == nil {
		w.error("hook not loaded")
		return
	}

	if !w.adapter.intOk {
		w.error(w.adapter.initError)
		return
	}

	path := w.http.Request.URL.Path

	if w.http.Request.Method == http.MethodGet {
		if tfile, ok := w.adapter.state.templates[path]; ok {
			w.serveTemplate(tfile)
			return
		}
	}

	if action, ok := w.adapter.state.routes[fmt.Sprintf("%s|%s", path, w.http.Request.Method)]; ok {
		w.serveAction(action)
		return
	}

	if action, ok := w.adapter.state.routes[path]; ok {
		w.serveAction(action)
		return
	}
}

func (w *WATarget) error(estr string) {
	w.http.Data(http.StatusBadRequest, "text/html", []byte(estr))
}

// target methods

func (w *WATarget) Type() string {
	return "web2agent"
}

func (w *WATarget) ExecuteMethod(method, path string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return w.executeMethod(method, path, data)
}

func (w *WATarget) UserContext() *invoker.User {
	return nil
}

func (w *WATarget) GetAttr(string) interface{}       { return nil }
func (w *WATarget) GetAttrs() map[string]interface{} { return nil }
