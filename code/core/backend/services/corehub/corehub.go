package corehub

import (
	sockdhub "github.com/temphia/temphia/code/core/backend/services/sockd/hub"

	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

var _ store.CoreHub = (*CoreHub)(nil)

type CoreHub struct {
	coredb   store.CoreDB
	sockdhub sockdhub.SockdHub
	cplane   xplane.ControlPlane
}

func New(coredb store.CoreDB, sockd sockdx.SockdCore, cplane xplane.ControlPlane) *CoreHub {
	return &CoreHub{
		coredb:   coredb,
		sockdhub: *sockdhub.New(sockd),
		cplane:   cplane,
	}
}
