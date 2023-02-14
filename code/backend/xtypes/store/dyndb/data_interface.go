package dyndb

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

type DataHub interface {
	DefaultSource(tenant string) DynSource
	GetSource(source, tenant string) DynSource
	ListSources(tenant string) ([]string, error)
	Inject(app xtypes.App)
}

type DynSource interface {
	Name() string

	NewGroup(model *xbprint.NewTableGroup) error
	EditGroup(gslug string, model *entities.TableGroupPartial) error
	ListGroup() ([]*entities.TableGroup, error)
	GetGroup(gslug string) (*entities.TableGroup, error)
	DeleteGroup(gslug string) error

	EditTable(gslug, tslug string, model *entities.TablePartial) error
	GetTable(gslug, tslug string) (*entities.Table, error)
	ListTables(gslug string) ([]*entities.Table, error)
	DeleteTable(gslug, tslug string) error

	EditColumn(gslug, tslug, cslug string, model *entities.ColumnPartial) error
	GetColumn(gslug, tslug, cslug string) (*entities.Column, error)
	ListColumns(gslug, tslug string) ([]*entities.Column, error)
	ListReverseColumnRef(gslug, tslug string) ([]*entities.Column, error)
	DeleteColumn(gslug, tslug, cslug string) error

	NewView(model *entities.DataView) error
	GetView(gslug, tslug string, id int64) (*entities.DataView, error)
	ModifyView(gslug, tslug string, id int64, data map[string]any) error
	ListView(gslug, tslug string) ([]*entities.DataView, error)
	DelView(gslug, tslug string, id int64) error

	QueryActivity(group, table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error)
	ListActivity(group, table string, rowId int) ([]*entities.DynActivity, error)
	NewActivity(group, table string, record *entities.DynActivity) error

	DataTableOps
}

type DataTableOps interface {
	NewRow(txid uint32, req NewRowReq) (int64, error)
	GetRow(txid uint32, req GetRowReq) (map[string]any, error)
	UpdateRow(txid uint32, req UpdateRowReq) (map[string]any, error)
	DeleteRowBatch(txid uint32, req DeleteRowBatchReq) error
	DeleteRowMulti(txid uint32, req DeleteRowMultiReq) error
	DeleteRow(txid uint32, req DeleteRowReq) error

	LoadTable(txid uint32, req LoadTableReq) (*LoadTableResp, error)

	SimpleQuery(txid uint32, req SimpleQueryReq) (*QueryResult, error)
	FTSQuery(txid uint32, req FTSQueryReq) (*QueryResult, error)
	RefResolve(txid uint32, gslug string, req *RefResolveReq) (*QueryResult, error)
	RefLoad(txid uint32, gslug string, req *RefLoadReq) (*QueryResult, error)
	ReverseRefLoad(txid uint32, gslug string, req *RevRefLoadReq) (*QueryResult, error)

	SqlQuery(txid uint32, req SqlQueryReq) (*SqlQueryResult, error)

	LiveSeed(group, table, userId string, max int) error
}

type DataSheetOps interface {
	ListSheetGroup(uclaim *claim.Data) (*ListSheetGroupResp, error)
	LoadSheet(uclaim *claim.Data, data *LoadSheetReq) (*LoadSheetResp, error)
	ListSheet(uclaim *claim.Data) ([]map[string]any, error)
	NewSheet(uclaim *claim.Data, data map[string]any) error
	GetSheet(uclaim *claim.Data, id int64) (map[string]any, error)
	UpdateSheet(uclaim *claim.Data, id int64, data map[string]any) error
	DeleteSheet(uclaim *claim.Data, id int64) error
	ListSheetColumn(uclaim *claim.Data, sid int64) ([]map[string]any, error)
	NewSheetColumn(uclaim *claim.Data, sid int64, data map[string]any) (int64, error)
	GetSheetColumn(uclaim *claim.Data, sid, cid int64) (map[string]any, error)
	UpdateSheetColumn(uclaim *claim.Data, sid, cid int64, data map[string]any) error
	DeleteSheetColumn(uclaim *claim.Data, sid, cid int64) error
	NewRowWithCell(uclaim *claim.Data, sid int64, data map[int64]map[string]any) (map[int64]map[string]any, error)
	UpdateRowWithCell(uclaim *claim.Data, sid, rid int64, data map[int64]map[string]any) (map[int64]map[string]any, error)
}

type DynDB interface {
	NewGroup(tenantId string, model *xbprint.NewTableGroup) error
	EditGroup(tenantId string, gslug string, model *entities.TableGroupPartial) error
	ListGroup(tenantId string) ([]*entities.TableGroup, error)
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

	NewView(model *entities.DataView) error
	GetView(tenantId, gslug, tslug string, id int64) (*entities.DataView, error)
	ModifyView(tenantId, gslug, tslug string, id int64, data map[string]any) error
	ListView(tenantId, gslug, tslug string) ([]*entities.DataView, error)
	DelView(tenantId, gslug, tslug string, id int64) error

	QueryActivity(tenantId, group, table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error)
	ListActivity(tenantId, group, table string, rowId int) ([]*entities.DynActivity, error)
	NewActivity(tenantId, group, table string, record *entities.DynActivity) (int64, error)

	// ops
	NewRow(txid uint32, req NewRowReq) (int64, error)
	NewBatchRows(txid uint32, req NewBatchRowReq) ([]int64, error)

	GetRow(txid uint32, req GetRowReq) (map[string]any, error)
	UpdateRow(txid uint32, req UpdateRowReq) (map[string]any, error)
	DeleteRowBatch(txid uint32, req DeleteRowBatchReq) error
	DeleteRowMulti(txid uint32, req DeleteRowMultiReq) error
	DeleteRow(txid uint32, req DeleteRowReq) error

	SimpleQuery(txid uint32, req SimpleQueryReq) (*QueryResult, error)
	FTSQuery(txid uint32, req FTSQueryReq) (*QueryResult, error)

	RefResolve(txid uint32, tenantId, gslug string, req *RefResolveReq) (*QueryResult, error)
	RefLoad(txid uint32, tenantId, gslug string, req *RefLoadReq) (*QueryResult, error)
	ReverseRefLoad(txid uint32, tenantId, gslug string, req *RevRefLoadReq) (*QueryResult, error)

	TemplateQuery(txid uint32, req TemplateQueryReq) (*QueryResult, error)

	SqlQueryRaw(txid uint32, tenantId, group, qstr string) (*SqlQueryResult, error)
	SqlQueryScopped(txid uint32, tenantId, group, qstr string) (*SqlQueryResult, error)

	GetCache() DCache
}
