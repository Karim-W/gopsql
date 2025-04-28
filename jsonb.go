package gopsql

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONB map[string]any

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	byts, ok := value.([]byte)
	if !ok {
		return errors.New("expected a marshalled json object")
	}

	return json.Unmarshal(byts, j)
}

func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	byts, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}

	return byts, nil
}

func (j *JSONB) MarshalJSON() ([]byte, error) {
	if j == nil {
		return json.Marshal(map[string]any{})
	}
	return json.Marshal(map[string]any(*j))
}

func (j *JSONB) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("gopsql: UnmarshalJSON on nil pointer")
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	*j = m
	return nil
}
