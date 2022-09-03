package cabinethub

import (
	"fmt"
	"strings"

	"github.com/thoas/go-funk"
)

const previewPrefix = "__preview__"

var (
	imageFileExt = []string{"jpeg", "png", "gif", "jpg"}
)

func previewName(filename string) string {
	return fmt.Sprintf("%s%s", previewPrefix, filename)
}

func canPreview(file string) bool {

	fileSegs := strings.Split(file, ".")
	if len(fileSegs) < 2 {
		return false
	}
	return funk.ContainsString(imageFileExt, fileSegs[1])
}

// func (b *blobAdapter) GeneratePreview(ctx context.Context, tenant, folder string, file string, contents []byte) error {
// 	pbytes, err := image.GeneratePreview(contents, file)
// 	if err != nil {
// 		return err
// 	}
// 	return b.BlobProvider.AddBlob(ctx, tenant, folder, file, pbytes)
// }

// func (b *blobAdapter) AddBlob(ctx context.Context, tenant, folder string, file string, contents []byte) error {

// 	err := b.BlobProvider.AddBlob(ctx, tenant, folder, file, contents)
// 	if err != nil {
// 		return err
// 	}
// 	if !b.canPreview(file) {
// 		return nil
// 	}

// 	err = b.GeneratePreview(ctx, tenant, folder, file, contents)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return nil
// }

// func (b *blobAdapter) GeneratePreview(ctx context.Context, tenant, folder string, file string, contents []byte) error {
// 	pbytes, err := image.GeneratePreview(contents, file)
// 	if err != nil {
// 		return err
// 	}
// 	return b.BlobProvider.AddBlob(ctx, tenant, folder, file, pbytes)
// }

// func (b *blobAdapter) AddBlobStreaming(ctx context.Context, tenant string, folder string, file string, contents io.ReadCloser) error {
// 	if !b.canPreview(file) {
// 		return b.BlobProvider.AddBlobStreaming(ctx, tenant, folder, file, contents)
// 	}

// 	defer contents.Close()

// 	bytes, err := ioutil.ReadAll(contents)
// 	if err != nil {
// 		return err
// 	}

// 	err = b.GeneratePreview(ctx, tenant, folder, file, bytes)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return nil

// }
