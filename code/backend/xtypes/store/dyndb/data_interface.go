package dyndb

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type DataHub interface {
	Inject(app xtypes.App)

	xplane.StateWatcher

	GetDataTableHub(tenantId, group string) DataTableHub
	GetDataSheetHub(tenantId, group string) DataSheetHub

	GetDynDB() DynDB
}

type DataTableHub interface {
	NewRow(txid uint32, req NewRowReq) (int64, error)
	NewBatchRows(txid uint32, req NewBatchRowReq) ([]int64, error)
	GetRow(txid uint32, req GetRowReq) (map[string]any, error)
	UpdateRow(txid uint32, req UpdateRowReq) (map[string]any, error)
	DeleteRowBatch(txid uint32, req DeleteRowBatchReq) ([]int64, error)
	DeleteRowMulti(txid uint32, req DeleteRowMultiReq) error
	DeleteRow(txid uint32, req DeleteRowReq) error

	LoadTable(txid uint32, req LoadTableReq) (*LoadTableResp, error)

	SimpleQuery(txid uint32, req SimpleQueryReq) (*QueryResult, error)
	JoinQuery(txid uint32, req JoinReq) (*JoinResult, error)
	MultiJoinQuery(txid uint32, req MultiJoinReq) (*MultiJoinResult, error)

	FTSQuery(txid uint32, req FTSQueryReq) (*QueryResult, error)
	RefResolve(txid uint32, gslug string, req *RefResolveReq) (*QueryResult, error)
	RefLoad(txid uint32, gslug string, req *RefLoadReq) (*QueryResult, error)
	ReverseRefLoad(txid uint32, gslug string, req *RevRefLoadReq) (*QueryResult, error)

	SqlQuery(txid uint32, req SqlQueryReq) (*SqlQueryResult, error)

	QueryActivity(table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error)
	ListActivity(table string, rowId int) ([]*entities.DynActivity, error)
	ListActivityByAlt(table string, alt string) ([]*entities.DynActivity, error)

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
	FTSQuery(txid uint32, data *FTSQuerySheet) (*QuerySheetResp, error)
	RefQuery(txid uint32, data *RefQuerySheet) (*QuerySheetResp, error)

	NewRowWithCell(txid uint32, sid int64, userId string, data map[int64]map[string]any) ([]map[string]any, error)
	UpdateRowWithCell(txid uint32, sid, rid int64, userId string, data map[int64]map[string]any) ([]map[string]any, error)
	DeleteRowWithCell(txid uint32, sid, rid int64, userId string) error

	GetRowHistory(rowid int64) ([]*entities.DynActivity, error)

	ExportSheets(txid uint32, opts ExportOptions) (*ExportData, error)
	ImportSheets(txid uint32, opts ImportOptions, data *ExportData) error

	GetRowRelations(txid uint32, sid, rid, refsheet, refcol int64) (*Relation, error)
}
