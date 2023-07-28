package xtypes

import "io/fs"

type DataBox interface {
	GetSchema(name string) ([]byte, error)
	GetTemplate(name string) ([]byte, error)
	GetIfaceFile(name string) ([]byte, error)

	GetAsset(atype, name string) ([]byte, error)
	AssetAdapter(atype string) fs.FS
}
