package agent

import (
	"fmt"
	"testing"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

func TestSPA(t *testing.T) {

	sb := SpaBuilder{
		opts: SpaBuilderOptions{
			Plug:         "test",
			Agent:        "default",
			APIBaseURL:   httpx.ApiBaseURL("example.com", xtypes.DefaultTenant),
			EntryName:    "loaderMain",
			ExecLoader:   "std.loader",
			TenantID:     xtypes.DefaultTenant,
			JsPlugScript: "main.js,lib1.mjs",
			StyleFile:    "main.css",
			ExtScripts:   make(map[string]interface{}),
		},
	}

	fmt.Println(string(sb.Build()))

}
