package invoker

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type User struct {
	Id        string
	Group     string
	SessionId int64
	DeviceId  string
}

type Invoker interface {
	Type() string
	ExecuteModule(module, action string, data xtypes.LazyData) (xtypes.LazyData, error)
	ListModules() []string
	UserContext() *User

	GetAttr(string) interface{}
	GetAttrs() map[string]interface{}
}

type DevOptions struct {
	App     xtypes.App
	HttpCtx *gin.Context
	Args    map[string]any
	PlugId  string
	AgentId string
	Claim   *claim.PlugDevTkt
}

type DevInvokerBuilder func(opts DevOptions) Invoker
