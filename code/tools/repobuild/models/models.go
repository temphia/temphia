package models

type RepoItem struct {
	GitURL     string `json:"git_url,omitempty" yaml:"git_url,omitempty"`
	Branch     string `json:"branch,omitempty" yaml:"branch,omitempty"`
	BuildCMD   string `json:"build_cmd,omitempty" yaml:"build_cmd,omitempty"`
	BprintFile string `json:"bprint_file,omitempty" yaml:"bprint_file,omitempty"`
}

type BuildConfig struct {
	Items        map[string]RepoItem `json:"items,omitempty" yaml:"items,omitempty"`
	BuildFolder  string              `json:"build_folder,omitempty" yaml:"build_folder,omitempty"`
	OutputFolder string              `json:"output_folder,omitempty" yaml:"output_folder,omitempty"`
	BuildImage   string              `json:"build_image,omitempty" yaml:"build_image,omitempty"`
}
