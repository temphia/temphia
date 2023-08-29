package xserver

import (
	"github.com/temphia/temphia/code/backend/xtypes"
)

// DEPRICATE use remote ???

type Mesh interface {
	Start(app xtypes.App) error
	Stop() error
	GetAddress() []string
}

type XRouter interface {
}
