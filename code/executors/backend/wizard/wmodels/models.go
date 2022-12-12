package wmodels

import "github.com/rs/xid"

type Wizard struct {
	Title       string             `json:"title,omitempty"`
	Splash      Splash             `json:"splash,omitempty"`
	Dev         bool               `json:"dev,omitempty"`
	StageGroups []StageGroup       `json:"stage_groups,omitempty"`
	Stages      map[string]*Stage  `json:"stages,omitempty"`
	Sources     map[string]*Source `json:"sources,omitempty"`
}

type StageGroup struct {
	Name        string   `json:"name,omitempty"`
	Stages      []string `json:"stages,omitempty"`
	LastMessage string   `json:"last_message,omitempty"`
	PreventBack bool     `json:"prevent_back,omitempty"`

	// server side hooks

	BeforeStart string `json:"before_start,omitempty"`
	AfterStart  string `json:"after_start,omitempty"`
	BeforeEnd   string `json:"before_end,omitempty"`
	AfterEnd    string `json:"after_end,omitempty"`
	BeforeNext  string `json:"before_next,omitempty"`
	BeforeBack  string `json:"before_back,omitempty"`

	// client side hooks
	ClientSide         bool   `json:"client_side,omitempty"`
	ClientBeforeSubmit string `json:"ui_before_submit,omitempty"`
	ClientAfterLoad    string `json:"ui_after_load,omitempty"`
	ClientOnDataChange string `json:"ui_on_data_change,omitempty"`
}

type Splash struct {
	Message  string   `json:"message,omitempty"`
	Fields   []*Field `json:"fields,omitempty"`
	Skip     bool     `json:"skip,omitempty"`
	OnLoad   string   `json:"on_load,omitempty"`
	OnSubmit string   `json:"on_submit,omitempty"`
}

// fixme => allow only "determisnistic" kinds of bindings

type Stage struct {
	Name           string   `json:"name,omitempty"`
	Message        string   `json:"message,omitempty"`
	Fields         []*Field `json:"fields,omitempty"`
	BeforeGenerate string   `json:"before_generate,omitempty"`
	AfterGenerate  string   `json:"after_generate,omitempty"`
	BeforeVerify   string   `json:"before_verify,omitempty"`
	AfterVerify    string   `json:"after_verify,omitempty"`
}

func (s *Stage) GetField(name string) *Field {
	for _, f := range s.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}

type Field struct {
	Name     string                 `json:"name,omitempty"`
	Info     string                 `json:"info,omitempty"`
	Type     string                 `json:"type,omitempty"`
	Pattern  string                 `json:"pattern,omitempty"`
	Optional bool                   `json:"optional,omitempty"`
	Attrs    map[string]interface{} `json:"attrs,omitempty"`
	Source   string                 `json:"source,omitempty"`
}

type Source struct {
	Name      string                 `json:"name,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Target    string                 `json:"target,omitempty"`
	Data      interface{}            `json:"data,omitempty"`
	ExtraMeta map[string]interface{} `json:"extra_meta,omitempty"`

	// auto_cache ? => by_submission by_plug by_user by_stage_group by_stage by_field by_ctx_var_key
}

type DataSource struct {
	Data     interface{} `json:"data,omitempty"`
	LazyLoad bool        `json:"lazyload,omitempty"`
}

// this will be send as opaque data
type Submission struct {
	Id               string                            `json:"id,omitempty"`
	StageGroup       string                            `json:"stage_group,omitempty"`
	CurrentStage     string                            `json:"curr_stage,omitempty"`
	Data             map[string]map[string]interface{} `json:"data,omitempty"` // all stage data
	SharedVars       map[string]interface{}            `json:"shared_vars,omitempty"`
	ParentSubId      string                            `json:"parent_subid,omitempty"`
	ParentStageGroup string                            `json:"parent_stage_group,omitempty"` // incase of nested group
	ParentStage      string                            `json:"parent_stage,omitempty"`
	VisitedStages    []string                          `json:"visited_stages,omitempty"`
}

func NewSub(pgroup, pstage, psubid, group, stage string) Submission {
	return Submission{
		Id:               xid.New().String(),
		StageGroup:       group,
		CurrentStage:     stage,
		Data:             make(map[string]map[string]interface{}),
		SharedVars:       make(map[string]interface{}),
		ParentStageGroup: pgroup,
		ParentStage:      pstage,
		ParentSubId:      psubid,
		VisitedStages:    []string{},
	}

}
