package dev

import (
	"embed"
	"io/fs"

	// repo provider
	rembed "github.com/temphia/temphia/code/backend/services/pacman/provider/embed"
	_ "github.com/temphia/temphia/code/backend/services/pacman/provider/github"
)

//go:embed devrepo
var DevRepo embed.FS

func init() {

	dfs, err := fs.Sub(DevRepo, "devrepo")
	if err != nil {
		panic(err)
	}

	rembed.Register("embed", dfs)

}
