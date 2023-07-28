package xtypes

type Mesh interface {
	Start(app App) error
	Stop() error
	GetAddress() []string
}
