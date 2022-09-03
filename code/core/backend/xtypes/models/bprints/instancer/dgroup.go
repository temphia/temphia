package instancer

type DataGroupRequest struct {
	DyndbSource   string                      `json:"dyndb_source,omitempty"`
	GroupName     string                      `json:"group_name,omitempty"`
	GroupSlug     string                      `json:"group_slug,omitempty"`
	CabinetSource string                      `json:"cabinet_source,omitempty"`
	CabinetFolder string                      `json:"cabinet_folder,omitempty"`
	SeedType      string                      `json:"seed_source,omitempty"`
	TableOptions  map[string]*DataTableOption `json:"table_options,omitempty"`
	UserId        string                      `json:"-"`
}

type DataGroupResponse struct {
	Source     string            `json:"source,omitempty"`
	GroupSlug  string            `json:"gslug,omitempty"`
	GroupName  string            `json:"name,omitempty"`
	SeedError  string            `json:"seed_error,omitempty"`
	ViewErrors map[string]string `json:"view_errors,omitempty"`
}

type DataTableOption struct {
	Name         string `json:"name,omitempty"`
	Slug         string `json:"slug,omitempty"`
	ActivityType string `json:"activity_type,omitempty"`
	SyncType     string `json:"sync_type,omitempty"`
	Seed         bool   `json:"seed,omitempty"`
}
