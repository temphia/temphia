package hubv2

type PlugSchema struct {
	Steps []Step `json:"steps,omitempty" yaml:"steps,omitempty"`
}

const (
	PlugStepNewPlug       = "new_plug"
	PlugStepNewAgent      = "new_agent"
	PlugStepUpdateAgent   = "update_agent"
	PlugStepRemoveAgent   = "remove_agent"
	PlugStepAddInnerLink  = "add_inner_link"
	PlugStepAddRemoveLink = "remove_inner_link"
)
