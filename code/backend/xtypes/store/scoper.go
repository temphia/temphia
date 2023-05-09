package store

type Scoper interface {
	HasScope(scope string) bool
}
