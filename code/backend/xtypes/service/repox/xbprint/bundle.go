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

const (
	StepNewData        = "new_data"
	StepUpdateData     = "update_data"
	StepNewPlug        = "new_plug"
	StepUpdatePlug     = "update_plug"
	StepNewResource    = "new_resource"
	StepUpdateResource = "update_resource"
)

type BundleV2 struct {
	Type  string       `json:"type,omitempty"`
	Items []BundleStep `json:"items,omitempty"`
}

type BundleStep struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	File string `json:"file,omitempty"`
}
