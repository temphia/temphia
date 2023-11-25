package models

type RepoItem struct {
	GitURL     string `json:"git_url,omitempty" toml:"git_url,omitempty"`
	Branch     string `json:"branch,omitempty" toml:"branch,omitempty"`
	BuildCMD   string `json:"build_cmd,omitempty" toml:"build_cmd,omitempty"`
	BprintFile string `json:"bprint_file,omitempty" toml:"bprint_file,omitempty"`
}

type BuildConfig struct {
	Items        map[string]RepoItem `json:"items,omitempty" toml:"items,omitempty"`
	BuildFolder  string              `json:"build_folder,omitempty" toml:"build_folder,omitempty"`
	OutputFolder string              `json:"output_folder,omitempty" toml:"output_folder,omitempty"`
	BuildImage   string              `json:"build_image,omitempty" toml:"build_image,omitempty"`
}
