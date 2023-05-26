package dynddl2

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/dyncore"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/dynddl2/meta"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/upper/db/v4"
)

type DynDDL struct {
	session    db.Session
	sharedLock service.DyndbLock
	logger     zerolog.Logger
	meta       dyncore.DynMeta
}

func New(session db.Session, sharedLock service.DyndbLock, logger zerolog.Logger) *DynDDL {

	return &DynDDL{
		session:    session,
		sharedLock: sharedLock,
		logger:     logger,
		meta:       meta.New(session, logger),
	}
}

func (d *DynDDL) RunNew(tenantId string, migctx MigrateContext) error {
	return d.newGroup(tenantId, migctx.StmtString, migctx.BaseSchema)
}

func (d *DynDDL) RunUpdate(tenantId string, migctx MigrateContext) error {
	return d.update(tenantId, migctx)
}
