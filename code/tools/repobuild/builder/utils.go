package builder

import (
	"crypto/sha1"
	"encoding/base64"
	"path"
)

func (rb *RepoBuilder) hashedBuidlPath(url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return path.Join(rb.config.BuildFolder, sha)
}
