package maybe

import "encoding/json"

// Int is an option data type, which means it can have a value or not.
type Int struct {
	value    int
	hasValue bool
}

// SetInt returns a Int option with a value.
func SetInt(value int) Int {
	return Int{
		value:    value,
		hasValue: true,
	}
}

// HasValue allows to check if the Int has a value.
func (mi Int) HasValue() bool {
	return mi.hasValue
}

// Value allows to check the Int value.
func (mi Int) Value() int {
	return mi.value
}

// UnmarshalJSON customises the deserialize behaviour for the Int option
func (mi *Int) UnmarshalJSON(data []byte) error {
	var i *int
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}

	if i != nil {
		*mi = SetInt(*i)
	}

	return nil
}

// MarshalJSON customises the serialize behaviour for the Int option
func (mi Int) MarshalJSON() ([]byte, error) {
	var i *int

	if mi.hasValue {
		i = &mi.value
	}

	return json.Marshal(i)
}
