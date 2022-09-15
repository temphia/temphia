package runner

import "github.com/temphia/temphia/code/core/backend/xtypes"

type Runner struct {
	app xtypes.App
}

func New(app xtypes.App) *Runner {
	return &Runner{
		app: app,
	}
}

func (r *Runner) WithBasic() error {

	return nil
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
