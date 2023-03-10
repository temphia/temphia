package ticket

type PlugState struct {
	KeyPrefix string `json:"key_prefix,omitempty"`
}

type CabinetFolder struct {
	Prefix      string   `json:"prefix,omitempty"`
	PinnedFiles []string `json:"pinned_files,omitempty"`
	Operations  []string `json:"ops,omitempty"`
}

type SockdRoom struct {
	AllowBroadcast string `json:"allow_broadcast,omitempty"`
}
