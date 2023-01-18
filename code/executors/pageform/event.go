package pageform

func (pf *Pageform) actionLoad(req LoadRequest) (*LoadResponse, error) {

	return &LoadResponse{
		Ok:       true,
		Message:  "",
		Items:    pf.model.Items,
		Data:     map[string]any{},
		OnLoad:   pf.model.ClientOnLoad,
		OnSubmit: pf.model.ClientOnSubmit,
	}, nil

}

func (pf *Pageform) actionSubmit(req SubmitRequest) (*SubmitResponse, error) {

	return &SubmitResponse{
		Ok:      true,
		Message: "",
		Items:   []ResultItem{},
	}, nil
}
