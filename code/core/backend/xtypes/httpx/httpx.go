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

	switch ffiles[1] {
	case "js":
		ctx.Writer.Header().Set("Content-Type", "application/javascript")
	case "css":
		ctx.Writer.Header().Set("Content-Type", "text/css")
	default:
		ctx.Writer.Header().Set("Content-Type", http.DetectContentType(data))
	}

	ctx.Writer.Write(data)
}
