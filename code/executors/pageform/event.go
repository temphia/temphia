package pageform

func (pf *Pageform) actionLoad(req LoadRequest) (*Response, error) {

	return &Response{
		Ok:       true,
		Message:  "",
		Items:    pf.model.Items["start"],
		Data:     map[string]any{},
		OnLoad:   pf.model.ClientOnLoad,
		OnSubmit: pf.model.ClientOnSubmit,
	}, nil

}

func (pf *Pageform) actionSubmit(req SubmitRequest) (*Response, error) {

	return &Response{
		Ok:      true,
		Message: "",
		Items:   pf.model.Items["aa"],
	}, nil
}
