package dyndb

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

type DynDB interface {
	MigrateSchema(tenantId string, opts xpackage.MigrateOptions) error

	NewGroup(tenantId string, model *xpackage.NewTableGroup) error
	EditGroup(tenantId string, gslug string, model *entities.TableGroupPartial) error
	ListGroup(tenantId string, cond map[string]any) ([]*entities.TableGroup, error)
	GetGroup(tenantId, gslug string) (*entities.TableGroup, error)
	DeleteGroup(tenantId, gslug string) error

	GetTable(tenantId, gslug, tslug string) (*entities.Table, error)
	EditTable(tenantId, gslug, tslug string, model *entities.TablePartial) error
	ListTables(tenantId, gslug string) ([]*entities.Table, error)
	DeleteTable(tenantId, gslug, tslug string) error

	GetColumn(tenantId, gslug, tslug, cslug string) (*entities.Column, error)
	EditColumn(tenantId, gslug, tslug, cslug string, model *entities.ColumnPartial) error
	ListColumns(tenantId, group_slug, tslug string) ([]*entities.Column, error)
	ListReverseColumnRef(tenantId, gslug, tslug string) ([]*entities.Column, error)
	DeleteColumn(tenantId, gslug, tslug, cslug string) error

	NewView(tenantId string, model *entities.DataView) error
	GetView(tenantId, gslug, tslug string, id int64) (*entities.DataView, error)
	ModifyView(tenantId, gslug, tslug string, id int64, data map[string]any) error
	ListView(tenantId, gslug, tslug string) ([]*entities.DataView, error)
	DelView(tenantId, gslug, tslug string, id int64) error

	QueryActivity(tenantId, group, table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error)
	ListActivity(tenantId, group, table string, rowId int) ([]*entities.DynActivity, error)
	ListActivityByAlt(tenantId, group, table string, alt string) ([]*entities.DynActivity, error)

	NewActivity(tenantId, group, table string, record *entities.DynActivity) (int64, error)

	// ops
	NewRow(txid uint32, req NewRowReq) (int64, error)
	NewBatchRows(txid uint32, req NewBatchRowReq) ([]int64, error)

	GetRow(txid uint32, req GetRowReq) (map[string]any, error)
	UpdateRow(txid uint32, req UpdateRowReq) (map[string]any, error)
	DeleteRowBatch(txid uint32, req DeleteRowBatchReq) ([]int64, error)
	DeleteRowMulti(txid uint32, req DeleteRowMultiReq) error
	DeleteRow(txid uint32, req DeleteRowReq) error

	SimpleQuery(txid uint32, req SimpleQueryReq) (*QueryResult, error)
	FTSQuery(txid uint32, req FTSQueryReq) (*QueryResult, error)
	JoinQuery(txid uint32, req JoinReq) (*JoinResult, error)
	MultiJoinQuery(txid uint32, req MultiJoinReq) (*MultiJoinResult, error)

	RefResolve(txid uint32, tenantId, gslug string, req *RefResolveReq) (*QueryResult, error)
	RefLoad(txid uint32, tenantId, gslug string, req *RefLoadReq) (*QueryResult, error)
	ReverseRefLoad(txid uint32, tenantId, gslug string, req *RevRefLoadReq) (*QueryResult, error)

	TemplateQuery(txid uint32, req TemplateQueryReq) (*QueryResult, error)

	SqlQueryRaw(txid uint32, tenantId, group, qstr string) (any, error)
	SqlQueryScopped(txid uint32, tenantId, group, qstr string) (any, error)

	GetDriver() any

	GetCache() DCache
}
