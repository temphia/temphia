package runner

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/tests"
	"github.com/temphia/temphia/code/core/backend/xtypes"
)

type Runner struct {
	handle *tests.TestHandle
	groups []tests.TestGroup
}

func New(app xtypes.App) *Runner {

	handle := &tests.TestHandle{
		App:     app,
		CoreHub: nil,
		CabHub:  nil,
		XPlane:  nil,
	}

	builders := tests.GlobalGet()

	groups := make([]tests.TestGroup, 0)

	for _, tgb := range builders {
		groups = append(groups, tgb(handle))
	}

	return &Runner{
		handle: handle,
		groups: groups,
	}
}

func (r *Runner) Run() error {

	for _, tg := range r.groups {
		fmt.Printf("Running test group |> %s \n", tg.Name())
		cases := tg.Cases()

		for _, tcase := range cases {
			fmt.Printf("\t Running case |> %s/%s \n", tg.Name(), tcase)
			err := tg.Run(tcase)
			if err == nil {
				fmt.Printf("Test run OK \n")
			} else {
				fmt.Printf("Test run error |> \n %s \n", err.Error())
			}
		}
	}

	return nil

}
