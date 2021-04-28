package maybe

import (
	"encoding/json"
)

// Bool is an option data type, which means it can have a value or not.
type Bool struct {
	value    bool
	hasValue bool
}

// SetBool returns a Bool option with a value.
func SetBool(value bool) Bool {
	return Bool{
		value:    value,
		hasValue: true,
	}
}

// HasValue allows to check if the Bool has a value.
func (mb Bool) HasValue() bool {
	return mb.hasValue
}

// Value allows to check the Bool value.
func (mb Bool) Value() bool {
	return mb.value
}

// UnmarshalJSON customises the deserialize behaviour for the Bool option
func (mb *Bool) UnmarshalJSON(data []byte) error {
	var b *bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}

	if b != nil {
		*mb = SetBool(*b)
	}

	return nil
}

// MarshalJSON customises the serialize behaviour for the Bool option
func (mb Bool) MarshalJSON() ([]byte, error) {
	var b *bool

	if mb.hasValue {
		b = &mb.value
	}

	return json.Marshal(b)
}
