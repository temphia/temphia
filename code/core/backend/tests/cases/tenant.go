package cases

import "github.com/temphia/temphia/code/core/backend/tests"

func RegisterTenant(h *tests.TestHandle) tests.TestGroup {

	grp := tests.NewGroup("tenant", h)

	grp.AddCase("add_tenant", func(th *tests.TestHandle) error {
		return nil
	})

	return grp

}
