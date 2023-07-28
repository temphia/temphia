package bundled

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/engine/invokers"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func NewWeb(http *gin.Context, app xtypes.App, eclaim *claim.Executor) invoker.Invoker {

	builder := NewBuilder(invokers.UserApp)

	builder.SetApp(app)
	builder.SetUserContextProvider(func() *claim.UserContext {
		return &claim.UserContext{
			TenantId:  eclaim.TenantId,
			UserID:    eclaim.UserId,
			UserGroup: eclaim.UserGroup,
			SessionID: eclaim.SessionId,
			DeviceId:  eclaim.DeviceId,
		}
	})

	// fixme => set modules bashed on eclaim.ExecType

	return builder.Build()
}

func NewAdmin(dclaim *claim.UserContext) invoker.Invoker {

	builder := NewBuilder(invokers.UserApp)

	builder.SetUserContextProvider(func() *claim.UserContext {
		return &claim.UserContext{
			TenantId:  dclaim.TenantId,
			UserID:    dclaim.UserID,
			UserGroup: dclaim.UserGroup,
			SessionID: dclaim.SessionID,
			DeviceId:  dclaim.DeviceId,
		}
	})

	// fixme => set modules bashed on eclaim.ExecType

	return builder.Build()
}
