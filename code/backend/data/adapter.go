package data

import "io/fs"

type AssetAdapter struct {
	d      *DataBox
	folder string
}

func (a AssetAdapter) Open(name string) (fs.File, error) {
	return a.d.tryReadAssetsFile(a.folder, name)
}
