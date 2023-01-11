package wizard

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/backend/wizard/wmodels"
	"github.com/thoas/go-funk"
)

func (sw *SimpleWizard) RunNestedStart(ev *event.Request) (interface{}, error) {
	data := wmodels.RequestStartNested{}

	err := json.Unmarshal(ev.Data, &data)
	if err != nil {
		return nil, err
	}

	sub, err := sw.getSub(data.ParentOpaqueData)
	if err != nil {
		return nil, err
	}

	if sub.ParentStageGroup != "" {
		panic("cannot have double nested stages")
	}

	pgroup := sw.getStageGroup(sub.StageGroup)
	if pgroup == nil {
		panic("Empty parent group to start nested stage")
	}

	if !funk.ContainsString(pgroup.Stages, sub.CurrentStage) {
		panic("stage is not in current group")
	}

	pstage := sw.model.Stages[sub.CurrentStage]
	if pstage == nil {
		panic("Empty parent stage to start nested stage")
	}

	field := pstage.GetField(data.Field)
	if field == nil {
		panic("field not found, to start nested stage")
	}

	return sw.runStart(sub.StageGroup, sub.CurrentStage, sub.Id, field.Attrs["nested_stage_group"].(string), data.StartRawData)
}
