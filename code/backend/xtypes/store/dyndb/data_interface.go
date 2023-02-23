package dyndb

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

type DataHub interface {
	DefaultSource(tenant string) DynSource
	GetSource(source, tenant string) DynSource
	ListSources(tenant string) ([]string, error)
	Inject(app xtypes.App)

	GetDataTableHub(source, tenantId, group string) DataTableHub
	GetDataSheetHub(source, tenantId, group string) DataSheetHub
}

type DynSource interface {
	Name() string

	NewGroup(tenantId string, model *xbprint.NewTableGroup) error
	EditGroup(tenantId, gslug string, model *entities.TableGroupPartial) error
	ListGroup(tenantId string) ([]*entities.TableGroup, error)
	GetGroup(tenantId, gslug string) (*entities.TableGroup, error)
	DeleteGroup(tenantId, gslug string) error

	EditTable(tenantId, gslug, tslug string, model *entities.TablePartial) error
	GetTable(tenantId, gslug, tslug string) (*entities.Table, error)
	ListTables(tenantId, gslug string) ([]*entities.Table, error)
	DeleteTable(tenantId, gslug, tslug string) error

	EditColumn(tenantId, gslug, tslug, cslug string, model *entities.ColumnPartial) error
	GetColumn(tenantId, gslug, tslug, cslug string) (*entities.Column, error)
	ListColumns(tenantId, gslug, tslug string) ([]*entities.Column, error)
	ListReverseColumnRef(tenantId, gslug, tslug string) ([]*entities.Column, error)
	DeleteColumn(tenantId, gslug, tslug, cslug string) error

	NewView(tenantId string, model *entities.DataView) error
	GetView(tenantId, gslug, tslug string, id int64) (*entities.DataView, error)
	ModifyView(tenantId, gslug, tslug string, id int64, data map[string]any) error
	ListView(tenantId, gslug, tslug string) ([]*entities.DataView, error)
	DelView(tenantId, gslug, tslug string, id int64) error

	QueryActivity(tenantId, group, table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error)
	ListActivity(tenantId, group, table string, rowId int) ([]*entities.DynActivity, error)
	NewActivity(tenantId, group, table string, record *entities.DynActivity) error

	GetDataTableHub(tenantId, group string) DataTableHub
	GetDataSheetHub(tenantId, group string) DataSheetHub
}

type DataTableHub interface {
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

	LiveSeed(table, userId string, max int) error
}

type DataSheetHub interface {
	ListSheetGroup(txid uint32) (*ListSheetGroupResp, error)

	ListSheet(txid uint32) ([]map[string]any, error)
	NewSheet(txid uint32, userId string, data map[string]any) error
	GetSheet(txid uint32, id int64) (map[string]any, error)
	UpdateSheet(txid uint32, id int64, userId string, data map[string]any) error
	DeleteSheet(txid uint32, id int64, userId string) error
	ListSheetColumn(txid uint32, sid int64) ([]map[string]any, error)
	NewSheetColumn(txid uint32, sid int64, userId string, data map[string]any) (int64, error)
	GetSheetColumn(txid uint32, sid, cid int64) (map[string]any, error)
	UpdateSheetColumn(txid uint32, sid, cid int64, userId string, data map[string]any) error
	DeleteSheetColumn(txid uint32, sid, cid int64, userId string) error

	LoadSheet(txid uint32, data *LoadSheetReq) (*LoadSheetResp, error)
	Query(txid uint32, data *QuerySheetReq) (*QuerySheetResp, error)

	NewRowWithCell(txid uint32, sid int64, userId string, data map[int64]map[string]any) (map[int64]map[string]any, error)
	UpdateRowWithCell(txid uint32, sid, rid int64, userId string, data map[int64]map[string]any) (map[int64]map[string]any, error)
}
