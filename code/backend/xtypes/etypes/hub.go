package etypes

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type EngineHub interface {
	GetEngine() Engine

	Start() error

	// RunDataHook()
	// RunWebApp()
	// RunStartupHooks()
}

type PoolManager interface {
	CloseSidePool(id int)
	GetSidePool() ExecPool

	RootPool() ExecPool
}

type ExecPool interface {
	Borrow(plugId, agentId string) (bindx.Bindings, int)
	Destroy(plugId string, agentIds []string)
	Return(b bindx.Bindings)
	SetEpoch(plug, agent string, e int64)
	PoolId() int
}
