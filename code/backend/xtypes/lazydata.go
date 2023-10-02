package xtypes

// LazyData is lazyly marsheled data, it might contain references
// to executor inner types so it should not be stored more than its
// request, response duration.
type LazyData interface {
	AsJsonBytes() ([]byte, error)
	AsObject(target any) error

	IsJsonBytes() bool
	IsObject() bool
	Inner() any
}

type BeBytes []byte
