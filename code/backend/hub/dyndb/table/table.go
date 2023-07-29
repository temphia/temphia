package table

import (
	"encoding/json"
	"fmt"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/backend/hub/dyndb/handle"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var _ dyndb.DataTableHub = (*Table)(nil)

type Table struct {
	inner    dyndb.DynDB
	handle   *handle.Handle
	source   string
	tenantId string
	group    string
}

func New(source, tenantId, group string, inner dyndb.DynDB, handle *handle.Handle) *Table {
	return &Table{
		inner:    inner,
		handle:   handle,
		source:   source,
		group:    group,
		tenantId: tenantId,
	}
}

func (t *Table) NewRow(txid uint32, req dyndb.NewRowReq) (int64, error) {

	id, err := t.inner.NewRow(txid, req)
	if err != nil {
		return 0, err
	}

	req.Data[dyndb.KeyPrimary] = id

	if txid != 0 || req.NoReact {
		// fixme => when txid != 0
		return id, nil
	}

	err = t.handle.SockdHub.PushNewRow(t.source, req.TenantId, req.Group, req.Table, []int64{id}, req.Data)
	if err != nil {
		pp.Println(err)
	}

	return id, nil
}

func (t *Table) NewBatchRows(txid uint32, req dyndb.NewBatchRowReq) ([]int64, error) {

	ids, err := t.inner.NewBatchRows(txid, req)
	if err != nil {
		return nil, err
	}

	if txid != 0 || req.NoReact {
		return ids, nil
	}

	err = t.handle.SockdHub.PushNewRow(t.source, req.TenantId, req.Group, req.Table, ids, req.Data)
	if err != nil {
		pp.Println(err)
	}

	return ids, nil
}

func (t *Table) GetRow(txid uint32, req dyndb.GetRowReq) (map[string]any, error) {
	if txid != 0 || req.SkipCache {
		return t.inner.GetRow(txid, req)
	}

	// fixme => use cache
	return t.inner.GetRow(txid, req)

}

func (t *Table) UpdateRow(txid uint32, req dyndb.UpdateRowReq) (map[string]any, error) {

	data, err := t.inner.UpdateRow(txid, req)
	if err != nil {
		return nil, err
	}

	if txid != 0 || req.NoReact {
		return data, nil
	}

	err = t.handle.SockdHub.PushUpdateRow(t.source, req.TenantId, req.Group, req.Table, []int64{req.Id}, req.Data)
	if err != nil {
		pp.Println(err)
	}

	return data, nil

}

func (t *Table) DeleteRowBatch(txid uint32, req dyndb.DeleteRowBatchReq) ([]int64, error) {
	ids, err := t.inner.DeleteRowBatch(txid, req)
	if err != nil {
		return nil, err
	}

	if txid != 0 || req.NoReact {
		return ids, nil
	}

	t.handle.SockdHub.PushDeleteRow(t.source, req.TenantId, req.Group, req.Table, ids)

	return ids, nil
}

func (t *Table) DeleteRowMulti(txid uint32, req dyndb.DeleteRowMultiReq) error {
	err := t.inner.DeleteRowMulti(txid, req)
	if err != nil {
		return err
	}

	if txid != 0 || req.NoReact {
		return nil
	}

	t.handle.SockdHub.PushDeleteRow(t.source, req.TenantId, req.Group, req.Table, req.Ids)

	return nil
}

func (t *Table) DeleteRow(txid uint32, req dyndb.DeleteRowReq) error {
	err := t.inner.DeleteRow(txid, req)
	if err != nil {
		return err
	}

	if txid != 0 || req.NoReact {
		return nil
	}

	t.handle.SockdHub.PushDeleteRow(t.source, req.TenantId, req.Group, req.Table, []int64{req.Id})

	return nil
}

func (t *Table) LoadTable(txid uint32, req dyndb.LoadTableReq) (*dyndb.LoadTableResp, error) {

	views, err := t.inner.ListView(req.TenantId, req.Group, req.Table)
	if err != nil {
		return nil, err
	}

	sqr := dyndb.SimpleQueryReq{
		TenantId:    req.TenantId,
		Table:       req.Table,
		Group:       req.Group,
		Count:       50,
		FilterConds: make([]dyndb.FilterCond, 0),
		Page:        0,
		Selects:     nil,
	}

	finalResp := &dyndb.LoadTableResp{
		ReverseRefs:   nil,
		Views:         views,
		DataWidgets:   nil,
		ActiveView:    "",
		FolderTickets: make(map[string]string),
		UserTickets:   make(map[string]string),
	}

	if req.ViewFilters != nil {
		sqr.FilterConds = req.ViewFilters
	} else if req.View != "" {
		for _, view := range views {
			if view.Name == req.View {
				sqr.Count = view.Count
				fconds, err := json.Marshal(view.FilterConds)
				if err == nil {
					json.Unmarshal(fconds, &sqr.FilterConds)
				}
				sqr.Selects = view.Selects
				finalResp.ActiveView = view.Name
				break
			}
		}
	}

	sqresp, err := t.SimpleQuery(txid, sqr)
	if err != nil {
		return nil, err
	}
	finalResp.QueryResponse = sqresp

	apps, err := t.handle.CoreHub.ListTargetAppByType(req.TenantId, entities.TargetAppTypeDataTableWidget, fmt.Sprintf("%s/%s/%s", t.source, req.Group, req.Table))
	if err == nil {
		finalResp.DataWidgets = apps
	}

	refCols, err := t.inner.ListReverseColumnRef(req.Table, req.Group, req.Table)
	if err == nil {
		finalResp.ReverseRefs = refCols
	}

	pp.Println("@final_resp", finalResp)

	return finalResp, nil

}

func (t *Table) SimpleQuery(txid uint32, req dyndb.SimpleQueryReq) (*dyndb.QueryResult, error) {
	if req.Count == 0 {
		req.Count = dyndb.DefaultQueryFetchCount
	}

	// fixme => try fetch count + 1 and check if we got exact else
	// to determine if its last rows and return all resp_rows -1
	// you have to check if order by is other than __id

	resp, err := t.inner.SimpleQuery(txid, req)
	if err != nil {
		return nil, err
	}

	if req.Count >= int64(len(resp.Rows)) {
		resp.Final = true
	}

	return resp, nil
}

func (t *Table) JoinQuery(txid uint32, req dyndb.JoinReq) (*dyndb.JoinResult, error) {
	return t.inner.JoinQuery(txid, req)
}

func (t *Table) MultiJoinQuery(txid uint32, req dyndb.MultiJoinReq) (*dyndb.MultiJoinResult, error) {
	return t.inner.MultiJoinQuery(txid, req)
}

func (t *Table) FTSQuery(txid uint32, req dyndb.FTSQueryReq) (*dyndb.QueryResult, error) {
	return t.inner.FTSQuery(txid, req)
}

func (t *Table) RefResolve(txid uint32, gslug string, req *dyndb.RefResolveReq) (*dyndb.QueryResult, error) {
	return t.inner.RefResolve(txid, t.tenantId, gslug, req)
}

func (t *Table) RefLoad(txid uint32, gslug string, req *dyndb.RefLoadReq) (*dyndb.QueryResult, error) {
	return t.inner.RefLoad(txid, t.tenantId, gslug, req)
}

func (t *Table) ReverseRefLoad(txid uint32, gslug string, req *dyndb.RevRefLoadReq) (*dyndb.QueryResult, error) {
	return t.inner.ReverseRefLoad(txid, t.tenantId, gslug, req)
}

func (t *Table) SqlQuery(txid uint32, req dyndb.SqlQueryReq) (*dyndb.SqlQueryResult, error) {

	var records any
	var err error

	if !req.Raw {
		records, err = t.inner.SqlQueryScopped(txid, t.tenantId, req.Group, req.QStr)
	} else {
		// fixme => check if tenant allows raw query
		records, err = t.inner.SqlQueryRaw(txid, t.tenantId, req.Group, req.QStr)
	}

	if err != nil {
		return nil, err
	}

	return &dyndb.SqlQueryResult{
		Records: records,
		Columns: nil,
	}, nil

}

func (t *Table) LiveSeed(table, userId string, max int) error {
	// lseder, err := seeder2.NewLiveSeeder(seeder2.LiveSeederOptions{
	// 	TenantId:  tenantId,
	// 	Group:     group,
	// 	Table:     table,
	// 	UserId:    userId,
	// 	Source:    nil,
	// 	MaxRecord: int(max),
	// })

	// if err != nil {
	// 	return err
	// }

	// return lseder.Seed()

	return nil
}

func (t *Table) QueryActivity(table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error) {
	return t.inner.QueryActivity(t.tenantId, t.group, table, query)
}

func (t *Table) ListActivity(table string, rowId int) ([]*entities.DynActivity, error) {
	return t.inner.ListActivity(t.tenantId, t.group, table, rowId)
}

func (t *Table) ListActivityByAlt(table string, alt string) ([]*entities.DynActivity, error) {
	return t.inner.ListActivityByAlt(t.tenantId, t.group, table, alt)
}
