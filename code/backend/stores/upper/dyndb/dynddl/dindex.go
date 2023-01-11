package dynddl

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

// fixme => finish/check all meta types

func (d *DynDDL) AddIndex(tenantId, gslug, tslug string, model *entities.Index) error {
	qstr, err := d.dyngen.AddIndex(tenantId, gslug, tslug, model.Slug, "model.Mtype", model.Spans)
	if err != nil {
		return err
	}
	return d.GroupExecute(tenantId, gslug, qstr)
}

func (d *DynDDL) GroupExecute(tenantId, gslug, qstr string) error {
	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		return err
	}

	defer d.sharedLock.GlobalUnLock(tenantId, utok)

	return dbutils.Execute(ucore.GetDriver(d.session), qstr)
}

func (d *DynDDL) AddUniqueIndex(tenantId, gslug, tslug string, model *entities.Index) error {
	qstr, err := d.dyngen.AddIndex(tenantId, gslug, tslug, model.Slug, "model.Mtype", model.Spans)
	// fixme => uniue ?
	if err != nil {
		return err
	}
	pp.Println(qstr)

	return easyerr.NotImpl()

	// return d.GroupExecute(tenantId, gslug, qstr)
}

func (d *DynDDL) AddFTSIndex(tenantId, gslug, tslug string, model *entities.FTSIndex) error {

	// fixme

	// d.dyngen.AddIndex()

	// qstr, err := d.devend.AddFTSIndex(model.Spans)
	// if err != nil {
	// 	return err
	// }
	// return d.GroupExecute(tenantId, gslug, qstr)
	return easyerr.NotImpl()
}

func (d *DynDDL) AddColumnFRef(tenantId, gslug, tslug string, model *entities.ColumnFKRef) error {
	// qstr, err := d.devend.AddFKRef(model)
	// if err != nil {
	// 	return err
	// }
	// return d.GroupExecute(tenantId, gslug, qstr)
	return easyerr.NotImpl()
}

func (d *DynDDL) ListIndex(tenantId, gslug, tslug string) ([]*entities.Index, error) {

	// d.dyngen.GetFKRefs(tenantId, gslug, tslug, func(query string) (map[string]interface{}, error) {
	// 	return nil, nil
	// })

	return nil, nil
}

func (d *DynDDL) ListFKRef(tenantId, gslug, tslug string) ([]*entities.Index, error) {
	// qstr, err := d.devend.ListFKRef()
	// if err != nil {
	// 	return nil, err
	// }

	// rows, err := d.session.SQL().Query(qstr)
	// if err != nil {
	// 	return nil, err
	// }
	// return d.devend.ExtractIndexs(rows)

	return nil, nil
}

func (d *DynDDL) RemoveIndex(tenantId, gslug, tslug, slug string) error {
	// qstr, err := d.devend.RemoveIndex(slug)
	// if err != nil {
	// 	return err
	// }
	// return d.GroupExecute(tenantId, gslug, qstr)
	return easyerr.NotImpl()
}

func (d *DynDDL) RemoveFKRef(tenantId, gslug, tslug, refslug string) error {
	// qstr, err := d.devend.RemoveFKRef(refslug)
	// if err != nil {
	// 	return err
	// }
	// return d.GroupExecute(tenantId, gslug, qstr)

	return easyerr.NotImpl()
}
