package wizard

import (
	"encoding/json"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/backend/wizard/lifecycle"
	"github.com/temphia/temphia/code/executors/backend/wizard/wmodels"
)

func (sw *SimpleWizard) GetSplash(ev *event.Request, msg string) (interface{}, error) {
	req := wmodels.RequestSplash{}

	err := json.Unmarshal(ev.Data, &req)
	if err != nil {
		return nil, err
	}

	return sw.getSplash(req.HasExecData, msg)
}

func (sw *SimpleWizard) getSplash(hasExecData bool, msg string) (interface{}, error) {

	dataSources := map[string]interface{}{}
	skipSplash := hasExecData

	if sw.model.Splash.OnLoad != "" {
		lf := lifecycle.OnSplashLoad{
			Models: &sw.model,
			SideEffects: lifecycle.OnSplashLoadSideEffects{
				FailErr:     "",
				SkipSplash:  skipSplash,
				DataSources: map[string]interface{}{},
			},
			HasExecData: skipSplash,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}

		if lf.SideEffects.FailErr != "" {
			return nil, easyerr.Error(lf.SideEffects.FailErr)
		}

		skipSplash = lf.SideEffects.SkipSplash
		dataSources = lf.SideEffects.DataSources
	}

	return wmodels.ResponseSplash{
		WizardTitle: sw.model.Title,
		Message:     msg,
		SkipSplash:  skipSplash,
		Fields:      sw.model.Splash.Fields,
		DataSources: dataSources,
	}, nil
}
