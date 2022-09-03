package xtypes

import "io/fs"

type DataBox interface {
	GetSchema(name string) ([]byte, error)
	GetTemplate(name string) ([]byte, error)
	GetIfaceFile(name string) ([]byte, error)
	GetAsset(name string) ([]byte, error)
	AssetAdapter() fs.FS
}
