package pageform

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

func (pf *Pageform) actionLoad(req LoadRequest) (*Response, error) {

	pp.Println("@model", pf.model)

	ctx := pf.pfCtx(map[string]any{
		"__pf_exec_data__": req,
	})
	if len(pf.model.ExecHint) != 0 {
		ctx.nextStage = pf.model.ExecHint[0]
	}

	ctx.bind()
	if pf.model.OnLoad != "" {
		err := ctx.execute(pf.model.OnLoad, "on_load", "")
		if err != nil {
			return nil, err
		}
	}

	if ctx.nextStage == "" {
		return nil, easyerr.NotFound()
	}

	return pf.generate(ctx)
}

func (pf *Pageform) actionSubmit(req SubmitRequest) (*Response, error) {

	currStage, ok := pf.model.Stages[req.Stage]
	if !ok {
		return nil, easyerr.NotFound()
	}

	ctx := pf.pfCtx(req.Data)
	for _, shook := range currStage.OnSubmit {
		err := ctx.execute(shook.Target, "on_submit", req.Stage)
		if err != nil {
			return nil, err
		}
	}

	if ctx.nextStage == "" {

		max := len(pf.model.ExecHint)

		for idx, shint := range pf.model.ExecHint {

			if shint == req.Stage {
				if max == (idx + 1) {
					return &Response{
						Ok:      true,
						Final:   true,
						Message: "Reached the final stage",
						Items:   []FormItem{},
					}, nil
				}

				ctx.nextStage = pf.model.ExecHint[idx+1]
				break
			}

			if max == (idx + 1) {
				return &Response{
					Ok:      false,
					Final:   true,
					Message: "Reached the final stage",
					Items:   []FormItem{},
				}, nil
			}

		}
	}

	return pf.generate(ctx)
}

func (pf *Pageform) generate(ctx *PfCtx) (*Response, error) {
	stage, ok := pf.model.Stages[ctx.nextStage]
	if !ok {
		return &Response{
			Ok:      false,
			Message: "stage not found",
		}, nil
	}

	ctx.applyData(stage.Data)
	for _, shook := range stage.OnBuild {
		err := ctx.execute(shook.Target, "on_submit", ctx.nextStage)
		if err != nil {
			return nil, err
		}
	}

	pp.Println("@ctx.message", ctx.message)

	return &Response{
		Ok:      ctx.ok,
		Final:   ctx.final,
		Message: ctx.message,
		Items:   stage.GetItems(ctx.disabledFields),
		Data:    ctx.data,
		Stage:   ctx.nextStage,
	}, nil
}

func (pf *Pageform) pfCtx(data map[string]any) *PfCtx {

	return &PfCtx{
		data:           data,
		model:          pf.model,
		disabledFields: make([]string, 0),
		rt:             pf.jsruntime,
		ok:             true,
	}

}
