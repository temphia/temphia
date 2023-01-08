package repox

type RepoImportOpts struct {
	Slug      string   `json:"slug,omitempty"`
	Group     string   `json:"group,omitempty"`
	Source    int64    `json:"source,omitempty"`
	SkipFiles []string `json:"skip_files,omitempty"`
	NewId     string   `json:"new_id,omitempty"`
}
