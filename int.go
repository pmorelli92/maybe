package maybe

import "encoding/json"

type Int struct {
	value    int
	hasValue bool
}

func SetInt(value int) Int {
	return Int{
		value:    value,
		hasValue: true,
	}
}

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

func (mi Int) MarshalJSON() ([]byte, error) {
	var i *int

	if mi.hasValue {
		i = &mi.value
	}

	return json.Marshal(i)
}
