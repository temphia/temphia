package dyndb

import (
	"encoding/json"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func (d *DynDB) migrateSchema(tenantId string, opts step.MigrateOptions) error {

	var baseSchema *xbprint.NewTableGroup
	var buf strings.Builder
	//	var rollbackBuf strings.Builder

	postitems := make([]postDDLItem, 0)

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

			addPostItem(mstep.Type, tschema)

			if opts.New {
				// fixme => think of other validation ? also check when not new load siblings

				for _, nt := range baseSchema.Tables {
					if nt.Slug == tschema.Slug {
						return easyerr.Error("dup table name")
					}
				}

				baseSchema.Tables = append(baseSchema.Tables, tschema)
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

			addPostItem(mstep.Type, tschema)

			// fixme => remove column, baseschema

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

			addPostItem(mstep.Type, tschema)

			// fixme add to baseschema

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

func (d *DynDB) postDDLCreate(tenantId string, opts step.MigrateOptions, items []postDDLItem) error {

	return nil
}
