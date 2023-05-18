package plugkv

import (
	"time"

	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/upper/db/v4"
)

const tableName = "plug_states"

type PlugKV struct {
	db     db.Session
	txn    dbutils.TxManager
	vendor string
}

func New(db db.Session, txn dbutils.TxManager, vendor string) *PlugKV {

	return &PlugKV{
		db:     db,
		txn:    txn,
		vendor: vendor,
	}

}

func (p *PlugKV) SetBatch(txid uint32, tenantId, plugId string, opts *store.SetBatchOptions) error {
	for _, record := range opts.Records {
		record["tenant_id"] = tenantId
		record["plug_id"] = plugId

		if _, ok := record["tag1"]; !ok {
			record["tag1"] = ""
		}
		if _, ok := record["tag2"]; !ok {
			record["tag2"] = ""
		}
		if _, ok := record["tag3"]; !ok {
			record["tag3"] = ""
		}
	}

	if txid == 0 {
		_txid, err := p.txn.NewTxn(p.db, 10)
		if err != nil {
			return err
		}

		txid = _txid
	}

	return p.stateTx(txid, func(tbl db.Collection) error {

		if opts.ClearBefore {
			err := tbl.Find(db.Cond{
				"tenant_id": tenantId,
				"plug_id":   plugId,
			}).Delete()
			if err != nil {
				return err
			}
		}

		inserter := tbl.Session().
			SQL().
			InsertInto(tableName).
			Batch(len(opts.Records))

		for _, data := range opts.Records {
			inserter.Values(data)
		}

		inserter.Done()
		return inserter.Err()
	})

}

func (p *PlugKV) Set(txid uint32, tenantId, plugId, key, value string, opts *store.SetOptions) error {
	return p.stateTx(txid, func(tbl db.Collection) error {

		cond := db.Cond{
			"tenant_id": tenantId,
			"plug_id":   plugId,
			"key":       key,
			"value":     value,
		}

		if opts != nil {
			if opts.Tag1 != "" {
				cond["tag1"] = opts.Tag1
			}
			if opts.Tag2 != "" {
				cond["tag2"] = opts.Tag2
			}

			if opts.Tag3 != "" {
				cond["tag3"] = opts.Tag3
			}

			if opts.TTL != 0 {
				cond["ttl"] = time.Now().Add(time.Second * time.Duration(opts.TTL))
			}
		}

		_, err := tbl.Insert(cond)

		return err
	})

}

func (p *PlugKV) Update(txid uint32, tenantId, plugId, key, value string, opts *store.UpdateOptions) error {
	return p.stateTx(txid, func(tbl db.Collection) error {
		sql := tbl.Session().SQL()
		sets := make([]interface{}, 0, 6)
		sets = append(sets, "value = ?", value)

		cond := db.Cond{
			"tenant_id": tenantId,
			"key":       key,
			"plug_id":   plugId,
		}

		if opts != nil {
			if opts.ForceVer && opts.WithVerison {
				return easyerr.Error("force version and with version in same statement")
			}

			if opts.WithVerison {
				cond["version"] = opts.Version
			}

			if opts.ForceVer {
				sets = append(sets, `version = ?`, opts.Version)
			} else {
				sets = append(sets, `version = (version +  ?)`, 1)
			}

			if opts.SetTag1 {
				cond["tag1"] = opts.Tag1
			}
			if opts.SetTag2 {
				cond["tag2"] = opts.Tag2
			}
			if opts.SetTag3 {
				cond["tag3"] = opts.Tag3
			}

			if opts.TTL != 0 {
				cond["ttl"] = time.Now().Add(time.Second * time.Duration(opts.TTL))
			}
		}

		_, err := sql.Update(tableName).
			Set(sets...).
			Where(cond).Exec()

		return err
	})

}

func (p *PlugKV) Get(txid uint32, tenantId, plugId, key string) (*entities.PlugKV, error) {
	pkv := &entities.PlugKV{}
	err := p.stateTx(txid, func(tbl db.Collection) error {
		return tbl.Find(db.Cond{
			"tenant_id": tenantId,
			"plug_id":   plugId,
			"key":       key,
		}).One(pkv)
	})

	if err != nil {
		return nil, err
	}

	return pkv, err
}

func (p *PlugKV) Del(txid uint32, tenantId, plugId, key string) error {

	return p.stateTx(txid, func(tbl db.Collection) error {
		return tbl.Find(db.Cond{
			"tenant_id": tenantId,
			"plug_id":   plugId,
			"key":       key,
		}).Delete()
	})
}

func (p *PlugKV) Query(txid uint32, tenantId, plugId string, query *store.PkvQuery) ([]*entities.PlugKV, error) {
	data := make([]*entities.PlugKV, 0)
	err := p.stateTx(txid, func(tbl db.Collection) error {
		sql := tbl.Session().SQL()
		conds := db.Cond{
			"tenant_id": tenantId,
			"plug_id":   plugId,
		}

		addTag(conds, query)

		if query.KeyPrefix != "" {
			conds["key LIKE"] = query.KeyPrefix + "%"
		}

		if query.KeyCursor != "" {
			conds["key >"] = query.KeyCursor
		}

		slect := sql.SelectFrom(tableName).Where(
			conds,
			db.Or(db.Cond{"ttl": nil}, db.Cond{"ttl >": time.Now()}),
		)

		if query.NoValue {
			slect = slect.Columns("key", "version", "tag1", "tag2", "tag3", "ttl", "plug_id", "tenant_id")
		}

		err := slect.Paginate(query.PageCount).Page(query.Page + 1).All(&data)
		return err
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *PlugKV) DelBatch(txid uint32, tenantId, plugId string, keys []string) error {
	return p.stateTx(txid, func(tbl db.Collection) error {
		return tbl.Find(db.Cond{
			"tenant_id": tenantId,
			"plug_id":   plugId,
			"key IN":    keys,
		}).Delete()
	})
}

// txn stuff
func (s *PlugKV) NewTxn() (uint32, error) {
	return s.txn.NewTxn(s.db, 10)
}

func (s *PlugKV) RollBack(txid uint32) error {
	return s.txn.RollbackTx(txid)
}

func (s *PlugKV) Commit(txid uint32) error {
	return s.txn.CommitTx(txid)
}

// private

func (s *PlugKV) stateTx(tx uint32, fn func(tbl db.Collection) error) error {
	return s.txn.TxOr(tx, s.db, func(sess db.Session) error {
		return fn(pkvTable(sess))
	})
}

func pkvTable(sess db.Session) db.Collection {
	return dbutils.Table(sess, tableName)
}

func addTag(conds db.Cond, query *store.PkvQuery) {
	switch len(query.Tag1s) {
	case 0:
		break
	case 1:
		conds["tag1"] = query.Tag1s[0]
	default:
		conds["tag1 IN"] = query.Tag1s
	}

	switch len(query.Tag1s) {
	case 0:
		break
	case 1:
		conds["tag1"] = query.Tag1s[0]
	default:
		conds["tag1 IN"] = query.Tag1s
	}

	switch len(query.Tag2s) {
	case 0:
		break
	case 1:
		conds["tag2"] = query.Tag2s[0]
	default:
		conds["tag2 IN"] = query.Tag2s
	}

	switch len(query.Tag3s) {
	case 0:
		break
	case 1:
		conds["tag3"] = query.Tag3s[0]
	default:
		conds["tag3 IN"] = query.Tag3s
	}

}
