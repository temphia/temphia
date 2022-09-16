package wazero

import "encoding/json"

func (e *Executor) getBytes(offset, count int32) []byte {
	out, ok := e.mem.Read(e.context, uint32(offset), uint32(count))
	if !ok {
		panic(ErrOutofIndex)
	}

	return (out)
}

func (e *Executor) getString(offset, count int32) string {
	out, ok := e.mem.Read(e.context, uint32(offset), uint32(count))
	if !ok {
		panic(ErrOutofIndex)
	}

	return string(out)
}

func (e *Executor) getJSON(optPtr, optLen int32, target any) error {
	out, ok := e.mem.Read(e.context, uint32(optPtr), uint32(optLen))
	if !ok {
		panic(ErrOutofIndex)
	}

	return json.Unmarshal(out, target)
}

func (e *Executor) getStrMap(optPtr, optLen int32) (map[string]string, error) {
	m := make(map[string]string)

	out, ok := e.mem.Read(e.context, uint32(optPtr), uint32(optLen))
	if !ok {
		panic(ErrOutofIndex)
	}

	err := json.Unmarshal(out, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
