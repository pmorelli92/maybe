package maybe

import "encoding/json"

type String struct {
	value    string
	hasValue bool
}

func SetString(value string) String {
	return String{
		value:    value,
		hasValue: true,
	}
}

func (ms String) HasValue() bool {
	return ms.hasValue
}

func (ms String) Value() string {
	return ms.value
}

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

func (ms String) MarshalJSON() ([]byte, error) {
	var s *string

	if ms.hasValue {
		s = &ms.value
	}

	return json.Marshal(s)
}
