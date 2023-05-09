package easypage

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

func (e *EasyPage) serveEditorFile(file string) ([]byte, error) {

	pp.Println("@server_easypage", file)

	switch file {
	case "main.js":
		return e.dataBox.GetAsset("build", "adapter_editor_easypage.js")
	case "main.css":
		return e.dataBox.GetAsset("build", "adapter_editor_easypage.css")
	default:
		return os.ReadFile(fmt.Sprintf("code/frontend/public/build/%s", file))
	}

}

func (e *EasyPage) preformEditorAction(name string, data []byte) (any, error) {

	switch name {
	// case "load":
	// 	return e.load()
	// case "update_pages":
	// 	return nil, e.updatePages(data)
	// case "get_page_data":
	// 	return e.getPageData(data)
	// case "set_page_data":
	// 	return nil, e.setPageData(data)
	// case "delete_page_data":
	// 	return nil, e.delPageData(data)
	default:
		return nil, easyerr.NotImpl()
	}
}
