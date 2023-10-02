package xplane

type ControlPlane interface {
	Start() error
	//	Inject(iapp any, config *config.AppConfig) error

	NotifyStat(stats NodeStat) error

	GetAppStatus() (*AppStatus, error)
	GetLocker() Locker

	GetSockdRouter() SockdRouter

	GetMsgBus() MsgBus

	GetNodeId() int64
	GetIdService() IDService
}
