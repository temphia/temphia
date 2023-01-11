package local

import (
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/data"
)

func init() {
	registry.SetRepoBuilder("embed", data.NewEmbed)
}
