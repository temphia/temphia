package easypage

import (
	"encoding/json"
	"fmt"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
)

func (e *EasyPage) serveEditorFile(file string) ([]byte, error) {

	switch file {
	case "main.js":
		return e.dataBox.GetAsset("build", "adapter_editor_easypage.js")
	case "main.css":
		return e.dataBox.GetAsset("build", "adapter_editor_easypage.css")
	}

	return []byte(``), nil
}

func (e *EasyPage) preformEditorAction(name string, data []byte) (any, error) {

	switch name {
	case "list_pages":
		return e.listPages(name)
	case "update_pages":
		return nil, e.updatePages(data)
	case "get_page_data":
		return e.getPageData(data)
	case "set_page_data":
		return nil, e.setPageData(data)
	case "delete_page_data":
		return nil, e.delPageData(data)
	default:
		return nil, easyerr.NotImpl()
	}
}

func (e *EasyPage) listPages(name string) ([]Page, error) {
	main, err := e.ahandle.KvGet("main")
	if err != nil {
		err2 := e.initState()
		if err2 != nil {
			return nil, err
		}
		return []Page{}, nil
	}

	out := make([]Page, 0)
	err = json.Unmarshal([]byte(main), &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (e *EasyPage) updatePages(data []byte) error {
	return e.ahandle.KvUpdate("main", string(data))
}

func (e *EasyPage) setPageData(data []byte) error {
	pd := NewPage{}
	err := json.Unmarshal(data, &pd)
	if err != nil {
		return err
	}

	err = e.ahandle.KvUpdate(pageKey(pd.Slug), pd.Data)
	if err != nil {
		err2 := e.ahandle.KvAdd(pageKey(pd.Slug), pd.Data)
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func (e *EasyPage) getPageData(data []byte) (string, error) {
	pslug := ""
	err := json.Unmarshal(data, &pslug)
	if err != nil {
		return "", err
	}

	return e.ahandle.KvGet(pageKey(pslug))
}

func (e *EasyPage) delPageData(data []byte) error {
	pslug := ""
	err := json.Unmarshal(data, &pslug)
	if err != nil {
		return err
	}

	return e.ahandle.KvRemove(pageKey(pslug))
}

func pageKey(slug string) string {
	return fmt.Sprintf("page-%s", slug)
}

func (e *EasyPage) initState() error {
	return e.ahandle.KvAdd("main", "[]")
}
