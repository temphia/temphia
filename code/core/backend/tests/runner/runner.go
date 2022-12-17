package runner

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
)

type RunnerTest func(app xtypes.App) error

type Runner struct {
	app   xtypes.App
	tests map[string][]RunnerTest
}

func New(app xtypes.App) *Runner {
	return &Runner{
		app:   app,
		tests: make(map[string][]RunnerTest),
	}
}

func (r *Runner) Register(name string, tests []RunnerTest) {
	r.tests[name] = tests
}

/*

	add_tenant
	add_user_group
	add_super_user
	add_normal_user
		change_password
		remove_session
		list_session
	add_repo
		import_bprint
		instance_bprint -> plug
		execute

		instance_bprint -> data


	add_domain




*/
