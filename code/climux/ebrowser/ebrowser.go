package ebrowser

import (
	"encoding/json"
	"os"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/climux"
	"github.com/temphia/temphia/code/distro"
	eb "github.com/temphia/temphia/code/ebrowser"
)

func init() {

	climux.Register(&climux.Action{
		Name: "ebrowser",
		Help: "Run embed browser with state folder",
		Func: RunCLI,
	})

	climux.DefaultCLI = "ebrowser"

}

func RunCLI(cctx climux.Context) error {

	file := "temphia.json"

	ew := eb.New(cctx)
	defer ew.Close()

	if !xutils.FileExists("./", file) {
		ew.RunWithStartPage(eb.TemplateOptions{
			LocalRunning: false,
			LocalExists:  false,
			LocalFile:    "",
		})
		return nil
	}

	out, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	conf := &config.Config{}
	err = json.Unmarshal(out, conf)
	if err != nil {
		panic(err)
	}

	if !distro.IsRunning(conf.DataFolder) {
		ew.RunWithStartPage(eb.TemplateOptions{
			LocalExists:  true,
			LocalRunning: false,
			LocalFile:    file,
		})

		return nil
	}

	ew.RunWithStartPage(eb.TemplateOptions{
		LocalExists:  true,
		LocalRunning: true,
		LocalFile:    file,
	})

	return nil

}
