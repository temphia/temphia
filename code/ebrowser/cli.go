package ebrowser

import (
	"encoding/json"
	"os"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/distro/climux"
	"github.com/temphia/temphia/code/distro/common"
)

func init() {

	climux.Register(&climux.CliAction{
		Name: "ebrowser",
		Help: "Run embed browser with state folder",
		Func: RunCLI,
	})

	climux.DefaultCLI = "ebrowser"

}

func RunCLI(cctx climux.Context) error {

	file := "temphia.json"

	ew := New()
	defer ew.Close()

	ew.clictx = cctx

	if !xutils.FileExists("./", file) {
		ew.RunWithStartPage(TemplateOptions{
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

	if !common.IsRunning(conf.DataFolder) {
		ew.RunWithStartPage(TemplateOptions{
			LocalExists:  true,
			LocalRunning: false,
			LocalFile:    file,
		})

		return nil
	}

	ew.RunWithStartPage(TemplateOptions{
		LocalExists:  true,
		LocalRunning: true,
		LocalFile:    file,
	})

	return nil

}
