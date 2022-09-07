package web

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
)

type request struct {
	Name string `json:"name,omitempty"`
}

func getTargetName(data xtypes.LazyData) (string, error) {
	req := &request{}
	err := data.AsObject(req)
	if err != nil {
		return "", err
	}
	return req.Name, nil
}
