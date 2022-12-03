package data

import (
	"io"
	"io/fs"
	"os"
	"path"
)

type Options struct {
	OverlayDataFolder        string
	OverlayBuildAssetsFolder string
	FS                       fs.FS
}

type DataBox struct {
	dataOverlay        string
	buildAssetsOverlay string
	fs                 fs.FS
}

func DefaultNew() *DataBox {
	if defaultDataBox == nil {
		lazyInit()
	}

	return defaultDataBox
}

func (d *DataBox) GetSchema(name string) ([]byte, error) {
	return d.tryRead(d.dataOverlay, "schema", name)
}

func (d *DataBox) GetTemplate(name string) ([]byte, error) {
	return d.tryRead(d.dataOverlay, "templates", name)
}

func (d *DataBox) GetIfaceFile(name string) ([]byte, error) {
	return d.tryRead(d.dataOverlay, "interfaces", name)
}

func (d *DataBox) GetAsset(folder, name string) ([]byte, error) {
	file, err := d.tryReadAssetsFile(folder, name)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(file)
}

func (d *DataBox) AssetAdapter(folder string) fs.FS {
	return &AssetAdapter{
		d:      d,
		folder: folder,
	}
}

func readFile(paths ...string) (fs.File, error) {
	return os.Open(path.Join(paths...))
}

// private

func (d *DataBox) tryReadAssetsFile(folder, name string) (fs.File, error) {
	if folder == "build" && d.buildAssetsOverlay != "" {
		file, err := readFile(d.buildAssetsOverlay, name)
		if err == nil {
			return file, nil
		}
	}

	return d.fs.Open(path.Join(d.dataOverlay, "assets", folder, name))
}

func (d *DataBox) tryRead(overlay, folder, name string) ([]byte, error) {
	file, err := d.tryReadFile(overlay, folder, name)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}

func (d *DataBox) tryReadFile(paths ...string) (fs.File, error) {

	if paths[0] != "" {
		bytes, err := readFile(paths...)
		if err == nil {
			return bytes, nil
		}
	}

	return d.fs.Open(path.Join(paths...))
}
