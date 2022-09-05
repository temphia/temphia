package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusNotFound)
	ctx.Writer.Write([]byte(`<h1>Not Found<h1>`))
}
