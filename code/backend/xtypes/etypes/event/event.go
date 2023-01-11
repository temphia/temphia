package event

type Request struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Data []byte `json:"data,omitempty"`
}

type Response struct {
	Payload []byte `json:"payload,omitempty"`
}
