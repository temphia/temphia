package cabinethub

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

// fixme => adapter stuff like preview generation and stuff

type cabinetSourced struct {
	source   string
	tenantId string
	provider store.CabinetSource
}

func (c *cabinetSourced) AddFolder(ctx context.Context, folder string) error {
	if folder == "" {
		return easyerr.NotFound()
	}

	return c.provider.AddFolder(ctx, c.tenantId, folder)
}

func (c *cabinetSourced) AddBlob(ctx context.Context, folder, file string, contents []byte) error {
	if folder == "" || file == "" {
		return easyerr.NotFound()
	}

	return c.provider.AddBlob(ctx, c.tenantId, folder, file, contents)
}

func (c *cabinetSourced) ListRoot(ctx context.Context) ([]string, error) {
	return c.provider.ListRoot(ctx, c.tenantId)
}

func (c *cabinetSourced) ListFolder(ctx context.Context, folder string) ([]*store.BlobInfo, error) {
	if folder == "" {
		return nil, easyerr.NotFound()
	}
	return c.provider.ListFolderBlobs(ctx, c.tenantId, folder)
}

func (c *cabinetSourced) GetBlob(ctx context.Context, folder, file string) ([]byte, error) {
	if folder == "" || file == "" {
		return nil, easyerr.NotFound()
	}
	return c.provider.GetBlob(ctx, c.tenantId, folder, file)
}

func (c *cabinetSourced) DeleteBlob(ctx context.Context, folder, file string) error {
	if folder == "" || file == "" {
		return easyerr.NotFound()
	}

	return c.provider.DeleteBlob(ctx, c.tenantId, folder, file)
}
