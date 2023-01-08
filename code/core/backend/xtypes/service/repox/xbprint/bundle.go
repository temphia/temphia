package xbprint

type Bundle struct {
	Type  string       `json:"type,omitempty"`
	Items []BundleItem `json:"items,omitempty"`
}

type BundleItem struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	File string `json:"file,omitempty"`
}
