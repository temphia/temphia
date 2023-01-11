package httpx

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NotFound(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusNotFound)
	ctx.Writer.Write([]byte(`<h1>Not Found<h1>`))
}

func WriteFile(file string, data []byte, ctx *gin.Context) {

	ffiles := strings.Split(file, ".")

	ctype := ""
	switch ffiles[1] {
	case "js":
		ctype = "application/javascript"
	case "css":
		ctype = "text/css"
	default:
		ctype = http.DetectContentType(data)
	}

	ctx.Data(http.StatusOK, ctype, data)
}
