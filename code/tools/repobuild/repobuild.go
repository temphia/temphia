package repobuild

import (
	"github.com/temphia/temphia/code/distro/climux"
	"github.com/temphia/temphia/code/tools/repobuild/cmd"
)

func init() {

	climux.Register(&climux.CliAction{
		Name: "repo",
		Help: "repositary related actions",
		Func: func(args []string) error {

			if args[0] == "build" {
				return cmd.Run(".repo.yaml")
			}

			return nil
		},
	})

}
