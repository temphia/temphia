package claim

type Executor struct {
	TenentId   string            `json:"-"`
	UserId     string            `json:"user_id,omitempty"`
	UserGroup  string            `json:"group,omitempty"`
	DeviceId   string            `json:"device_id,omitempty"`
	Type       string            `json:"type,omitempty"`
	SessionId  int64             `json:"session_id,omitempty"`
	ExecId     int64             `json:"exec_id,omitempty"`
	PlugId     string            `json:"plug_id,omitempty"`
	AgentId    string            `json:"agent_id,omitempty"`
	ExecType   string            `json:"exec_type,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type Data struct {
	TenentId   string   `json:"-"`
	Type       string   `json:"type,omitempty"`
	UserID     string   `json:"user_id,omitempty"`
	UserGroup  string   `json:"ugroup,omitempty"`
	SessionID  int64    `json:"session_id,omitempty"`
	DeviceId   string   `json:"device_id,omitempty"`
	DataSource string   `json:"source,omitempty"`
	DataGroup  string   `json:"group,omitempty"`
	DataTables []string `json:"tables,omitempty"`
	IsExec     bool     `json:"is_exec,omitempty"`
}

type Folder struct {
	TenentId  string `json:"-"`
	UserId    string `json:"user_id,omitempty"`
	SessionID int64  `json:"session_id,omitempty"`
	DeviceId  string `json:"device_id,omitempty"`
	Type      string `json:"type,omitempty"`
	Expiry    int64  `json:"expiry,omitempty"`
	Source    string `json:"source,omitempty"`
	Folder    string `json:"folder,omitempty"`
}
