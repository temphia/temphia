package tests

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var globalTests []TestGroupBuilder = []TestGroupBuilder{}
var glock sync.Mutex = sync.Mutex{}

func Add(tgb TestGroupBuilder) {
	glock.Lock()
	globalTests = append(globalTests, tgb)
	glock.Unlock()
}

func GlobalGet() []TestGroupBuilder {
	return globalTests
}

type TestHandle struct {
	App     xtypes.App
	CoreHub store.CoreDB
	CabHub  store.CabinetHub
	XPlane  xplane.ControlPlane
}

type TestGroup interface {
	Name() string
	Cases() []string
	Run(string) error
}

type TestGroupBuilder func(handle *TestHandle) TestGroup
