package wizard

import (
	"encoding/json"

	"github.com/temphia/temphia/code/executors/backend/wizard/wmodels"
	"github.com/thoas/go-funk"
)

func (sw *SimpleWizard) getStageGroup(group string) *wmodels.StageGroup {
	for _, grp := range sw.model.StageGroups {
		if group != grp.Name {
			continue
		}
		return &grp
	}

	return nil
}

func (sw *SimpleWizard) stageIndex(group *wmodels.StageGroup, stage string) int {
	return funk.IndexOfString(group.Stages, stage)
}

func (sw *SimpleWizard) getSub(opData []byte) (*wmodels.Submission, error) {
	subData := wmodels.Submission{}
	err := json.Unmarshal(opData, &subData)
	if err != nil {
		return nil, err
	}

	if subData.Data == nil {
		subData.Data = make(map[string]map[string]interface{})
	}

	if subData.SharedVars == nil {
		subData.SharedVars = make(map[string]interface{})
	}

	return &subData, nil
}

func (sw *SimpleWizard) updateSub(sdata *wmodels.Submission) ([]byte, error) {
	return json.Marshal(sdata)
}
