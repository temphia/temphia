package extension

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

type Handle interface {
	SetExecutorBuilder(name string, builder etypes.ExecutorBuilder)
	SetModuleBuilder(name string, builder etypes.ModuleBuilder)
	SetAdapterBuilder(name string, builder httpx.Builder)
	SetScript(name string, s func(tenantId string, ctx any) error)

	SetOnStart(hook func())
	SetOnExit(hook func())
}

type Accesser interface {
	GetExecutorBuilder() map[string]etypes.ExecutorBuilder
	GetModuleBuilder() map[string]etypes.ModuleBuilder
	GetAdapterBuilder() map[string]httpx.Builder
	GetScript() map[string]func(tenantId string, ctx any) error
}

type Builder func(app xtypes.App, handle Handle) error
