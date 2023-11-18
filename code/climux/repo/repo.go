package repo

import (
	"github.com/temphia/temphia/code/climux"
	"github.com/temphia/temphia/code/tools/repobuild/cmd"
)

func init() {

	climux.Register(&climux.Action{
		Name: "repo",
		Help: "repositary related actions",
		Func: func(cctx climux.Context) error {

			if cctx.Args[0] == "build" {
				return cmd.Run(".repo.yaml")
			}

			return nil
		},
	})

}
