package pagequery

func (pf *PageQuery) load(req *LoadRequest) (*LoadResponse, error) {
	return &LoadResponse{
		Title:  pf.model.Title,
		Stages: pf.model.Stages,
	}, nil
}

func (pf *PageQuery) submit(req *SubmitRequest) (*SubmitResponse, error) {

	return nil, nil
}
