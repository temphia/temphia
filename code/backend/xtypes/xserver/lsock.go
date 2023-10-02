package xserver

import "encoding/json"

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

// remote execution info

type REInfo struct {
	Addr          string `json:"addr,omitempty"`
	RPXPrefix     string `json:"rpx_prefix,omitempty"`
	ControlPrefix string `json:"control_prefix,omitempty"`
	ReplyToken    string `json:"reply_token,omitempty"`
}

type REPacketIn struct {
	Name    string          `json:"name,omitempty"`
	UserCtx any             `json:"user_ctx,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

type REPacketOut struct {
	Name    string `json:"name,omitempty"`
	UserCtx any    `json:"user_ctx,omitempty"`
	Data    any    `json:"data,omitempty"`
}
