package instance

type DataTableRequest struct {
	Name         string `json:"name,omitempty"`
	Slug         string `json:"slug,omitempty"`
	ActivityType string `json:"activity_type,omitempty"`
	SyncType     string `json:"sync_type,omitempty"`
	SeedType     string `json:"seed_source,omitempty"`
	Seed         bool   `json:"seed,omitempty"`
}

type DataTableResponse struct {
	Source     string            `json:"source,omitempty"`
	GroupSlug  string            `json:"gslug,omitempty"`
	GroupName  string            `json:"gname,omitempty"`
	TableSlug  string            `json:"tslug,omitempty"`
	TableName  string            `json:"tname,omitempty"`
	SeedError  string            `json:"seed_error,omitempty"`
	ViewErrors map[string]string `json:"view_errors,omitempty"`
}
