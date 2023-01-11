package job

import (
	"encoding/json"
)

func (j *Job) Serilize() ([]byte, error) {
	return json.Marshal(j)
}

func DeSerilize(data []byte) (*Job, error) {
	j := &Job{}

	err := json.Unmarshal(data, j)
	if err != nil {
		return nil, err
	}
	return j, nil
}
