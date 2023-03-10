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
	builder.SetUserContextProvider(func() *invoker.User {
		return &invoker.User{
			Id:        eclaim.UserId,
			Group:     eclaim.UserGroup,
			SessionId: eclaim.SessionId,
			DeviceId:  eclaim.DeviceId,
		}
	})

	// fixme => set modules bashed on eclaim.ExecType

	return builder.Build()
}

func NewAdmin(dclaim *claim.PlugDevTkt) invoker.Invoker {

	builder := NewBuilder(invokers.UserApp)

	builder.SetUserContextProvider(func() *invoker.User {
		return &invoker.User{
			Id:        dclaim.UserId,
			Group:     dclaim.UserGroup,
			SessionId: 0,
			DeviceId:  0,
		}
	})

	// fixme => set modules bashed on eclaim.ExecType

	return builder.Build()
}
