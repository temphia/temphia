package cache

import (
	"bytes"
	"io"

	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
)

type BprintLoader struct {
	pacaman  repox.Hub
	bprintid string
	tenantId string
}

func NewBprintLoader(tenantId, bprintid string, pcaman repox.Hub) *BprintLoader {

	return &BprintLoader{
		pacaman:  pcaman,
		bprintid: bprintid,
	}
}

func (b *BprintLoader) Has(key string) (bool, error) {
	panic("not implemented")
}

func (b *BprintLoader) Get(dst io.Writer, key string) error {
	out, err := b.pacaman.BprintGetBlob(b.tenantId, b.bprintid, key)
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, bytes.NewReader(out))
	return err
}

func (b *BprintLoader) Put(key string, src io.Reader) error {
	panic("not implemented")
}
