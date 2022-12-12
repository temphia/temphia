package wizard

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

	"github.com/thoas/go-funk"
)

func (sw *SimpleWizard) RunBack(ev *event.Request) (interface{}, error) {

	req := wmodels.RequestBack{}

	err := json.Unmarshal(ev.Data, &req)
	if err != nil {
		return nil, err
	}

	sub, err := sw.getSub(req.OpaqueData)
	if err != nil {
		return nil, err
	}

	if len(sub.VisitedStages) == 0 {
		panic("cannot back further")
	}

	nextstage := ""
	currIndex := funk.IndexOfString(sub.VisitedStages, sub.CurrentStage)
	switch currIndex {
	case -1:
		nextstage = sub.VisitedStages[len(sub.VisitedStages)-1]
	case 0:
		nextstage = sub.CurrentStage
	default:
		nextstage = sub.VisitedStages[currIndex-1]
	}

	pp.Println(sub, nextstage)

	// json.Unmarshal()

	return nil, nil
}
