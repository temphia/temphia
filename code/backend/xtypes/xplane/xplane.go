package xplane

import "github.com/temphia/temphia/code/backend/xtypes/etypes/job"

type ControlPlane interface {
	Start() error
	//	Inject(iapp any, config *config.AppConfig) error
	SetJobChan(chan *job.Job)
	NotifyStat(stats NodeStat) error

	GetAppStatus() (*AppStatus, error)
	GetLocker() Locker
	GetRouter() Router

	GetSockdRouter() SockdRouter

	GetMsgBus() MsgBus

	GetNodeId() int64
	GetIdService() IDService
}
