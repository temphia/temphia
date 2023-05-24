package hubv2

import "encoding/json"

type PlugSchema struct {
	Steps []PlugSteps `json:"steps,omitempty" yaml:"steps,omitempty"`
}

type PlugSteps struct {
	Name string          `json:"name,omitempty" yaml:"name,omitempty"`
	Type string          `json:"type,omitempty" yaml:"type,omitempty"`
	Data json.RawMessage `json:"data,omitempty" yaml:"data,omitempty"`
}

const (
	PlugStepNewPlug            = "new_plug"
	PlugStepNewAgent           = "new_agent"
	PlugStepUpdateAgent        = "update_agent"
	PlugStepRemoveAgent        = "remove_agent"
	PlugStepAddResourceLink    = "add_resource_link"
	PlugStepRemoveResourceLink = "remove_resource_link"
	PlugStepAddInnerLink       = "add_inner_link"
	PlugStepAddRemoveLink      = "remove_inner_link"

	PlugStepAddResourceModule    = "add_resource"
	PlugStepUpdateResourceModule = "update_resource"
	PlugStepRemoveResourceModule = "remove_resource"

	PlugStepAddTargetApp    = "add_target_app"
	PlugStepUpdateTargetApp = "update_target_app"
	PlugStepDeleteTargetApp = "delete_target_app"

	PlugStepAddTargetHook    = "add_target_hook"
	PlugStepUpdateTargetHook = "update_target_hook"
	PlugStepDeleteTargetHook = "delete_target_hook"
)
