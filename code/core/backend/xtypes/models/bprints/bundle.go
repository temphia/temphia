package bprints

type Bundle map[string]BundleItem

type BundleItem struct {
	Type string `json:"type,omitempty"`
	File string `json:"file,omitempty"`
}
