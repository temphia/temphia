package xbprint

type BundleV1 struct {
	Type         string       `json:"type,omitempty"`
	Items        []BundleItem `json:"items,omitempty"`
	PostInstance string       `json:"post_instance,omitempty"`
	PostUpgrade  string       `json:"post_upgrade,omitempty"`
}

type BundleItem struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	File string `json:"file,omitempty"`
}
