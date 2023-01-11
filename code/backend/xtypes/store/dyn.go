package store

import (
	"strings"

	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"

	"github.com/thoas/go-funk"
)

const (

	// column types ctype
	CtypeShortText = "shorttext"
	CtypePhone     = "phonenumber"
	CtypeSelect    = "select"
	CtypeRFormula  = "rowformula"
	CtypeFile      = "file"
	CtypeMultiFile = "multifile"
	CtypeCheckBox  = "checkbox"
	CtypeCurrency  = "currency"
	CtypeNumber    = "number"
	CtypeLocation  = "location"
	CtypeDateTime  = "datetime"

	CtypeMultSelect  = "multiselect"
	CtypeLongText    = "longtext"
	CtypeSingleUser  = "singleuser"
	CtypeMultiUser   = "multiuser"
	CtypeEmail       = "email"
	CtypeJSON        = "json"
	CtypeRangeNumber = "rangenumber"
	CtypeColor       = "color"
)

const (
	RefHardPriId = "hard_pri"
	RefSoftPriId = "soft_pri"
	RefHardText  = "hard_text"
	RefSoftText  = "soft_text"
	RefHardMulti = "hard_multi"
)

const (
	// index types
	IndexUnique = "unique"
	IndexNormal = "normal"
	IndexFTS    = "fts"
)

const (
	// meta keys
	KeyPrimary = "__id"
	KeyVersion = "__version"
	KeyModSig  = "__mod_sig"

	// meta reference keys
	KeyForceVersion     = "__force_version__"
	KeySecondary        = "__secondary_keys__"
	KeyErrorAfterUpdate = "__error_after_update__"
)

const (
	DynActivityTypeNone   = "none"
	DynActivityTypeStrict = "strict"
	DynActivityTypeLazy   = "lazy"
)

const (
	DynSyncTypeNone         = "none"
	DynSyncTypeEventOnly    = "event_only"
	DynSyncTypeEventAndData = "event_and_data"
)

const (
	DynSeedTypeData    = "data"
	DynSeedTypeAutogen = "autogen"
)

var (
	MetaRefKeys = []string{KeyForceVersion, KeySecondary, KeyErrorAfterUpdate}
	MetaKeys    = []string{KeyPrimary, KeyVersion, KeyModSig}
)

func IsMeta(name string) bool {
	if !strings.HasPrefix(name, "__") {
		return false
	}

	return funk.ContainsString(MetaKeys, name)
}

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

	AddTable(gslug string, model *xbprint.NewTable) error
	EditTable(gslug, tslug string, model *entities.TablePartial) error
	GetTable(gslug, tslug string) (*entities.Table, error)
	ListTables(gslug string) ([]*entities.Table, error)
	DeleteTable(gslug, tslug string) error

	AddColumn(gslug, tslug string, model *xbprint.NewColumn) error
	EditColumn(gslug, tslug, cslug string, model *entities.ColumnPartial) error
	GetColumn(gslug, tslug, cslug string) (*entities.Column, error)
	ListColumns(gslug, tslug string) ([]*entities.Column, error)
	DeleteColumn(gslug, tslug, cslug string) error

	AddIndex(gslug, tslug string, model *entities.Index) error
	AddUniqueIndex(gslug, tslug string, model *entities.Index) error
	AddFTSIndex(gslug, tslug string, model *entities.FTSIndex) error
	AddColumnFRef(gslug, tslug string, model *entities.ColumnFKRef) error
	ListIndex(gslug, tslug string) ([]*entities.Index, error)
	ListColumnRef(gslug, tslug string) ([]*entities.ColumnFKRef, error)
	RemoveIndex(gslug, tslug, slug string) error

	NewView(model *entities.DataView) error
	GetView(gslug, tslug string, id int64) (*entities.DataView, error)
	ModifyView(gslug, tslug string, id int64, data map[string]any) error
	ListView(gslug, tslug string) ([]*entities.DataView, error)
	DelView(gslug, tslug string, id int64) error

	QueryActivity(group, table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error)
	ListActivity(group, table string, rowId int) ([]*entities.DynActivity, error)
	NewActivity(group, table string, record *entities.DynActivity) error

	NewRow(txid uint32, req NewRowReq) (int64, error)
	GetRow(txid uint32, req GetRowReq) (map[string]any, error)
	UpdateRow(txid uint32, req UpdateRowReq) (map[string]any, error)
	DeleteRows(txid uint32, req DeleteRowReq) error
	SimpleQuery(txid uint32, req SimpleQueryReq) (*QueryResult, error)
	FTSQuery(txid uint32, req FTSQueryReq) (*QueryResult, error)
	RefResolve(txid uint32, gslug string, req *RefResolveReq) (*QueryResult, error)
	RefLoad(txid uint32, gslug string, req *RefLoadReq) (*QueryResult, error)
	ReverseRefLoad(txid uint32, gslug string, req *RevRefLoadReq) (*QueryResult, error)

	SqlQuery(txid uint32, req SqlQueryReq) (*SqlQueryResult, error)

	LiveSeed(group, table, userId string, max int) error
}

