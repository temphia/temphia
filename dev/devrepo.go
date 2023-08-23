package dev

import (
	"embed"
	"io/fs"

	"github.com/temphia/temphia/code/backend/app/registry"

	// repo provider
	rembed "github.com/temphia/temphia/code/backend/services/pacman/provider/embed"
	_ "github.com/temphia/temphia/code/backend/services/pacman/provider/github"
	_ "github.com/temphia/temphia/code/backend/services/pacman/provider/local"
)

//go:embed devrepo
var DevRepo embed.FS

func init() {

	dfs, err := fs.Sub(DevRepo, "devrepo")
	if err != nil {
		panic(err)
	}

	registry.SetRepoBuilder("embed", rembed.NewEmbed("embed", dfs))

}
