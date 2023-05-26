package migreflect

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/dyncore"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/upper/db/v4"
)

type MigReflect struct {
	session    db.Session
	sharedLock service.DyndbLock
	logger     zerolog.Logger
}

func New(session db.Session, sharedLock service.DyndbLock, logger zerolog.Logger) *MigReflect {
	return &MigReflect{
		session:    session,
		sharedLock: sharedLock,
		logger:     logger,
	}
}

func (d *MigReflect) dataTableGroups() db.Collection {
	return dyncore.GroupTable(d.session)
}

func (d *MigReflect) dataTables() db.Collection {
	return dyncore.Table(d.session)
}

func (d *MigReflect) dataTableColumns() db.Collection {
	return dyncore.TableColumn(d.session)
}
