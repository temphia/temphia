package wizard

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/backend/wizard/lifecycle"
	"github.com/temphia/temphia/code/executors/backend/wizard/sloader"
	"github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

	"github.com/thoas/go-funk"
)

func (sw *SimpleWizard) RunNext(ev *event.Request) (interface{}, error) {

	req := wmodels.RequestNext{}

	err := json.Unmarshal(ev.Data, &req)
	if err != nil {
		return nil, err
	}

	subData, err := sw.getSub(req.OpaqueData)
	if err != nil {
		return nil, err
	}

	currstage := sw.model.Stages[subData.CurrentStage]
	if currstage == nil {
		return nil, easyerr.NotFound()
	}

	currgroup := sw.getStageGroup(subData.StageGroup)
	if currgroup == nil {
		return nil, easyerr.NotFound()
	}

	var nstage string
	var skipChecks []string

	if currgroup.BeforeNext != "" {
		lf := lifecycle.BeforeNext{
			CurrentData: req.Data,
			SubData:     subData,
			SideEffects: lifecycle.BeforeNextSideEffects{},
		}
		err := lf.Execute()
		if err != nil {
			return nil, err
		}
		if len(lf.SideEffects.Errors) > 0 {
			return wmodels.ResponseNext{
				Errors: lf.SideEffects.Errors,
				Ok:     false,
			}, nil
		}
		if lf.SideEffects.NextStage != "" {
			nstage = lf.SideEffects.NextStage
		}

		if lf.SideEffects.SkipCheck != nil {
			skipChecks = lf.SideEffects.SkipCheck
		}
	}

	if currstage.BeforeVerify != "" {

		lf := lifecycle.StageBeforeVerify{
			Models: &sw.model,
			SideEffects: lifecycle.StageBeforeVerifyEffect{
				SkipCheck: skipChecks,
			},
			SubData: subData,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}
		if lf.SideEffects.NextStage != "" {
			nstage = lf.SideEffects.NextStage
		}

		if len(lf.SideEffects.Errors) > 0 {
			return wmodels.ResponseNext{
				Errors: lf.SideEffects.Errors,
				Ok:     false,
			}, nil
		}

		if lf.SideEffects.SkipCheck != nil {
			skipChecks = lf.SideEffects.SkipCheck
		}

		if lf.SideEffects.NextStage != "" {
			nstage = lf.SideEffects.NextStage
		}
	}

	pp.Println(nstage)

	errors := make(map[string]string)
	for _, field := range currstage.Fields {
		if funk.ContainsString(skipChecks, field.Name) {
			continue
		}
		ferr := sw.validateField(field, req.Data[field.Name])
		if ferr != "" {
			errors[field.Name] = ferr
		}
	}

	if len(errors) > 0 {
		return wmodels.ResponseNext{
			Errors: errors,
			Ok:     false,
		}, nil
	}

	{
		subData.Data[subData.CurrentStage] = req.Data
		subData.VisitedStages = append(subData.VisitedStages, subData.CurrentStage)
	}

	if nstage == "" {
		idx := sw.stageIndex(currgroup, currstage.Name)

		switch idx {
		case -1:
			pp.Println("@group", currgroup)
			pp.Println("@subdata", subData)
			// this should not happen
			return nil, easyerr.NotFound()
		case len(currgroup.Stages) - 1:
			return sw.endStageGroup(currgroup, subData)
		default:
			nstage = currgroup.Stages[idx+1]
			subData.CurrentStage = nstage
		}

	} else {
		if !funk.ContainsString(currgroup.Stages, nstage) {
			return nil, easyerr.NotFound()
		}
	}

	return sw.generate(subData, currgroup, nstage)
}

func (sw *SimpleWizard) generate(sub *wmodels.Submission, group *wmodels.StageGroup, nStage string) (interface{}, error) {
	datasources := make(map[string]interface{})

	stage := sw.model.Stages[nStage]
	if stage == nil {
		pp.Println("@next_stage", nStage)
		pp.Println("@all_stage", sw.model.Stages)
		return nil, easyerr.NotFound()
	}

	if stage.BeforeGenerate != "" {
		lf := lifecycle.StageBeforeGenerate{
			Models: &sw.model,
			SideEffects: lifecycle.StageBeforeGenerateEffects{
				DataSources: datasources,
			},
			SubData: sub,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}
	}

	loader := sloader.SLoader{
		Binding:     sw.binding,
		Model:       &sw.model,
		SubData:     sub,
		Stage:       stage,
		Group:       group,
		DataSources: datasources,
	}

	err := loader.Process()
	if err != nil {
		return nil, err
	}

	if stage.AfterGenerate != "" {
		lf := lifecycle.StageAfterGenerate{
			Models: &sw.model,
			SideEffects: lifecycle.StageAfterGenerateEffects{
				DataSources: datasources,
			},
			SubData: sub,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}
	}

	opdata, err := sw.updateSub(sub)
	if err != nil {
		return nil, err
	}

	return wmodels.ResponseNext{
		StageTitle:  stage.Name,
		Fields:      stage.Fields,
		DataSources: datasources,
		Message:     stage.Message,
		OpaqueData:  opdata,
		Ok:          true,
		Final:       false,
	}, nil

}

func (sw *SimpleWizard) endStageGroup(group *wmodels.StageGroup, subData *wmodels.Submission) (interface{}, error) {

	if group.BeforeEnd == "" {
		lf := lifecycle.BeforeEnd{
			SideEffects: lifecycle.BeforeEndSideEffects{},
			SubData:     subData,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}

		if lf.SideEffects.GotoStage != "" {
			// fixme => gotomessage
			return sw.generate(subData, group, lf.SideEffects.GotoStage)
		}
	}

	msg := group.LastMessage

	if group.AfterEnd != "" {
		lf := lifecycle.AfterEnd{
			SubData:     subData,
			SideEffects: lifecycle.AfterEndSideEffects{},
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}

		if lf.SideEffects.Message != "" {
			msg = lf.SideEffects.Message
		}
	}

	return wmodels.ResponseFinal{
		Ok:          true,
		LastMessage: msg,
		Final:       true,
	}, nil
}
