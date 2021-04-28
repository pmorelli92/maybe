package maybe

import (
	"encoding/json"
)

// Bool allows
type Bool struct {
	value    bool
	hasValue bool
}

func (mb Bool) HasValue() bool {
	return mb.hasValue
}

func (mb Bool) Value() bool {
	return mb.value
}

func SetBool(value bool) Bool {
	return Bool{
		value:    value,
		hasValue: true,
	}
}

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

func (mb Bool) MarshalJSON() ([]byte, error) {
	var b *bool

	if mb.hasValue {
		b = &mb.value
	}

	return json.Marshal(b)
}
