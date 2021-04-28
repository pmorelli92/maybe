package maybe

import (
	"encoding/json"
	"time"
)

// Time is an option data type, which means it can have a value or not.
type Time struct {
	value    time.Time
	hasValue bool
}

// SetTime returns a Time option with a value.
func SetTime(value time.Time) Time {
	return Time{
		value:    value,
		hasValue: true,
	}
}

// HasValue allows to check if the Time has a value.
func (mt Time) HasValue() bool {
	return mt.hasValue
}

// Value allows to check the Time value.
func (mt Time) Value() time.Time {
	return mt.value
}

// UnmarshalJSON customises the deserialize behaviour for the Time option
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

// MarshalJSON customises the serialize behaviour for the Time option
func (mt Time) MarshalJSON() ([]byte, error) {
	var t *time.Time

	if mt.hasValue {
		t = &mt.value
	}

	return json.Marshal(t)
}
