package sockd

import (
	"github.com/temphia/temphia/code/backend/services/sockd/core"
	"github.com/temphia/temphia/code/backend/services/sockd/syncer"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
)

type Sockd struct {
	sockd sockdx.SockdCore

	dataSyncer sockdx.DataSyncer
	userSyncer sockdx.UserSyncer
}

func New(opts sockdx.Options) *Sockd {

	sockd := core.New(opts)

	return &Sockd{
		sockd:      sockd,
		dataSyncer: syncer.NewData(sockd),
		userSyncer: syncer.NewUser(sockd),
	}
}

func (s *Sockd) GetSockd() sockdx.SockdCore { return s.sockd }

func (s *Sockd) GetDataSyncer() sockdx.DataSyncer { return s.dataSyncer }

func (s *Sockd) GetUserSyncer() sockdx.UserSyncer { return s.userSyncer }
