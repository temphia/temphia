package xserver

type LSock interface {
	Register(s LSubcriber) int64
	SendWS(eid int64, name string, data []byte)
}

type LSubcriber interface {
	Handle(name string, data []byte)
	HandleWS(data []byte)
}
