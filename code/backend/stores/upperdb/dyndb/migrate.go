package dyndb

import (
	"encoding/json"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/stores/upperdb/dyndb/dynddl2"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/upper/db/v4"
)

func (d *DynDB) migrateSchema(tenantId string, opts step.MigrateOptions) error {

	// debug all this

	var (
		baseSchema  *xbprint.NewTableGroup
		buf         = strings.Builder{}
		lastMigHead = ""
		nextMigHead = ""
	)

	postitems := make([]dynddl2.PostDDLItem, 0)

	siblings := make(map[string]map[string]string)

	if opts.New {
		firstStep := opts.Steps[0]
		nextMigHead = firstStep.Name

		if firstStep.Type != step.MigTypeNewGroup {
			return easyerr.Error("wrong type as first migration step")
		}

		baseSchema = &xbprint.NewTableGroup{}

		err := json.Unmarshal(firstStep.Data, baseSchema)
		if err != nil {
			return err
		}

		baseSchema.Slug = opts.Gslug

		stmt, err := d.dyngen.NewGroup(tenantId, baseSchema)
		if err != nil {
			return err
		}

		stmtstr := stmt.String()

		buf.WriteString(stmtstr)
		opts.Steps = opts.Steps[1:]

		postitems = append(postitems, dynddl2.PostDDLItem{
			Name:  firstStep.Name,
			Mtype: firstStep.Type,
			Data:  baseSchema,
			Stmt:  stmtstr,
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

		group, err := d.GetGroup(tenantId, opts.Gslug)
		if err != nil {
			pp.Println(tenantId, opts.Gslug)
			pp.Println("@not_found")
			return err
		}

		if group.BprintId != opts.BprintId {
			return easyerr.Error("wrong bprint_id")
		}

		if group.BprintInstanceId != opts.BprintInstanceId {
			return easyerr.Error("wrong bprint_instance_id")
		}

		if group.BprintItemId != opts.BprintItemId {
			return easyerr.Error("wrong bprint_item_id")
		}

		if group.BprintStepHead == "" {
			return easyerr.Error("bprint_step_head empty")
		}

		lastMigHead = group.BprintStepHead
		found := false
		for idx, step := range opts.Steps {
			if idx+1 == len(opts.Steps) {
				// "No more steps left"
				return nil
			}

			if step.Name == lastMigHead {
				found = true
				opts.Steps = opts.Steps[idx+1:]
				break
			}
		}

		if !found {
			return easyerr.Error("bprint_step_head not found")
		}

		pp.Println("@found/last_head", lastMigHead)
		pp.Println("@found/steps", opts.Steps)

		cols := make([]*entities.Column, 0)
		err = d.dataTableColumns().Find(db.Cond{
			"group_id":  opts.Gslug,
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
		pp.Println(mstep.Type)

		switch mstep.Type {
		case step.MigTypeAddTable:

			tschema := &xbprint.NewTable{}
			err := json.Unmarshal(mstep.Data, tschema)
			if err != nil {
				return err
			}

			tstmt, err := d.dyngen.NewTable(tenantId, opts.Gslug, tschema, []string{})
			if err != nil {
				return err
			}

			stmtstr := tstmt.String()

			buf.WriteString(stmtstr)

			if opts.New {
				baseSchema.Tables = append(baseSchema.Tables, tschema)
			} else {
				postitems = append(postitems, dynddl2.PostDDLItem{
					Name:  mstep.Name,
					Mtype: mstep.Type,
					Data:  tschema,
					Stmt:  stmtstr,
				})
			}

		case step.MigTypeRemoveTable:
			tschema := &xbprint.RemoveTable{}
			err := json.Unmarshal(mstep.Data, tschema)
			if err != nil {
				return err
			}

			stmt, err := d.dyngen.DropTable(tenantId, opts.Gslug, tschema.Slug)
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
				postitems = append(postitems, dynddl2.PostDDLItem{
					Name:  mstep.Name,
					Mtype: mstep.Type,
					Data:  tschema,
					Stmt:  stmt,
				})
			}

		case step.MigTypeAddColumn:
			tschema := &xbprint.NewColumn{}
			err := json.Unmarshal(mstep.Data, tschema)
			if err != nil {
				return err
			}

			stmtstr, err := d.dyngen.AddColumn(tenantId, opts.Gslug, tschema.Table, tschema.Slug, tschema)
			if err != nil {
				return err
			}

			buf.WriteString(stmtstr)

			if opts.New {
				found := false

				for _, table := range baseSchema.Tables {
					if table.Slug == tschema.Table {
						found = true
						table.Columns = append(table.Columns, tschema)
					}
				}

				if !found {
					return easyerr.Error("table not found to add column")
				}
			} else {

				postitems = append(postitems, dynddl2.PostDDLItem{
					Name:  mstep.Name,
					Mtype: mstep.Type,
					Data:  tschema,
					Stmt:  stmtstr,
				})

			}

		case step.MigTypeRemoveColumn:

			tschema := &xbprint.RemoveColumn{}
			err := json.Unmarshal(mstep.Data, tschema)
			if err != nil {
				return err
			}

			stmtstr, err := d.dyngen.DropColumn(tenantId, opts.Gslug, tschema.Table, tschema.Slug)
			if err != nil {
				return err
			}

			buf.WriteString(stmtstr)

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

				postitems = append(postitems, dynddl2.PostDDLItem{
					Name:  mstep.Name,
					Mtype: mstep.Type,
					Data:  tschema,
					Stmt:  stmtstr,
				})

			}

		default:
			panic("not implemented")
		}
	}

	if len(opts.Steps) > 0 {
		nextMigHead = opts.Steps[len(opts.Steps)-1].Name
	}

	mctx := dynddl2.MigrateContext{
		BaseSchema:  baseSchema,
		StmtString:  buf.String(),
		PostItems:   postitems,
		Siblings:    siblings,
		LastMigHead: lastMigHead,
		Options:     opts,
		NextMigHead: nextMigHead,
		Gslug:       opts.Gslug,
	}

	pp.Println("@mctx", mctx)

	if opts.DryRun {
		pp.Println("@dry_run_mctx", mctx)
		return nil
	}

	// fixme use txn
	runner := dynddl2.New(d.session, d.sharedLock, d.loggerBuilder())

	if opts.New {
		return runner.RunNew(tenantId, mctx)
	}

	return runner.RunUpdate(tenantId, mctx)
}