type DynDB interface {
	NewGroup(tenantId string, model *xbprint.NewTableGroup) error
	EditGroup(tenantId string, gslug string, model *entities.TableGroupPartial) error
	ListGroup(tenantId string) ([]*entities.TableGroup, error)
	GetGroup(tenantId, gslug string) (*entities.TableGroup, error)
	DeleteGroup(tenantId, gslug string) error

	AddTable(tenantId, gslug string, model *xbprint.NewTable) error
	GetTable(tenantId, gslug, tslug string) (*entities.Table, error)
	EditTable(tenantId, gslug, tslug string, model *entities.TablePartial) error
	ListTables(tenantId, gslug string) ([]*entities.Table, error)
	DeleteTable(tenantId, gslug, tslug string) error

	AddColumn(tenantId, gslug, tslug string, model *xbprint.NewColumn) error
	GetColumn(tenantId, gslug, tslug, cslug string) (*entities.Column, error)
	EditColumn(tenantId, gslug, tslug, cslug string, model *entities.ColumnPartial) error
	ListColumns(tenantId, group_slug, tslug string) ([]*entities.Column, error)
	DeleteColumn(tenantId, gslug, tslug, cslug string) error

	AddIndex(tenantId, gslug, tslug string, model *entities.Index) error
	AddUniqueIndex(tenantId, gslug, tslug string, model *entities.Index) error
	AddFTSIndex(tenantId, gslug, tslug string, model *entities.FTSIndex) error
	AddColumnFRef(tenantId, gslug, tslug string, model *entities.ColumnFKRef) error
	ListIndex(tenantId, gslug, tslug string) ([]*entities.Index, error)
	ListColumnRef(tenantId, gslug, tslug string) ([]*entities.ColumnFKRef, error)
	RemoveIndex(tenantId, gslug, tslug, slug string) error

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
	DeleteRows(txid uint32, req DeleteRowReq) error

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

type DCache interface {
	CachedColumns(tenantId, group, table string) (map[string]*entities.Column, error)
	EvictColumns(tenantId, group, table string)
}

func ExtractColumns(m *xbprint.NewTable, tenantId, gslug string) []*entities.Column {
	indexedCol := make(map[string]*entities.Column)

	for _, nc := range m.Columns {
		newcol := nc.To(tenantId, gslug, m.Slug)
		indexedCol[nc.Slug] = newcol

	}

	for _, colref := range m.ColumnRef {
		if colref.Slug == "" {
			colref.Slug = xid.New().String()
		}

		if colref.Type == RefHardPriId || colref.Type == RefSoftPriId {
			if len(colref.ToCols) == 0 {
				colref.ToCols = []string{KeyPrimary}
			}
		}

		for idx, colId := range colref.FromCols {
			col := indexedCol[colId]
			col.RefId = colref.Slug
			col.RefType = colref.Type
			col.RefTarget = colref.Target
			col.RefObject = colref.ToCols[idx]
			col.RefCopy = colref.RefCopy
		}
	}

	cols := make([]*entities.Column, 0, len(m.Columns)+len(m.ColumnRef))
	for _, v := range indexedCol {
		cols = append(cols, v)
	}

	return cols

}

type (
	Schema struct {
		Group   *entities.TableGroup
		Tables  map[string]*entities.Table
		Columns map[string]*entities.Column
	}
)

func (s *Schema) AddColumn(cols ...*entities.Column) {
	for _, col := range cols {
		s.Columns[col.TableID+col.Slug] = col
	}
}

func (s *Schema) GetColumn(table, column string) *entities.Column {
	return s.Columns[table+column]
}
