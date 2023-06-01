package store

type Error struct {
	Inner   string
	LogId   string
	CtxData any
}

func (e *Error) Error() string {
	return e.Inner
}
