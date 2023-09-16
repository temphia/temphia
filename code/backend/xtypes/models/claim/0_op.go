package claim

type Operator struct {
	XID          string `json:"xid,omitempty"`
	Type         string `json:"type,omitempty"`
	BindDeviceId string `json:"bind_device,omitempty"`
}

type LSock struct {
	Type  string `json:"type,omitempty"`
	SID   int64  `json:"sid,omitempty"`
	Plug  string `json:"plug,omitempty"`
	Agent string `json:"agent,omitempty"`
}
