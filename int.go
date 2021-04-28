package maybe

import "encoding/json"

type Int struct {
	Value    int
	HasValue bool
}

func (ms *Int) UnmarshalJSON(data []byte) error {
	var s *int
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s != nil {
		*ms = Int{
			Value:    *s,
			HasValue: true,
		}
	}

	return nil
}

func (ms Int) MarshalJSON() ([]byte, error) {
	var s *int

	if ms.HasValue {
		s = &ms.Value
	}

	return json.Marshal(s)
}
