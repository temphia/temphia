package middleware

import (
	"mime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
)

type Middleware struct {
	Signer service.Signer
	Logger zerolog.Logger
}

func (m *Middleware) AsFile(data []byte, ext string) func(ctx *gin.Context) {
	exmime := mime.TypeByExtension(ext)
	clen := strconv.FormatInt(int64(len(data)), 10)

	// fixme => also set etag ?

	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", exmime)
		ctx.Header("Cache-Control", `public, max-age=86400`)
		ctx.Header("Content-Length", clen)
		ctx.Writer.Write(data)
	}
}
