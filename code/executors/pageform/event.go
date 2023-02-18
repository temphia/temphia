package pageform

import "github.com/temphia/temphia/code/backend/libx/easyerr"

func (pf *Pageform) actionLoad(req LoadRequest) (*Response, error) {

	ctx := pf.pfCtx(map[string]any{
		"__pf_exec_data__": req,
	}, pf.model.ExecHint[0])

	ctx.bind()
	if pf.model.OnLoad != "" {

		err := ctx.execute(pf.model.OnLoad, "on_load")
		if err != nil {
			return nil, err
		}
	}

	if ctx.nextStage == "" {
		return nil, easyerr.NotFound()
	}

	stage, ok := pf.model.Stages[ctx.nextStage]
	if !ok {
		return &Response{
			Ok:      false,
			Message: "stage not found",
		}, nil
	}

	if stage.OnGenerate != "" {
		ctx.applyData(stage.Data)
		ctx.execute(stage.OnGenerate, "on_generate")
	}

	return &Response{
		Ok:      ctx.message == "",
		Message: ctx.message,
		Items:   stage.GetItems(ctx.disabledFields),
		Data:    stage.Data,
		Stage:   ctx.nextStage,
	}, nil

}

func (pf *Pageform) actionSubmit(req SubmitRequest) (*Response, error) {

	// pp.Println(req)

	// message := ""
	// ok := true
	// items := pf.model.Items["second"]
	// if items == nil {
	// 	ok = false
	// 	message = "stage not found"
	// }

	// return &Response{
	// 	Ok:      ok,
	// 	Message: message,
	// 	Items:   items,
	// }, nil

	return nil, nil
}

func (pf *Pageform) pfCtx(data map[string]any, stage string) *PfCtx {

	return &PfCtx{
		data:           data,
		model:          pf.model,
		disabledFields: make([]string, 0),
		message:        "",
		nextStage:      stage,
	}

}
