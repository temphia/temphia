package xtension

type Builder func(ctx *Context) (Xtension, error)

type Xtension interface {
	Start() error
	Close() error
}
