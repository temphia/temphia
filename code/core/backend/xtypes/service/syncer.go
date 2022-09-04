package service

import "github.com/temphia/temphia/code/core/backend/xtypes/store"

type Syncer interface {
	store.SyncDB
	GetInnerDriver() any
}
