package maybe

import (
	"encoding/json"
	"time"
)

type Time struct {
	value    time.Time
	hasValue bool
}

func SetTime(value time.Time) Time {
	return Time{
		value:    value,
		hasValue: true,
	}
}

func (mt Time) HasValue() bool {
	return mt.hasValue
}

func (mt Time) Value() time.Time {
	return mt.value
}

func (mt *Time) UnmarshalJSON(data []byte) error {
	var t *time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	if t != nil {
		*mt = SetTime(*t)
	}

	return nil
}

func (mt Time) MarshalJSON() ([]byte, error) {
	var t *time.Time

	if mt.hasValue {
		t = &mt.value
	}

	return json.Marshal(t)
}
