package cases

import (
	"github.com/temphia/temphia/code/backend/tests"
)

func init() {
	tests.Add(RegisterRepo)
	tests.Add(RegisterTenant)
}

func RegisterRepo(h *tests.TestHandle) tests.TestGroup {
	grp := tests.NewGroup("repo", h)

	grp.AddCase("add_repo", addRepo)

	return grp
}

func addRepo(th *tests.TestHandle) error {

	return nil
}
