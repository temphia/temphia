package pagedash

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

func (pd *PageDash) actionLoad(req LoadRequest) (*LoadResponse, error) {

	var data map[string]any

	if len(pd.model.OnLoad) == 0 {
		data = pd.model.StaticData
	} else {
		data = make(map[string]any)
		data["__exec_data__"] = req.ExecData
		for k, v := range pd.model.StaticData {
			data[k] = v
		}

		for _, hook := range pd.model.OnLoad {
			pctx := pd.new(data)
			pctx.bind()

			switch hook.Type {
			case "script":
				err := pctx.execute(hook.Target, req.Version)
				if err != nil {
					return nil, err
				}
			default:
				return nil, easyerr.NotImpl()
			}
		}
	}

	return &LoadResponse{
		Name:     pd.model.Name,
		Data:     data,
		Sections: pd.model.Sections,
	}, nil
}

func (pd *PageDash) actionBuild(req BuildRequest) (*BuildRespone, error) {

	return nil, nil
}
