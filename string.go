package maybe

import "encoding/json"

// String is an option data type, which means it can have a value or not.
type String struct {
	value    string
	hasValue bool
}

// SetString returns a String option with a value.
func SetString(value string) String {
	return String{
		value:    value,
		hasValue: true,
	}
}

// HasValue allows to check if the String has a value.
func (ms String) HasValue() bool {
	return ms.hasValue
}

// Value allows to check the String value.
func (ms String) Value() string {
	return ms.value
}

// UnmarshalJSON customises the deserialize behaviour for the String option
func (ms *String) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s != nil {
		*ms = SetString(*s)
	}

	return nil
}

// MarshalJSON customises the serialize behaviour for the String option
func (ms String) MarshalJSON() ([]byte, error) {
	var s *string

	if ms.hasValue {
		s = &ms.value
	}

	return json.Marshal(s)
}
