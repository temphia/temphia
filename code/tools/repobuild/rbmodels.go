package repobuild

type RepoItem struct {
	GitURL       string `json:"git_url,omitempty" yaml:"git_url,omitempty"`
	Branch       string `json:"branch,omitempty" yaml:"branch,omitempty"`
	BuildCommand string `json:"build_command,omitempty" yaml:"build_command,omitempty"`
	OutputFolder string `json:"output_folder,omitempty" yaml:"output_folder,omitempty"`
}

type BuildConfig struct {
	Items        map[string]RepoItem `json:"items,omitempty" yaml:"items,omitempty"`
	BuildFolder  string              `json:"build_folder,omitempty" yaml:"build_folder,omitempty"`
	OutputFolder string              `json:"output_folder,omitempty" yaml:"output_folder,omitempty"`
}

type BuildResult struct {
	ErroredItems map[string]error
	Outputs      map[string]string
}
