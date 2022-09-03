package data

import (
	"io/fs"
)

type AssetsFSAdapter struct {
	as AssetStore
}

func (a *AssetStore) AssetAdapter() fs.FS {
	return &AssetsFSAdapter{
		as: *a,
	}
}

func (a *AssetsFSAdapter) Open(name string) (fs.File, error) {
	return a.as.tryReadFile(a.as.assetsFolder, "assets", name)
}
