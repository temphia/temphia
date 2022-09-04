package entities

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/k0kubun/pp"
)

func JSONDriverValue(val any) (driver.Value, error) {
	if val == nil {
		return "", nil
	}

	out, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	return string(out), nil
}

func JSONDriverScan(target any, value any) error {
	if value == nil {
		return nil
	}

	pp.Println("target @=>", target, value)

	un := func(b []byte) error {
		return json.Unmarshal(b, target)
	}
	switch s := value.(type) {
	case string:

		pp.Println("##")
		if s == "" {
			return nil
		}

		return un([]byte(s))
	case []byte:
		pp.Println("##")
		return un(s)

	}
	return nil
}

type JsonMap map[string]any

func (j JsonMap) Value() (driver.Value, error) {
	return JSONDriverValue(j)
}
func (j JsonMap) Scan(value any) error {
	return JSONDriverScan(&j, value)
}

type JsonStrMap map[string]string

func (j JsonStrMap) Value() (driver.Value, error) {
	return JSONDriverValue(j)
}
func (j JsonStrMap) Scan(value any) error {
	return JSONDriverScan(&j, value)
}

type JsonArray []string

func (j *JsonArray) Value() (driver.Value, error) {
	if j == nil {
		return "", nil
	}

	out, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return string(out), nil
}

func (j *JsonArray) Scan(value any) error {
	if value == nil {
		return nil
	}

	switch s := value.(type) {
	case string:
		if s == "" {
			return nil
		}
		return json.Unmarshal([]byte(s), &j)
	case []byte:
		return json.Unmarshal(s, j)
	}
	return nil
}
