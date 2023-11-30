package cabinethub

import (
	"context"
	"path"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

/*
	future
	- needsCache  bool
	- cacheBudget int64
*/

type CabinetHub struct {
	inner store.FileStore
}

func New(impl store.FileStore) *CabinetHub {
	return &CabinetHub{
		inner: impl,
	}
}

func (f *CabinetHub) Start(mb xplane.MsgBus) error {
	return nil
}

func (f *CabinetHub) ListFolder(ctx context.Context, tenantId, fpath string) ([]*store.BlobInfo, error) {

	if !valid(fpath) {
		return nil, easyerr.NotSupported()
	}

	return f.inner.ListFolder(ctx, tenantId, fpath)
}

func (f *CabinetHub) NewFolder(ctx context.Context, tenantId, fpath, name string) error {
	if !valid(fpath) || !valid(name) || strings.Contains(name, "/") {
		return easyerr.NotSupported()
	}

	return f.inner.NewFolder(ctx, tenantId, fpath, name)

}

func (f *CabinetHub) DeleteFolder(ctx context.Context, tenantId, fpath string) error {
	if !valid(fpath) {
		return easyerr.NotSupported()
	}

	return f.inner.DeleteFolder(ctx, tenantId, fpath)

}

func (f *CabinetHub) RenameFolder(ctx context.Context, tenantId, fpath, newname string) error {
	if !valid(fpath) || !valid(newname) || strings.Contains(newname, "/") {
		return easyerr.NotSupported()
	}

	return f.inner.RenameFolder(ctx, tenantId, fpath, newname)

}
func (f *CabinetHub) CompressFolder(ctx context.Context, tenantId, fpath string) (store.FData, error) {
	if !valid(fpath) {
		return nil, easyerr.NotSupported()
	}

	return f.inner.CompressFolder(ctx, tenantId, fpath)

}
func (f *CabinetHub) TreeFolder(ctx context.Context, tenantId, fpath string) ([]*store.BlobInfo, error) {
	if !valid(fpath) {
		return nil, easyerr.NotSupported()
	}

	return f.inner.TreeFolder(ctx, tenantId, fpath)

}

func (f *CabinetHub) GetFile(ctx context.Context, tenantId, fpath string) (store.FData, error) {
	if !valid(fpath) {
		return nil, easyerr.NotSupported()
	}

	return f.inner.GetFile(ctx, tenantId, fpath)

}

func (f *CabinetHub) RenameFile(ctx context.Context, tenantId, fpath, name, newname string) error {
	if !valid(fpath) || !valid(name) || !valid(newname) || strings.Contains(name, "/") || strings.Contains(newname, "/") {
		return easyerr.NotSupported()
	}

	return f.inner.RenameFile(ctx, tenantId, fpath, name, newname)

}

func (f *CabinetHub) DuplicateFile(ctx context.Context, tenantId, fpath, name, newname string) error {
	if !valid(fpath) || !valid(name) || !valid(newname) || strings.Contains(name, "/") || strings.Contains(newname, "/") {
		return easyerr.NotSupported()
	}

	return f.inner.DuplicateFile(ctx, tenantId, fpath, name, newname)
}

func (f *CabinetHub) MoveFile(ctx context.Context, tenantId, fpath, newfpath string) error {
	if !valid(fpath) || !valid(newfpath) {
		return easyerr.NotSupported()
	}

	return f.inner.MoveFile(ctx, tenantId, fpath, newfpath)
}

func (f *CabinetHub) NewFile(ctx context.Context, tenantId, fpath, name string, data store.FData) error {
	if !valid(fpath) || !valid(name) {
		return easyerr.NotSupported()
	}

	if strings.Contains(name, "/") {
		dir, file := path.Split(name)
		fpath = path.Join(fpath, dir)
		name = file
	}

	return f.inner.NewFile(ctx, tenantId, fpath, name, data)

}

func (f *CabinetHub) UpdateFile(ctx context.Context, tenantId, fpath, name string, data store.FData) error {
	if !valid(fpath) || !valid(name) {
		return easyerr.NotSupported()
	}

	if strings.Contains(name, "/") {
		dir, file := path.Split(name)
		fpath = path.Join(fpath, dir)
		name = file
	}

	return f.inner.UpdateFile(ctx, tenantId, fpath, name, data)
}

func (f *CabinetHub) DeleteFile(ctx context.Context, tenantId, fpath, name string) error {
	if !valid(fpath) || !valid(name) {
		return easyerr.NotSupported()
	}

	if strings.Contains(name, "/") {
		dir, file := path.Split(name)
		fpath = path.Join(fpath, dir)
		name = file
	}

	return f.inner.DeleteFile(ctx, tenantId, fpath, name)

}

func (f *CabinetHub) CompressFiles(ctx context.Context, tenantId, fpath string, files []string) (store.FData, error) {
	if !valid(fpath) {
		return nil, easyerr.NotSupported()
	}

	for _, f := range files {
		if !valid(f) {
			return nil, easyerr.NotAuthorized()
		}

	}

	return f.inner.CompressFiles(ctx, tenantId, fpath, files)

}

// private

func valid(fpath string) bool {
	return !strings.Contains(fpath, "..")
}
