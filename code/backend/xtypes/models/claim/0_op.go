package claim

type Operator struct {
	XID          string `json:"xid,omitempty"`
	Type         string `json:"type,omitempty"`
	BindDeviceId string `json:"bind_device,omitempty"`
}
