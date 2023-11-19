package client

type HttpData interface {
	IsJSON() bool
	AsJSON(target any) error
	AsRaw() []byte
}

type HttpBase interface {
	GET(path string) (HttpData, error)
	POST(path string, data any) (HttpData, error)
	PUT(path string, data any) (HttpData, error)
	PATCH(path string, data any) (HttpData, error)
	DELETE(path string) (HttpData, error)
}
