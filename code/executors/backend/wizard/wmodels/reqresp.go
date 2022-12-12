package wmodels

type RequestSplash struct {
	HasExecData bool `json:"has_exec_data,omitempty"`
}
type ResponseSplash struct {
	WizardTitle string                 `json:"wizard_title,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Fields      []*Field               `json:"fields,omitempty"`
	DataSources map[string]interface{} `json:"data_sources,omitempty"`
	SkipSplash  bool                   `json:"skip_splash,omitempty"`
}

type RequestStart struct {
	SplashData   map[string]interface{} `json:"splash_data,omitempty"`
	StartRawData interface{}            `json:"start_raw_data,omitempty"` // exec_data
}

type ResponseStart struct {
	StartStage  bool                   `json:"stage_started,omitempty"`
	StageTitle  string                 `json:"stage_title,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Fields      []*Field               `json:"fields,omitempty"`
	DataSources map[string]interface{} `json:"data_sources,omitempty"`
	PrevData    map[string]interface{} `json:"prev_data,omitempty"`
	OpaqueData  []byte                 `json:"odata,omitempty"`
	Ok          bool                   `json:"ok"`
}

type RequestStartNested struct {
	ParentOpaqueData []byte      `json:"parent_odata,omitempty"`
	Field            string      `json:"field,omitempty"`
	StartRawData     interface{} `json:"start_raw_data,omitempty"`
}

type RequestNext struct {
	Data       map[string]interface{} `json:"data,omitempty"`
	OpaqueData []byte                 `json:"odata,omitempty"`
}

type ResponseNext struct {
	StageTitle  string                 `json:"stage_title,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Fields      []*Field               `json:"fields,omitempty"`
	DataSources map[string]interface{} `json:"data_sources,omitempty"`
	PrevData    map[string]interface{} `json:"prev_data,omitempty"`
	OpaqueData  []byte                 `json:"odata,omitempty"`
	Ok          bool                   `json:"ok"`
	Final       bool                   `json:"final"`
	Errors      map[string]string      `json:"errors,omitempty"`
}

type ResponseFinal struct {
	LastMessage string      `json:"last_message,omitempty"`
	Ok          bool        `json:"ok"`
	Final       bool        `json:"final"`
	FinalData   interface{} `json:"final_data"`
}

type RequestBack struct {
	OpaqueData []byte `json:"odata,omitempty"`
}

type ResponseBack struct {
	StageTitle  string                 `json:"stage_title,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Fields      []*Field               `json:"fields,omitempty"`
	DataSources map[string]interface{} `json:"data_sources,omitempty"`
	PrevData    map[string]interface{} `json:"prev_data,omitempty"`
	OpaqueData  []byte                 `json:"odata,omitempty"`
	Ok          bool                   `json:"ok"`
	Final       bool                   `json:"final"`
}
