package erepo

import (
	"embed"
	"io/fs"

	// repo provider
	rembed "github.com/temphia/temphia/code/backend/services/pacman/provider/embed"
	_ "github.com/temphia/temphia/code/backend/services/pacman/provider/github"
)

//go:embed data
var Repo embed.FS

func init() {

	dfs, err := fs.Sub(Repo, "data")
	if err != nil {
		panic(err)
	}

	rembed.Register("embed", dfs)

}
