package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) dev(rg *gin.RouterGroup) {

	rg.GET("/executor", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", []byte(`
		<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
    <title>Play</title>
    <link rel="stylesheet" href="/z/assets/build/executor_pageform.css" />
    <script defer src="/z/assets/build/executor_pageform.js"></script>
  </head>
  	<body></body>  </html>`))
	})

	rg.POST("/executor/preform/:name", func(ctx *gin.Context) {

	})
}
