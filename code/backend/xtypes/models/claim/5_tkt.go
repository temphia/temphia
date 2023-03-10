package claim

type SockdTkt struct {
	UserId    string `json:"user_id,omitempty"`
	Type      string `json:"type,omitempty"`
	DeviceId  int64  `json:"device_id,omitempty"`
	SessionId int64  `json:"session_id,omitempty"`
	Room      string `json:"room,omitempty"`
}

type RoomTagTkt struct {
	Type   string   `json:"type"`
	Tags   []string `json:"tags"`
	Room   string   `json:"room"`
	Expiry int64    `json:"expiry"`
}

type PlugDevTkt struct {
	TenantId  string   `json:"-"`
	UserId    string   `json:"user_id"`
	UserGroup string   `json:"user_group"`
	BprintId  string   `json:"bprint_id"`
	PlugIds   []string `json:"plug_ids"`
	AllPlugs  bool     `json:"all_plugs"`
}

type UserMgmtTkt struct {
	TenantId    string   `json:"-"`
	Type        string   `json:"type"`
	Group       string   `json:"group"`
	Update      bool     `json:"update"`
	List        bool     `json:"list"`
	ListScope   []string `json:"list_scope"`
	UpdateScope []string `json:"update_scope"`
}

type PlugState struct {
	TenantId  string `json:"-"`
	Type      string `json:"type"`
	UserId    string `json:"user_id"`
	DeviceId  int64  `json:"device_id,omitempty"`
	SessionId int64  `json:"session_id,omitempty"`
	ExecId    int64  `json:"exec_id,omitempty"`
	PlugId    string `json:"plug_id,omitempty"`
	AgentId   string `json:"agent_id,omitempty"`
	KeyPrefix string `json:"key_prefix,omitempty"`
}

type AdviseryTkt struct {
	Type   string `json:"type"`
	XID    string `json:"xid"`
	Expiry int64  `json:"expiry"`
	Data   []byte `json:"data"`
}
