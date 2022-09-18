package sockdhub

import (
	"github.com/temphia/temphia/code/core/backend/services/sockdhub/core"
	"github.com/temphia/temphia/code/core/backend/services/sockdhub/syncer"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
)

type SockdHub struct {
	sockd sockdx.SockdCore

	dataSyncer sockdx.DataSyncer
	userSyncer sockdx.UserSyncer
}

func New(opts sockdx.Options) *SockdHub {

	sockd := core.New(opts)

	return &SockdHub{
		sockd:      sockd,
		dataSyncer: syncer.NewData(sockd),
		userSyncer: syncer.NewUser(sockd),
	}
}

func (s *SockdHub) GetSockd() sockdx.SockdCore { return s.sockd }

func (s *SockdHub) GetDataSyncer() sockdx.DataSyncer { return s.dataSyncer }

func (s *SockdHub) GetUserSyncer() sockdx.UserSyncer { return s.userSyncer }
