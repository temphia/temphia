package builder

import (
	"crypto/sha1"
	"encoding/base64"
	ppath "path"

	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

func (rb *RepoBuilder) hashedBuidlPath(url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return ppath.Join(rb.config.BuildFolder, sha)
}

func ZipIt(bprint *xpackage.Manifest, outFile string) error {

	z, err := NewZipper(bprint, outFile)
	if err != nil {
		return err
	}

	return z.Build()

}
