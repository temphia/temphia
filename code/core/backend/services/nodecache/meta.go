package nodecache

import (
	"encoding/binary"
)

type meta struct {
	version int64
	expire  int64
}

func getMeta(in []byte) meta {
	return meta{
		version: int64(binary.LittleEndian.Uint64(in[:8])),
		expire:  int64(binary.LittleEndian.Uint64(in[8:])),
	}
}

func (m meta) toByte() []byte {
	out := make([]byte, 16)
	binary.LittleEndian.PutUint64(out[:8], uint64(m.version))
	binary.LittleEndian.PutUint64(out[8:], uint64(m.expire))
	return out
}
