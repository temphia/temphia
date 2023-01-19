package web2agent

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WATarget struct {
	adapter *Web2Agent
	rid     int64
	http    *gin.Context
}

func (w *WATarget) handle() {
	w.http.Data(http.StatusOK, "text/html", []byte(`not implemented`))
}
