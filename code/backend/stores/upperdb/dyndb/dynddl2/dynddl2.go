package dynddl2

import (
	"github.com/rs/zerolog"

	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/upper/db/v4"
)

type DynDDL struct {
	session    db.Session
	sharedLock service.DyndbLock
	logger     zerolog.Logger
}

func New(session db.Session, sharedLock service.DyndbLock, logger zerolog.Logger) *DynDDL {

	return &DynDDL{
		session:    session,
		sharedLock: sharedLock,
		logger:     logger,
	}
}

func (d *DynDDL) RunNew(tenantId string, migctx MigrateContext) error {
	return d.runNew(tenantId, migctx)
}

func (d *DynDDL) RunUpdate(tenantId string, migctx MigrateContext) error {
	return d.update(tenantId, migctx)
}
