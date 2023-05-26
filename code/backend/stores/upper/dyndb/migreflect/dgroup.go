package migreflect

import (
	"log"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"

	"github.com/upper/db/v4"
)

func (d *MigReflect) NewGroup(tenantId, stmt string, model *xbprint.NewTableGroup) error {

	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		return err
	}

	defer d.sharedLock.GlobalUnLock(tenantId, utok)

	err = d.newGroupRef(tenantId, model)
	if err != nil {
		return err
	}

	pp.Println("METADATA TABLE CREATED... now executing schema string")

	err = dbutils.Execute(ucore.GetDriver(d.session), stmt)
	if err != nil {
		d.rollbackGroupRef(tenantId, model.Slug)
	}

	return err
}

func (d *MigReflect) EditGroup(tenantId string, gslug string, model *entities.TableGroupPartial) error {
	return d.dataTableGroups().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      gslug,
	}).Update(model)
}

func (d *MigReflect) GetGroup(tenantId, gslug string) (*entities.TableGroup, error) {
	tg := &entities.TableGroup{}

	err := d.dataTableGroups().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      gslug,
	}).One(tg)

	return tg, err
}

func (d *MigReflect) ListGroup(tenantId string) ([]*entities.TableGroup, error) {
	tgs := make([]*entities.TableGroup, 0, 10)
	err := d.dataTableGroups().Find("tenant_id", tenantId).All(&tgs)
	return tgs, err
}

func (d *MigReflect) DeleteGroup(tenantId, gslug string) error {

	d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).Delete()

	d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).Delete()

	return d.dataTableGroups().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      gslug,
	}).Delete()

}

func (d *MigReflect) newGroupRef(tenantId string, model *xbprint.NewTableGroup) (err error) {

	clear := false
	defer func() {
		if clear {
			d.rollbackGroupRef(tenantId, model.Slug)
		}
	}()

	_, err = d.dataTableGroups().Insert(model.To(tenantId))
	if err != nil {
		return
	}

	clear = true

	for _, tbl := range model.Tables {
		err = d.AddTableRef(tenantId, model.Slug, tbl)
		if err != nil {
			pp.Println(err)
			return err
		}
	}
	clear = false

	return
}

func (d *MigReflect) rollbackGroupRef(tenantId string, gslug string) error {
	log.Println("ROLLING BACK GROUP....", tenantId, gslug)

	err := d.DeleteGroup(tenantId, gslug)
	if err != nil {
		log.Println(err)
	}
	return err
}
