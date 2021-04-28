package maybe

import "encoding/json"

type String struct {
	Value    string
	HasValue bool
}

func (ms *String) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s != nil {
		*ms = String{
			Value:    *s,
			HasValue: true,
		}
	}

	return nil
}

func (ms String) MarshalJSON() ([]byte, error) {
	var s *string

	if ms.HasValue {
		s = &ms.Value
	}

	return json.Marshal(s)
}
