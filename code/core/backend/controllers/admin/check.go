package admin

import (
	"errors"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/upper/db/v4"
)

func (c *Controller) CheckBprint(uclaim *claim.Session, bid string) error {
	_, err := c.coredb.BprintGet(uclaim.TenentId, bid)
	if err == nil {
		return easyerr.Error("Already Exists")
	}

	if errors.Is(err, db.ErrNoMoreRows) {
		return nil
	}

	return err
}

func (c *Controller) CheckPlug(uclaim *claim.Session, pid string) error {
	_, err := c.coredb.PlugGet(uclaim.TenentId, pid)
	if err == nil {
		return easyerr.Error("Already Exists")
	}

	if errors.Is(err, db.ErrNoMoreRows) {
		return nil
	}
	return nil
}

func (c *Controller) CheckDataGroup(uclaim *claim.Session, source, gid string) error {

	return nil
}

func (c *Controller) CheckDataTable(uclaim *claim.Session, source, gid, tid string) error {
	return nil
}
