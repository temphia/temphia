package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type Request struct {
	Http    *gin.Context
	Session *claim.Session
}
