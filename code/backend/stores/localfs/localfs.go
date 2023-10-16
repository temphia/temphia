package localfs

import (
	"archive/zip"
	"context"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var _ store.FileStore = (*NativeBlob)(nil)

var (
	ErrInvalidPath = errors.New("invalid path")
)

type NativeBlob struct {
	rootFilePath string
}

func (n *NativeBlob) AddFolder(ctx context.Context, tenant, folder string) error {
	err := n.isValid(folder)
	if err != nil {
		return err
	}

	return os.MkdirAll(n.FolderPath(folder), 0755)
}

func (n *NativeBlob) AddBlob(ctx context.Context, tenant, folder string, file string, contents []byte) error {

	err := n.isValid(folder)
	if err != nil {
		return err
	}

	err = n.isValid(file)
	if err != nil {
		return err
	}

	ok := xutils.FileExists(n.rootFilePath, folder)
	if !ok {
		err := n.AddFolder(ctx, tenant, folder)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(n.filePath(folder, file), contents, 0755)

	pp.Println("@add_blob", folder, file, len(contents))

	return err

}

func (n *NativeBlob) ListRoot(ctx context.Context, tenant string) ([]string, error) {
	files, err := os.ReadDir(n.rootFilePath)
	if err != nil {
		return nil, err
	}

	respFiles := make([]string, 0, len(files))
	for _, f := range files {
		respFiles = append(respFiles, f.Name())
	}
	return respFiles, nil
}

func (n *NativeBlob) ListFolderBlobs(ctx context.Context, tenant, folder string) ([]*store.BlobInfo, error) {
	err := n.isValid(folder)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(n.FolderPath(folder))
	if err != nil {
		return nil, err
	}

	respblobs := make([]*store.BlobInfo, 0, len(files))
	for _, f := range files {

		i, err := f.Info()
		if err != nil {
			continue
		}

		respblobs = append(respblobs, &store.BlobInfo{
			Name:         f.Name(),
			Size:         int(i.Size()),
			IsDir:        f.IsDir(),
			LastModified: i.ModTime().String(),
		})
	}

	return respblobs, nil
}

func (n *NativeBlob) GetBlob(ctx context.Context, tenant, folder string, file string) ([]byte, error) {
	err := n.isValid(folder)
	if err != nil {
		return nil, err
	}

	err = n.isValid(file)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(n.filePath(folder, file))
}

func (n *NativeBlob) DeleteBlob(ctx context.Context, tenant, folder string, file string) error {
	err := n.isValid(folder)
	if err != nil {
		return err
	}

	err = n.isValid(file)
	if err != nil {
		return err
	}

	return os.Remove(n.filePath(folder, file))
}

func (n *NativeBlob) GetFolderAsZip(ctx context.Context, tenant, folder string) (string, error) {
	err := n.isValid(folder)
	if err != nil {
		return "", err
	}

	folderToZip := n.FolderPath(folder)

	zfile, err := os.CreateTemp("", "*temphia_cab_folder.zip")
	if err != nil {
		return "", err
	}

	defer zfile.Close()

	zipWriter := zip.NewWriter(zfile)
	defer zipWriter.Close()

	err = filepath.Walk(folderToZip, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		zipEntry, err := zipWriter.Create(filePath)
		if err != nil {
			return err
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(zipEntry, file)
		return err
	})

	if err != nil {
		return "", err
	}

	return path.Join(folderToZip, zfile.Name()), nil

}

func (n *NativeBlob) FolderPath(folder string) string {
	return filepath.Join(n.rootFilePath, folder)
}

func (n *NativeBlob) filePath(folder, file string) string {
	return filepath.Join(n.rootFilePath, folder, file)
}

func (n *NativeBlob) AddBlobStreaming(ctx context.Context, tenant string, folder string, file string, contents io.ReadCloser) error {
	return easyerr.NotImpl()
}

func (n *NativeBlob) isValid(folder string) error {
	if strings.Contains(folder, "..") {
		return ErrInvalidPath
	}

	return nil
}
