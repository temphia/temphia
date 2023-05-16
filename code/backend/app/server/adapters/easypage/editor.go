package easypage

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
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

func (e *EasyPage) preformEditorAction(uclaim *claim.UserContext, name string, data []byte) (any, error) {
	if name != "load" {
		return nil, easyerr.NotFound("editor perform action")
	}

	return e.signer.SignPlugState(e.options.TenantId, &claim.PlugState{
		TenantId:  e.options.TenantId,
		UserId:    uclaim.UserID,
		DeviceId:  uclaim.DeviceId,
		SessionId: uclaim.SessionID,
		ExecId:    0,
		PlugId:    e.editorHook.PlugId,
		AgentId:   e.editorHook.AgentId,
	})
}
