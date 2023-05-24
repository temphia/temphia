package xbprint

type BundleV1 struct {
	Type  string        `json:"type,omitempty"`
	Items []InstallItem `json:"items,omitempty"`
}

type InstallItem struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	File string `json:"file,omitempty"`
}
