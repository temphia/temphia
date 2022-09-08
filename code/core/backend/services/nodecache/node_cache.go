package nodecache

import (
	"errors"
	"time"

	"github.com/temphia/temphia/code/core/backend/libx/xutils/kosher"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	bolt "go.etcd.io/bbolt"
)

var (
	ErrVersionMismatch = errors.New("VERSION MISMATCH")
	ErrNotFound        = errors.New("ERROR NOT FOUND")
)

var _ service.NodeCache = (*NodeCache)(nil)

type NodeCache struct {
	db *bolt.DB
}

func New(path string) *NodeCache {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		panic(err)
	}

	return &NodeCache{
		db: db,
	}
}

func (nc *NodeCache) Put(tenantId, space, key string, value []byte, expire int64) error {
	return nc.db.Update(func(tx *bolt.Tx) error {
		buk, err := tx.CreateBucketIfNotExists([]byte(space))
		if err != nil {
			return err
		}

		err = buk.Put(spaceKey(tenantId, key), value)
		if err != nil {
			return err
		}

		meta := meta{
			version: 0,
			expire:  expire,
		}

		mkey := spaceKeyMeta(tenantId, key)
		out := buk.Get(mkey)
		if out != nil {
			meta = getMeta(out)
		}

		meta.version = meta.version + 1
		return buk.Put(mkey, meta.toByte())
	})

}

func (nc *NodeCache) PutCAS(tenantId, space, key string, value []byte, version, expire int64) error {

	return nc.db.Update(func(tx *bolt.Tx) error {
		buk, err := tx.CreateBucketIfNotExists([]byte(space))
		if err != nil {
			return err
		}

		meta := meta{
			version: version + 1,
			expire:  getCurrTimr() + expire,
		}

		mkey := spaceKeyMeta(tenantId, key)
		out := buk.Get(mkey)
		if out == nil && version != 0 {
			return ErrVersionMismatch
		} else if out != nil {
			m2 := getMeta(out)
			if m2.version != version {
				return ErrVersionMismatch
			}
		}

		err = buk.Put(mkey, meta.toByte())
		if err != nil {
			return err
		}

		return buk.Put(spaceKey(tenantId, key), value)
	})

}

func (nc *NodeCache) Get(tenantId, space, key string) (data []byte, version int64, expire int64, err error) {

	var needsDel bool

	err = nc.db.View(func(tx *bolt.Tx) error {

		mkey := spaceKeyMeta(tenantId, key)
		kkey := spaceKey(tenantId, key)

		buk := tx.Bucket([]byte(space))
		data = buk.Get(kkey)
		if data == nil {
			return ErrNotFound
		}

		meta := getMeta(buk.Get(mkey))
		ctime := getCurrTimr()
		if ctime > meta.expire {
			needsDel = true
			return ErrNotFound
		}

		version = meta.version
		expire = meta.expire

		return nil
	})
	if err != nil {
		return
	}

	if needsDel {
		nc.Expire(tenantId, space, key)
	}

	return
}

func (nc *NodeCache) Expire(tenantId, space, key string) error {

	return nc.db.Update(func(tx *bolt.Tx) error {
		mkey := spaceKeyMeta(tenantId, key)
		kkey := spaceKey(tenantId, key)

		buk := tx.Bucket([]byte(space))
		err1 := buk.Delete(kkey)
		if err1 != nil {
			return err1
		}

		err1 = (buk.Delete(mkey))
		if err1 != nil {
			return err1
		}
		return nil
	})

}

// private

var metaPrefix = []byte(`__meta_x__`)

func spaceKey(tenantId, key string) []byte {
	out := make([]byte, len(tenantId)+len(key))
	out = append(out, kosher.Byte(tenantId)...)
	out = append(out, kosher.Byte(key)...)
	return out
}

func spaceKeyMeta(tenantId, key string) []byte {
	out := make([]byte, len(tenantId)+len(key)+len(metaPrefix))

	out = append(out, metaPrefix...)
	out = append(out, kosher.Byte(tenantId)...)
	out = append(out, kosher.Byte(key)...)
	return out
}

func getCurrTimr() int64 {
	return time.Now().UTC().UnixMilli()
}
