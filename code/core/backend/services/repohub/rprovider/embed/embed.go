package local

import (
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/data"
)

func init() {
	registry.SetRepoBuilder("embed", data.NewEmbed)
}
