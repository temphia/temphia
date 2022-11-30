package data

import (
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/k0kubun/pp"
	"gitlab.com/mr_balloon/golib"
)

var (
	defaultAssetStore *AssetStore
)

func lazyInit() {
	var frontendBuild = "code/core/frontend/public/build"
	if exits, _ := golib.FileExists(frontendBuild); !exits {
		frontendBuild = ""
	}

	pp.Println("Using overlay assets from", frontendBuild)

	defaultAssetStore = &AssetStore{
		dataFolder:     "",
		schemaFolder:   "",
		templateFolder: "",
		ifaceFolder:    "",
		assetsFolder:   frontendBuild,
		fallbackFS:     dataDir,
	}

}

type Options struct {
	OverlayDataFolder      string
	OverlaySchemaFolder    string
	OverlayTemplateFolder  string
	OverlayInterfaceFolder string
	OverlayAssetsFolder    string
	FS                     fs.FS
}

// AssetStore tries to read from specific overlay folder
// (OverlaySchemaFolder, OverlayTemplateFolder, OverlayBuildFolder)
// if its not given then try to read from global overlay (OverlayDataFolder)
// if it still cannot read atlast read from embed fs
type AssetStore struct {
	dataFolder     string
	schemaFolder   string
	templateFolder string
	assetsFolder   string
	ifaceFolder    string
	fallbackFS     fs.FS
}

func DefaultNew() *AssetStore {
	if defaultAssetStore == nil {
		lazyInit()
	}

	return defaultAssetStore
}

func New(opts Options) *AssetStore {
	ass := &AssetStore{
		dataFolder:     opts.OverlayDataFolder,
		schemaFolder:   opts.OverlaySchemaFolder,
		templateFolder: opts.OverlayTemplateFolder,
		assetsFolder:   opts.OverlayAssetsFolder,
		ifaceFolder:    opts.OverlayInterfaceFolder,
		fallbackFS:     opts.FS,
	}

	if ass.fallbackFS == nil {
		ass.fallbackFS = dataDir
	}

	return ass
}

func (a *AssetStore) GetSchema(name string) ([]byte, error) {
	return a.tryRead(a.schemaFolder, "schema", name)
}

func (a *AssetStore) GetIfaceFile(name string) ([]byte, error) {
	return a.tryRead(a.ifaceFolder, "interfaces", name)
}

func (a *AssetStore) GetTemplate(name string) ([]byte, error) {
	return a.tryRead(a.templateFolder, "templates", name)
}

func (a *AssetStore) GetAsset(name string) ([]byte, error) {
	return a.tryRead(a.assetsFolder, "assets", name)
}

func (a *AssetStore) tryRead(overlay, folder, name string) ([]byte, error) {
	file, err := a.tryReadFile(overlay, folder, name)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}

func (a *AssetStore) tryReadFile(overlay, folder, name string) (fs.File, error) {
	if overlay != "" {
		bytes, err := readFile(overlay, name)
		if err == nil {
			return bytes, nil
		}
	}

	subFolder := path.Join(folder, name)
	if a.dataFolder != "" {
		bytes, err := readFile(a.dataFolder, subFolder)
		if err == nil {
			return bytes, nil
		}
	}
	return a.fallbackFS.Open(subFolder)
}

func readFile(folder, file string) (fs.File, error) {
	return os.Open(path.Join("./", folder, file))
}
