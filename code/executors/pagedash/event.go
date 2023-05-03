package pagedash

import (
	"errors"
	"fmt"
	"time"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/upper/db/v4"
)

func (pd *PageDash) actionLoad(req LoadRequest) (*LoadResponse, error) {

	idx, err := pd.getIndex()
	if err != nil && !errors.Is(err, db.ErrNoMoreRows) {
		return nil, err
	}
	if idx == nil {
		idx = &Index{}
	}

	if req.Id == "" {
		req.Id = idx.ActiveKey
	}

	data, err := pd.runLoadHook(req.ExecData, req.Id)
	if err != nil {
		return nil, err
	}

	return &LoadResponse{
		Name:     pd.model.Name,
		Data:     data,
		Sections: pd.model.Sections,

		Sources:      pd.model.Sources,
		ActiveKey:    idx.ActiveKey,
		BuildDate:    idx.BuildDate,
		PreviousKeys: idx.PreviousKeys,
	}, nil
}

func (pd *PageDash) runLoadHook(execData any, key string) (map[string]any, error) {
	var data map[string]any
	if key != "" {
		_data, err := pd.getBuildData(key)
		if err != nil {
			return nil, err
		}
		data = _data
		data["__build_ok_"] = true
	}

	if data == nil {
		data = map[string]any{}
	}

	if len(pd.model.OnLoad) == 0 {
		for k, v := range pd.model.StaticData {
			data[k] = v
		}
	}

	data["__exec_data__"] = execData
	for k, v := range pd.model.StaticData {
		data[k] = v
	}

	pctx := pd.new(data)
	pctx.bind()

	for _, hook := range pd.model.OnLoad {

		switch hook.Type {
		case "script":
			err := pctx.execute(hook.Target, key)
			if err != nil {
				return nil, err
			}
		case "func":
			hfunc, ok := hookFuncs[hook.Target]
			if !ok {
				return nil, easyerr.NotFound("hookfunc")
			}

			err := hfunc(pctx)
			if err != nil {
				return nil, err
			}

		default:
			return nil, easyerr.NotImpl()
		}
	}

	return data, nil

}

func (pd *PageDash) actionBuild(req BuildRequest) (*BuildRespone, error) {

	idx, err := pd.getIndex()
	if err != nil && !errors.Is(err, db.ErrNoMoreRows) {
		return nil, err
	}

	if err != nil {
		idx = &Index{
			ActiveKey:    "",
			BuildDate:    "",
			PreviousKeys: make(map[string]string),
			Version:      0,
		}
	}

	nexid, err := xutils.GenerateRandomString(10)
	if err != nil {
		return nil, err
	}

	idx.PreviousKeys[idx.ActiveKey] = idx.BuildDate
	idx.ActiveKey = nexid
	idx.BuildDate = time.Now().Format(time.RFC850)

	pctx := pd.new(map[string]any{})
	pctx.bind()

	for _, buildHook := range pd.model.OnBuild {

		switch buildHook.Type {
		case "script":
			err := pctx.execute(buildHook.Target, idx.ActiveKey)
			if err != nil {
				return nil, err
			}
		case "func":
			hfunc, ok := hookFuncs[buildHook.Target]
			if !ok {
				return nil, easyerr.NotFound("hookfunc")
			}

			err := hfunc(pctx)
			if err != nil {
				return nil, err
			}

		default:
			return nil, easyerr.NotImpl()
		}
	}

	err = pd.setBuildData(idx.ActiveKey, pctx.Data)
	if err != nil {
		pd.binder.Log(fmt.Sprintf("could not set build data %v", idx))
		return nil, err
	}

	err = pd.setIndex(idx)
	if err != nil {
		pd.binder.Log(fmt.Sprintf("could not update index %v", idx))
		pd.delBuldData(idx.ActiveKey)
	}

	return &BuildRespone{
		Id: idx.ActiveKey,
	}, nil
}
