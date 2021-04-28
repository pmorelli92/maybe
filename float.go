package maybe

import "encoding/json"

// Float is an option data type, which means it can have a value or not.
type Float struct {
	value    float64
	hasValue bool
}

// SetFloat returns a Float option with a value.
func SetFloat(value float64) Float {
	return Float{
		value:    value,
		hasValue: true,
	}
}

// HasValue allows to check if the Float has a value.
func (mf Float) HasValue() bool {
	return mf.hasValue
}

// Value allows to check the Float value.
func (mf Float) Value() float64 {
	return mf.value
}

// UnmarshalJSON customises the deserialize behaviour for the Float option
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

// MarshalJSON customises the serialize behaviour for the Float option
func (mf Float) MarshalJSON() ([]byte, error) {
	var f *float64

	if mf.hasValue {
		f = &mf.value
	}

	return json.Marshal(f)
}
