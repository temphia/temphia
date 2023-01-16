package claim

type SockdTkt struct {
	UserId      string   `json:"user_id"`
	Type        string   `json:"type"`
	XID         string   `json:"xid"`
	ConnectTags []string `json:"tags"`
	Room        string   `json:"room"`
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

type AdviseryTkt struct {
	Type   string `json:"type"`
	XID    string `json:"xid"`
	Expiry int64  `json:"expiry"`
	Data   []byte `json:"data"`
}
