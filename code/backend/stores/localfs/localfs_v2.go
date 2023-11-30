package localfs

import (
	"archive/zip"
	"context"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/fdatautil"
)

// fixme => send generic error and log platfrom specific error

type LocalFS struct {
	basePath string
}

func (l *LocalFS) ListFolder(ctx context.Context, tenantId, fpath string) ([]*store.BlobInfo, error) {

	files, err := os.ReadDir(l.folderPath(fpath))
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

func (l *LocalFS) NewFolder(ctx context.Context, tenantId, fpath, name string) error {
	return os.MkdirAll(path.Join(l.basePath, fpath, name), 0755)
}

func (l *LocalFS) DeleteFolder(ctx context.Context, tenantId, fpath string) error {
	return os.RemoveAll(l.folderPath(fpath))
}

func (l *LocalFS) RenameFolder(ctx context.Context, tenantId, fpath, newname string) error {
	dir, _ := path.Split(fpath)

	// check if its a folder and exists

	return os.Rename(l.folderPath(fpath), path.Join(l.basePath, dir, newname))
}

func (l *LocalFS) CompressFolder(ctx context.Context, tenantId, fpath string) (store.FData, error) {

	folderToZip := l.folderPath(fpath)

	zfile, err := os.CreateTemp("", "*temphia_cab_folder.zip")
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return fdatautil.NewFromFile(path.Join(folderToZip, zfile.Name()), true), nil

}

func (l *LocalFS) TreeFolder(ctx context.Context, tenantId, fpath string) ([]*store.BlobInfo, error) {
	folderToZip := l.folderPath(fpath)
	respblobs := make([]*store.BlobInfo, 0)

	err := filepath.Walk(folderToZip, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		respblobs = append(respblobs, &store.BlobInfo{
			Name:         info.Name(),
			Size:         int(info.Size()),
			IsDir:        info.IsDir(),
			LastModified: info.ModTime().String(),
		})

		return nil

	})
	if err != nil {
		return nil, err
	}

	return respblobs, nil
}

func (l *LocalFS) GetFile(ctx context.Context, tenantId, fpath string) (store.FData, error) {

	ffile := l.folderPath(fpath)
	exist, err := l.fileExists(ffile)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, easyerr.NotFound("file")
	}

	return fdatautil.NewFromFile(ffile, false), nil
}

func (l *LocalFS) RenameFile(ctx context.Context, tenantId, fpath, name, newname string) error {

	ffile := l.filePath(fpath, name)

	exist, err := l.fileExists(ffile)
	if err != nil {
		return err
	}

	if !exist {
		return easyerr.NotFound("file")
	}

	return os.Rename(ffile, l.filePath(fpath, newname))
}

func (l *LocalFS) DuplicateFile(ctx context.Context, tenantId, fpath, name, newname string) error {

	ffile := l.filePath(fpath, name)

	exist, err := l.fileExists(ffile)
	if err != nil {
		return err
	}

	if !exist {
		return easyerr.NotFound("file")
	}

	srcFile, err := os.Open(ffile)
	if err != nil {
		return err
	}

	defer srcFile.Close()

	destFile, err := os.Create(l.filePath(fpath, newname))
	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return destFile.Sync()

}

func (l *LocalFS) MoveFile(ctx context.Context, tenantId, fpath, newfpath string) error {
	ffile := l.filePath(fpath, "")

	exist, err := l.fileExists(ffile)
	if err != nil {
		return err
	}

	if !exist {
		return easyerr.NotFound("file")
	}

	return os.Rename(ffile, l.folderPath(newfpath))
}

func (l *LocalFS) NewFile(ctx context.Context, tenantId, fpath, name string, data store.FData) error {
	return l.writeFile(ctx, fpath, name, data)
}

func (l *LocalFS) UpdateFile(ctx context.Context, tenantId, fpath, name string, data store.FData) error {
	return l.writeFile(ctx, fpath, name, data)
}

func (l *LocalFS) DeleteFile(ctx context.Context, tenantId, fpath, name string) error {
	return os.Remove(l.filePath(fpath, name))
}

func (l *LocalFS) CompressFiles(ctx context.Context, tenantId, fpath string, files []string) (store.FData, error) {

	zfile, err := os.CreateTemp("", "*temphia_cab_folder.zip")
	if err != nil {
		return nil, err
	}

	defer zfile.Close()

	zipWriter := zip.NewWriter(zfile)
	defer zipWriter.Close()

	cleanExist := false
	defer func() {
		if cleanExist {
			return
		}

		os.Remove(zfile.Name())
	}()

	for _, f := range files {

		ffile := l.filePath(fpath, f)

		zipEntry, err := zipWriter.Create(ffile)
		if err != nil {
			return nil, err
		}

		file, err := os.Open(ffile)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		_, err = io.Copy(zipEntry, file)
		if err != nil {
			return nil, err
		}
	}

	cleanExist = true

	return fdatautil.NewFromFile(zfile.Name(), true), nil
}

// private

func (l *LocalFS) writeFile(ctx context.Context, fpath, name string, data store.FData) error {

	defer data.Close()

	reader, err := data.AsReader()
	if err != nil {
		return err
	}

	destFile, err := os.Create(path.Join(l.basePath, fpath, name))
	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, reader)
	if err != nil {
		return err
	}

	return destFile.Sync()
}

func (l *LocalFS) filePath(folder, file string) string {
	return path.Join(l.basePath, folder, file)
}

func (l *LocalFS) folderPath(folder string) string {
	return path.Join(l.basePath, folder)
}

func (l *LocalFS) fileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return false, err
		}
	}

	return true, nil

}
