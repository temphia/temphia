package dyndb

import (
	"encoding/json"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/upper/db/v4"
)

type (
	migrateContext struct {
		baseSchema *xbprint.NewTableGroup
		stmtString string
		postItems  []postDDLItem
		siblings   map[string]map[string]string
	}

	postDDLItem struct {
		mtype string
		data  any
	}
)

func (d *DynDB) migrateSchema(tenantId string, opts step.MigrateOptions) error {

	var baseSchema *xbprint.NewTableGroup
	var buf strings.Builder

	postitems := make([]postDDLItem, 0)

	siblings := make(map[string]map[string]string)

	if opts.New {
		firstStep := opts.Steps[0]

		if firstStep.Type != step.MigTypeNewGroup {
			return easyerr.Error("wrong type as first migration step")
		}

		baseSchema := &xbprint.NewTableGroup{}

		err := json.Unmarshal(opts.Steps[0].Data, baseSchema)
		if err != nil {
			return err
		}

		stmt, err := d.dyngen.NewGroup(tenantId, baseSchema)
		if err != nil {
			return err
		}

		buf.WriteString(stmt.String())
		opts.Steps = opts.Steps[1:]

		postitems = append(postitems, postDDLItem{
			mtype: firstStep.Type,
			data:  baseSchema,
		})

		for _, table := range baseSchema.Tables {
			if _, ok := siblings[table.Slug]; ok {
				return easyerr.Error("duplicate tables")
			}

			cols := make(map[string]string, len(table.Columns))
			for _, col := range table.Columns {
				if _, ok := cols[col.Slug]; ok {
					return easyerr.Error("duplicate columns")
				}

				cols[col.Slug] = col.Ctype
			}
		}

	} else {
		cols := make([]*entities.Column, 0)
		err := d.dataTableColumns().Find(db.Cond{
			"group_id":  opts.Slug,
			"tenant_id": tenantId,
		}).All(&cols)
		if err != nil {
			return err
		}

		for _, col := range cols {

			scols, ok := siblings[col.TableID]
			if !ok {
				scols = make(map[string]string)
				siblings[col.TableID] = scols
			}

			scols[col.Slug] = col.Ctype
		}

	}

	for _, mstep := range opts.Steps {

		switch mstep.Type {
		case step.MigTypeAddTable:

			tschema := &xbprint.NewTable{}
			err := json.Unmarshal(mstep.Data, baseSchema)
			if err != nil {
				return err
			}

			tstmt, err := d.dyngen.NewTable(tenantId, opts.Slug, tschema, []string{})
			if err != nil {
				return err
			}

			baseSchema.Tables = append(baseSchema.Tables, tschema)
			buf.WriteString(tstmt.String())

			if opts.New {
				baseSchema.Tables = append(baseSchema.Tables, tschema)
			} else {
				postitems = append(postitems, postDDLItem{
					mtype: mstep.Type,
					data:  tschema,
				})
			}

		case step.MigTypeRemoveTable:
			tschema := xbprint.RemoveTable{}
			err := json.Unmarshal(mstep.Data, baseSchema)
			if err != nil {
				return err
			}

			stmt, err := d.dyngen.DropTable(tenantId, opts.Slug, tschema.Slug)
			if err != nil {
				return err
			}

			buf.WriteString(stmt)

			if opts.New {

				newtables := make([]*xbprint.NewTable, 0, len(baseSchema.Tables))

				found := false
				for _, tbl := range baseSchema.Tables {
					if tbl.Slug == tschema.Slug {
						found = true
						continue
					}
					newtables = append(newtables, tbl)
				}
				if !found {
					return easyerr.Error("table to remove not found")
				}
				baseSchema.Tables = newtables
			} else {
				postitems = append(postitems, postDDLItem{
					mtype: mstep.Type,
					data:  tschema,
				})
			}

		case step.MigTypeAddColumn:
			tschema := xbprint.NewColumn{}
			err := json.Unmarshal(mstep.Data, baseSchema)
			if err != nil {
				return err
			}

			sout, err := d.dyngen.AddColumn(tenantId, opts.Slug, tschema.Table, tschema.Slug, &tschema)
			if err != nil {
				return err
			}

			buf.WriteString(sout)

			if opts.New {
				found := false

				for _, table := range baseSchema.Tables {
					if table.Slug == tschema.Table {
						found = true
						table.Columns = append(table.Columns, &tschema)
					}
				}

				if !found {
					return easyerr.Error("table not found to add column")
				}
			} else {

				postitems = append(postitems, postDDLItem{
					mtype: mstep.Type,
					data:  tschema,
				})

			}

		case step.MigTypeRemoveColumn:

			tschema := xbprint.RemoveColumn{}
			err := json.Unmarshal(mstep.Data, baseSchema)
			if err != nil {
				return err
			}

			sout, err := d.dyngen.DropColumn(tenantId, opts.Slug, tschema.Table, tschema.Slug)
			if err != nil {
				return err
			}

			buf.WriteString(sout)

			if opts.New {
				found := false

				for _, table := range baseSchema.Tables {
					if table.Slug == tschema.Table {
						found = true
						newcols := make([]*xbprint.NewColumn, 0, len(table.Columns))

						for _, nc := range table.Columns {
							if nc.Slug == tschema.Slug {
								continue
							}

							newcols = append(newcols, nc)
						}

						table.Columns = newcols
					}
				}

				if !found {
					return easyerr.Error("table not found to remove column")
				}

			} else {

				postitems = append(postitems, postDDLItem{
					mtype: mstep.Type,
					data:  tschema,
				})

			}

		default:
			panic("not implemented")
		}
	}

	mctx := migrateContext{
		baseSchema: baseSchema,
		stmtString: buf.String(),
		postItems:  postitems,
		siblings:   siblings,
	}

	if opts.New {
		return d.performNewMigrate(tenantId, mctx)
	}

	return d.performUpdateMigrate(tenantId, mctx)

}
