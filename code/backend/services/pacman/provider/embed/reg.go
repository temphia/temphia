package embed

import (
	"io/fs"

	"github.com/temphia/temphia/code/backend/app/registry"
)

func Register(name string, fsys fs.FS) {
	registry.SetRepoBuilder(name, NewEmbed(name, fsys))
}
