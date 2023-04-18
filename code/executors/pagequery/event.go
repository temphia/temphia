package pagequery

func (pf *PageQuery) load(req *LoadRequest) (*LoadResponse, error) {
	return &LoadResponse{
		Title:      pf.model.Title,
		Stages:     pf.model.Stages,
		FirstStage: pf.model.FirstStage,
	}, nil
}

func (pf *PageQuery) submit(req *SubmitRequest) (*SubmitResponse, error) {

	script := req.Script

	if script == "" {
		stage := pf.model.Stages[req.Stage]
		script = stage.Script
	}

	pfctx := pf.new(req.ExecData, req.ParamData, req.Stage)

	resp, err := pfctx.execute(script)
	if err != nil {
		return nil, err
	}

	return &SubmitResponse{
		Stage:    req.Stage,
		Data:     resp.Data,
		Elements: resp.Elements,
	}, nil
}
