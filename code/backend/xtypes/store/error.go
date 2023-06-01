package store

type Error struct {
	Inner string
	LogId string
}

func (e *Error) Error() string {
	return e.Inner
}
