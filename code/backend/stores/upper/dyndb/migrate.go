package dyndb

import (
	"encoding/json"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (d *DynDB) migrateSchema(tenantId string, opts dyndb.MigrateOptions) error {

	var baseSchema *xbprint.NewTableGroup
	var buf strings.Builder

	postitems := make([]postDDLItem, 0)

	if opts.New {
		firstStep := opts.Steps[0]

		if firstStep.Type != dyndb.MigTypeNewGroup {
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
	}

	addPostItem := func(mtype string, data any) {

		if !opts.New {
			postitems = append(postitems, postDDLItem{
				mtype: mtype,
				data:  data,
			})
		}

	}

	for _, mstep := range opts.Steps {

		switch mstep.Type {
		case dyndb.MigTypeAddTable:

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

			addPostItem(mstep.Type, tschema)

			// fixme => add table to base schema

		case dyndb.MigTypeRemoveTable:
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

			addPostItem(mstep.Type, tschema)

			// fixme => remove column, baseschema

		case dyndb.MigTypeAddColumn:
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

			addPostItem(mstep.Type, tschema)

			// fixme add to baseschema

		case dyndb.MigTypeRemoveColumn:

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

			addPostItem(mstep.Type, tschema)

			// fixme remove column from baseschema

		default:
			panic("not implemented")
		}

	}

	return d.postDDLCreate(tenantId, opts, postitems)
}

type postDDLItem struct {
	mtype string
	data  any
}

func (d *DynDB) postDDLCreate(tenantId string, opts dyndb.MigrateOptions, items []postDDLItem) error {

	return nil
}
