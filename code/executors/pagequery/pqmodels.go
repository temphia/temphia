package pagequery

type PgModel struct {
	Title      string                `json:"name,omitempty" yaml:"name,omitempty"`
	Stages     map[string]QueryStage `json:"stages,omitempty" yaml:"stages,omitempty"`
	FirstStage string                `json:"first_stage,omitempty" yaml:"first_stage,omitempty"`
}

type QueryStage struct {
	Script         string             `json:"script,omitempty" yaml:"script,omitempty"`
	About          string             `json:"about,omitempty" yaml:"about,omitempty"`
	NeedsParamData bool               `json:"needs_param_data,omitempty" yaml:"needs_param_data,omitempty"`
	Parameters     map[string]Element `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	ParamGetAction string             `json:"param_get_actions,omitempty" yaml:"param_get_actions,omitempty"`
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
	ExecData map[string]any `json:"exec_data,omitempty" yaml:"exec_data,omitempty"`
}

type LoadResponse struct {
	Title      string                `json:"name,omitempty" yaml:"name,omitempty"`
	Stages     map[string]QueryStage `json:"stages,omitempty" yaml:"stages,omitempty"`
	FirstStage string                `json:"first_stage,omitempty" yaml:"first_stage,omitempty"`
}

type SubmitRequest struct {
	Stage     string         `json:"stage,omitempty" yaml:"stage,omitempty"`
	ParamData map[string]any `json:"param_data,omitempty" yaml:"param_data,omitempty"`
	Script    string         `json:"script,omitempty" yaml:"script,omitempty"`
}

type SubmitResponse struct {
	Stage    string             `json:"stage,omitempty" yaml:"stage,omitempty"`
	Data     map[string]any     `json:"data,omitempty" yaml:"data,omitempty"`
	Elements map[string]Element `json:"elements,omitempty" yaml:"elements,omitempty"`
}
