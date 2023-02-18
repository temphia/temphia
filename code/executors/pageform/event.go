package pageform

import "github.com/k0kubun/pp"

func (pf *Pageform) actionLoad(req LoadRequest) (*Response, error) {

	return &Response{
		Ok:       true,
		Message:  "",
		Items:    pf.model.Items["start"],
		Data:     map[string]any{},
		OnLoad:   pf.model.ClientOnLoad,
		OnSubmit: pf.model.ClientOnSubmit,
		Stage:    "start",
	}, nil

}

func (pf *Pageform) actionSubmit(req SubmitRequest) (*Response, error) {

	pp.Println(req)

	message := ""
	ok := true
	items := pf.model.Items["second"]
	if items == nil {
		ok = false
		message = "stage not found"
	}

	return &Response{
		Ok:      ok,
		Message: message,
		Items:   items,
	}, nil
}
