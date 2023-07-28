package dyndb

import (
	"strings"

	"github.com/thoas/go-funk"
)

const (
	// meta keys
	KeyPrimary = "__id"
	KeyVersion = "__version"
	KeyModSig  = "__mod_sig"

	// meta reference keys
	KeyForceVersion     = "__force_version__"
	KeySecondary        = "__secondary_keys__"
	KeyErrorAfterUpdate = "__error_after_update__"
)

const (
	DynActivityTypeNone   = "none"
	DynActivityTypeStrict = "strict"
	DynActivityTypeLazy   = "lazy"
)

const (
	DynSyncTypeNone         = "none"
	DynSyncTypeEventOnly    = "event_only"
	DynSyncTypeEventAndData = "event_and_data"
)

const (
	DynSeedTypeData    = "data"
	DynSeedTypeAutogen = "autogen"
)

const (
	DefaultSeedNo          = 10
	DefaultQueryFetchCount = 10
)

var (
	MetaRefKeys = []string{KeyForceVersion, KeySecondary, KeyErrorAfterUpdate}
	MetaKeys    = []string{KeyPrimary, KeyVersion, KeyModSig}
)

func IsMeta(name string) bool {
	if !strings.HasPrefix(name, "__") {
		return false
	}

	return funk.ContainsString(MetaKeys, name)
}
