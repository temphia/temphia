package runner

import (
	"testing"

	epg "github.com/fergusstrange/embedded-postgres"
)

func TestApp(t *testing.T) {

	eargs := epg.NewDatabase(epg.DefaultConfig())

	eargs.Start()
	defer eargs.Stop()

}
