package pagequery

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/executors/execlib/goja2db"
	"github.com/temphia/temphia/code/executors/execlib/helper"
	"gopkg.in/yaml.v2"

	gojaExec "github.com/temphia/temphia/code/backend/engine/executors/javascript1/goja"
)

type PgBuilder struct {
	app    xtypes.App
	helper *helper.ExecutorHelper
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	h := helper.New("executor_pagequery", true)

	return &PgBuilder{
		app:    app.(xtypes.App),
		helper: h,
	}, nil
}

func (pg *PgBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	ffile := opts.EnvVars["model_file"]
	if ffile == "" {
		ffile = "model1.yaml"
	}

	out, _, err := opts.Binder.GetFileWithMeta(ffile)
	if err != nil {
		return nil, err
	}

	pp.Println("@model", string(out))

	model := &PgModel{}
	err = yaml.Unmarshal(out, model)
	if err != nil {
		return nil, err
	}

	rt := goja.New()

	if opts.File != "" {
		out, _, err := opts.Binder.GetFileWithMeta(opts.File)
		if err != nil {
			pp.Println(fmt.Sprintf("server.js could not load: %s", err.Error()))
			return nil, err
		}

		_, err = gojaExec.New(opts.Binder, rt)
		if err != nil {
			return nil, err
		}

		v, err := rt.RunString(string(out))

		fmt.Printf("@running_server.js %v %v \n", v, err)
	}

	dhub := pg.app.GetDeps().DataHub().(dyndb.DataHub)

	gdb := goja2db.New(opts.TenantId, dhub, rt)
	gdb.Bind()

	return &PageQuery{
		builder:   pg,
		model:     model,
		jsruntime: rt,
		binder:    opts.Binder,
		tenantId:  opts.TenantId,
		datahub:   dhub,
	}, nil
}

func (pg *PgBuilder) IfaceFile() (*etypes.ExecutorIface, error) {
	return &etypes.ExecutorIface{
		Methods: map[string]*etypes.Method{},
		Bindings: map[string]*etypes.Method{
			// core
			"get_execdata":       nil,
			"get_execdata_item":  nil,
			"get_paramdata":      nil,
			"get_paramdata_item": nil,
			"get_stage":          nil,

			// table
			"table_query":      nil,
			"table_join_query": nil,
			"new_row":          nil,
			"get_row":          nil,
			"update_row":       nil,
			"delete_row_batch": nil,
			"delete_row_multi": nil,
			"delete_row":       nil,
			"load_table":       nil,
			"fts_query":        nil,
			"ref_resolve":      nil,
			"ref_load":         nil,
			"reverse_ref_load": nil,
			"sql_query":        nil,
			// sheet
			"list_sheet_group":     nil,
			"list_sheet":           nil,
			"new_sheet":            nil,
			"get_sheet":            nil,
			"update_sheet":         nil,
			"delete_sheet":         nil,
			"list_sheet_column":    nil,
			"new_sheet_column":     nil,
			"get_sheet_column":     nil,
			"update_sheet_column":  nil,
			"delete_sheet_column":  nil,
			"load_sheet":           nil,
			"query_sheet":          nil,
			"fts_query_sheet":      nil,
			"new_row_with_cell":    nil,
			"update_row_with_cell": nil,
			"delete_row_with_cell": nil,
			"get_row_relations":    nil,
		},
	}, nil
}

func (pg *PgBuilder) ExecFile(file string) ([]byte, error) {
	return pg.helper.Serve(file)
}
