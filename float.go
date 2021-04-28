package maybe

import "encoding/json"

type Float struct {
	value    float64
	hasValue bool
}

func SetFloat(value float64) Float {
	return Float{
		value:    value,
		hasValue: true,
	}
}

func (mf Float) HasValue() bool {
	return mf.hasValue
}

func (mf Float) Value() float64 {
	return mf.value
}

func (mf *Float) UnmarshalJSON(data []byte) error {
	var f *float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	if f != nil {
		*mf = SetFloat(*f)
	}

	return nil
}

func (mf Float) MarshalJSON() ([]byte, error) {
	var f *float64

	if mf.hasValue {
		f = &mf.value
	}

	return json.Marshal(f)
}
