package claim

type Executor struct {
	TenentId   string            `json:"-"`
	UserId     string            `json:"user,omitempty"`
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

// Auth related claims

type OauthState struct {
	TenantId  string `json:"tenant_id,omitempty"`
	AuthId    int64  `json:"id,omitempty"`
	UserGroup string `json:"user_group,omitempty"`
	DeviceId  string `json:"device_id,omitempty"`
}

type AuthFirst struct {
	AuthId    int64  `json:"auth_id,omitempty"`
	NewUser   bool   `json:"new_user,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	UserGroup string `json:"user_group,omitempty"`
	UserEmail string `json:"user_email,omitempty"`
	Type      string `json:"type,omitempty"`
	Expiry    int64  `json:"expiry,omitempty"`
	DeviceId  string `json:"device_id,omitempty"`
}

type AuthNext struct {
	UserId      string `json:"user_id,omitempty"`
	UserGroup   string `json:"user_group,omitempty"`
	UserEmail   string `json:"user_email,omitempty"`
	DeviceId    string `json:"device_id,omitempty"`
	EmailVerify bool   `json:"email_verify,omitempty"`
	PassChange  bool   `json:"pass_change,omitempty"`
}

type PreAuthed struct {
	UserID     string `json:"user_id,omitempty"`
	UserGroup  string `json:"user_group,omitempty"`
	UserEmail  string `json:"user_email,omitempty"`
	AuthId     int64  `json:"auth_id,omitempty"`
	NeedsProof bool   `json:"needs_proof,omitempty"`
	DeviceId   string `json:"device_id,omitempty"`
}
