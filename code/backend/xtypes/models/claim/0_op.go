package claim

type Operator struct {
	XID          string `json:"xid,omitempty"`
	Type         string `json:"type,omitempty"`
	BindDeviceId string `json:"bind_device,omitempty"`
}

type LSock struct {
	Type  string `json:"type,omitempty"`
	IID   int64  `json:"iid,omitempty"` // instance id
	EID   int64  `json:"eid,omitempty"` // execution id
	Plug  string `json:"plug,omitempty"`
	Agent string `json:"agent,omitempty"`
}
