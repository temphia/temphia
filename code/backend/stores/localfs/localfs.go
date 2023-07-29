package localfs

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var _ store.FileStore = (*NativeBlob)(nil)

type NativeBlob struct {
	rootFilePath string
}

func (n *NativeBlob) InitilizeTenent(tenant string, folders []string) error {

	for _, folder := range folders {
		err := n.AddFolder(context.TODO(), tenant, folder)
		log.Println(err)
	}

	return nil
}

func (n *NativeBlob) AddFolder(ctx context.Context, tenant, folder string) error {
	return os.Mkdir(n.FolderPath(tenant, folder), 0755)
}

func (n *NativeBlob) AddBlob(ctx context.Context, tenant, folder string, file string, contents []byte) error {

	ok := xutils.FileExists(n.rootFilePath, folder)
	if !ok {
		err := n.AddFolder(ctx, tenant, folder)
		if err != nil {
			return err
		}
	}

	err := os.WriteFile(n.filePath(tenant, folder, file), contents, 0755)

	pp.Println(err, contents)

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
	files, err := ioutil.ReadDir(n.FolderPath(tenant, folder))
	if err != nil {
		return nil, err
	}
	respblobs := make([]*store.BlobInfo, 0, len(files))
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		respblobs = append(respblobs, &store.BlobInfo{
			Name:         f.Name(),
			Size:         int(f.Size()),
			IsDir:        f.IsDir(),
			LastModified: f.ModTime().String(),
		})
	}

	return respblobs, nil
}

func (n *NativeBlob) GetBlob(ctx context.Context, tenant, folder string, file string) ([]byte, error) {
	return os.ReadFile(n.filePath(tenant, folder, file))
}

func (n *NativeBlob) DeleteBlob(ctx context.Context, tenant, folder string, file string) error {
	return os.Remove(n.filePath(tenant, folder, file))
}

func (n *NativeBlob) FolderPath(tenant, folder string) string {
	return filepath.Join(n.rootFilePath, tenant, folder)
}

func (n *NativeBlob) filePath(tenant, folder, file string) string {
	return filepath.Join(n.rootFilePath, tenant, folder, file)
}

func (n *NativeBlob) AddBlobStreaming(ctx context.Context, tenant string, folder string, file string, contents io.ReadCloser) error {
	return easyerr.NotImpl()
}

// func (n *NativeBlob) CheckPreSignedReadToken(ctx context.Context, tenant string, token string) error {
// 	return easyerr.NotImpl()
// }
// func (n *NativeBlob) GetPreSignedReadToken(ctx context.Context, tenant, folder string, file string) (string, error) {
// 	return "", easyerr.NotImpl()
// }
// func (n *NativeBlob) GetPreSignedWriteToken(ctx context.Context, tenant, folder string, file string) (string, error) {
// 	return "", easyerr.NotImpl()
// }