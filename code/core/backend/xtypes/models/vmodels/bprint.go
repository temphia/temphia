package vmodels

import "encoding/json"

type RepoImportOpts struct {
	Slug      string   `json:"slug,omitempty"`
	Source    string   `json:"source,omitempty"`
	Group     string   `json:"group,omitempty"`
	SkipFiles []string `json:"skip_files,omitempty"`
	NewId     string   `json:"new_id,omitempty"`
}

type RepoInstallOpts struct {
	BprintId string          `json:"bprint_id,omitempty"`
	Data     json.RawMessage `json:"data,omitempty"`
	UserId   string          `json:"-"`
}

// DEPRICATE
type DGroupInstallOptions struct {
	DyndbSource   string `json:"dyndb_source,omitempty"`
	GroupName     string `json:"group_name,omitempty"`
	GroupSlug     string `json:"group_slug,omitempty"`
	Schema        string `json:"schema,omitempty"`
	CabinetSource string `json:"cabinet_source,omitempty"`
	CabinetFolder string `json:"cabinet_folder,omitempty"`
	SeedFrom      string `json:"seed_from,omitempty"`
	UserId        string `json:"-"`
}

// DEPRICATE
type DTableInstallOptions struct {
	TargetSource  string `json:"target_source,omitempty"`
	TargetGroupId string `json:"target_group_id,omitempty"`
	TableSlug     string `json:"table_slug,omitempty"`
	Schema        string `json:"schema,omitempty"`
	SeedRandom    bool   `json:"seed_random,omitempty"`
}

// DEPRICATE
type PlugInstallOptions struct {
	NewPlugId string   `json:"new_plug_id,omitempty"`
	Agents    []string `json:"agents,omitempty"`
	Resources []string `json:"resources,omitempty"`
	Schema    string   `json:"schema,omitempty"`
}

type PlugInstallResponse struct {
	Agents       map[string]string `json:"agents,omitempty"`
	Resources    map[string]string `json:"resources,omitempty"`
	ErrAgents    map[string]string `json:"err_agents,omitempty"`
	ErrResources map[string]string `json:"err_resources,omitempty"`
}

type DynData struct {
	Data map[string][]map[string]interface{} `json:"data,omitempty"`
}
