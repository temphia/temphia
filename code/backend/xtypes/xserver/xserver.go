package xserver

import (
	"github.com/temphia/temphia/code/backend/xtypes/xserver/remote"
)

type Server interface {
	Listen() error
	Close() error
	BuildRoutes() error
	GetAdapterHub() any
	HandleRemote(req *remote.Request) (*remote.Response, error)
}
