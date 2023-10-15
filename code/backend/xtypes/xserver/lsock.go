package xserver

// LSock is service that allows subprocesses to communicate with main app instance
// they are mostly locally run processes example use it in remote executor that is
// spawned in nsjail that want to call bindings or perform action in resources we
// do not enforce domain be localhost but send a token when spawning subprocess
// TEMPHIA_LSOCK_TOKEN which is needed to connect to this service which makes this safe

type LSock interface {
	Register(s LSubcriber) int64
	SendRPX(iid int64, name string, data []byte) ([]byte, error)
}

type LSubcriber interface {
	Handle(name string, data []byte) ([]byte, error)
}

type LSOptions struct {
	TenantId      string `json:"tenant_id,omitempty"`
	PlugId        string `json:"plug_id,omitempty"`
	AgentId       string `json:"agent_id,omitempty"`
	Addr          string `json:"addr,omitempty"`
	RPXPrefix     string `json:"rpx_prefix,omitempty"`
	ControlPrefix string `json:"control_prefix,omitempty"`
	ReplyToken    string `json:"reply_token,omitempty"`
}
