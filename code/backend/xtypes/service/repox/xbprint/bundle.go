package xbprint

type Install struct {
	Type  string        `json:"type,omitempty"`
	Items []InstallItem `json:"items,omitempty"`
}

type InstallItem struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	File string `json:"file,omitempty"`
}

type Upgrade struct {
	MinVersion string `json:"min_version,omitempty"`
}
