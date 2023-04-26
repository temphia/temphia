package pagequery

type PgModel struct {
	Title      string                `json:"name,omitempty" yaml:"name,omitempty"`
	Stages     map[string]QueryStage `json:"stages,omitempty" yaml:"stages,omitempty"`
	FirstStage string                `json:"first_stage,omitempty" yaml:"first_stage,omitempty"`
}

type QueryStage struct {
	Script    string            `json:"script,omitempty" yaml:"script,omitempty"`
	About     string            `json:"about,omitempty" yaml:"about,omitempty"`
	ParamForm map[string]string `json:"param_form,omitempty" yaml:"param_form,omitempty"`
}

type Element struct {
	Name     string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type     string         `json:"type,omitempty" yaml:"type,omitempty"`
	Info     string         `json:"info,omitempty" yaml:"info,omitempty"`
	ViewOpts map[string]any `json:"view_opts,omitempty" yaml:"view_opts,omitempty"`
	DataOpts map[string]any `json:"data_opts,omitempty" yaml:"data_opts,omitempty"`
	Source   string         `json:"source,omitempty" yaml:"source,omitempty"`
}

type LoadRequest struct {
	ExecData map[string]any `json:"exec_data,omitempty"`
}

type LoadResponse struct {
	Title      string                `json:"name,omitempty"`
	Stages     map[string]QueryStage `json:"stages,omitempty"`
	FirstStage string                `json:"first_stage,omitempty"`
}

type SubmitRequest struct {
	Stage     string            `json:"stage,omitempty"`
	Script    string            `json:"script,omitempty"`
	ExecData  map[string]any    `json:"exec_data,omitempty"`
	ParamData map[string]string `json:"param_data,omitempty"`
}

type SubmitResponse struct {
	Stage    string         `json:"stage,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
	Elements []Element      `json:"elements,omitempty"`
}
