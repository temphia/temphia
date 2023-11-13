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
		Func: func(cctx climux.Context) error {
			Run()
			return nil
		},
	})

}

func Run() {

	file := "temphia.json"

	ew := New()
	defer ew.Close()

	if !xutils.FileExists("./", file) {
		ew.RunWithStartPage(TemplateOptions{
			LocalRunning: false,
			LocalExists:  false,
			LocalFile:    "",
		})
		return
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
		return
	}

	ew.RunWithStartPage(TemplateOptions{
		LocalExists:  true,
		LocalRunning: true,
		LocalFile:    file,
	})

}
