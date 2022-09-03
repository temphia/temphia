package local

import (
	"github.com/temphia/core/data"
	"github.com/temphia/temphia/code/core/backend/app/registry"
)

func init() {
	registry.SetRepoBuilder("embed", data.NewEmbed)
}
