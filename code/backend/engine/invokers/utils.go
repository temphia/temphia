package invokers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service"
)

var (
	ErrInvokerActionNotImplemented = errors.New("INVOKER ACTION NOT IMPLEMENTED")
	ErrInvokerActionNotAllowed     = errors.New("INVOKER ACTION NOT ALLOWED")
)

const (
	Forked          = "forked"
	Admin           = "admin"
	UserApp         = entities.TargetAppTypeUserGroupApp
	DataTableWidget = entities.TargetAppTypeDataTableWidget
	DataSheetWidget = entities.TargetAppTypeDataTableWidget
)

func ParseClaim(signer service.Signer, ctx *gin.Context) (*claim.Session, error) {
	return signer.ParseSession(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
}
