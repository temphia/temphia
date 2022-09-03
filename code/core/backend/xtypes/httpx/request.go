package httpx

import "github.com/gin-gonic/gin"

type Request struct {
	Http    *gin.Context
	Session any
}
