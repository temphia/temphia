package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type Request struct {
	Id      int64
	Http    *gin.Context
	Session *claim.Session
}
