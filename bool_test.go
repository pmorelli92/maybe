package maybe

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_SetBool(t *testing.T) {
	tests := []struct {
		name string
		args bool
		want Bool
	}{
		{
			name: "Valid bool",
			args: true,
			want: Bool{hasValue: true, value: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetBool(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Bool_HasValue(t *testing.T) {
	tests := []struct {
		name string
		arg  Bool
		want bool
	}{
		{
			name: "Has value",
			arg:  Bool{hasValue: true},
			want: true,
		},
		{
			name: "Hasn't value",
			arg:  Bool{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.HasValue(); got != tt.want {
				t.Errorf("HasValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Bool_Value(t *testing.T) {
	tests := []struct {
		name string
		arg  Bool
		want bool
	}{
		{
			name: "Value is set",
			arg:  SetBool(true),
			want: true,
		},
		{
			name: "Value is not set",
			arg:  Bool{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.Value(); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Bool_Marshal(t *testing.T) {
	type person struct {
		IsCitizen Bool `json:"is_citizen"`
	}
	tests := []struct {
		name    string
		data    person
		want    []byte
		wantErr bool
	}{
		{
			name:    "Property get serialised",
			data:    person{IsCitizen: SetBool(false)},
			want:    []byte(`{"is_citizen":false}`),
			wantErr: false,
		},
		{
			name:    "Property does not get serialised",
			data:    person{IsCitizen: Bool{}},
			want:    []byte(`{"is_citizen":null}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func Test_Bool_Unmarshal(t *testing.T) {
	type person struct {
		IsCitizen Bool `json:"is_citizen"`
	}
	tests := []struct {
		name    string
		data    []byte
		want    person
		wantErr bool
	}{
		{
			name:    "Unmarshal with value",
			data:    []byte(`{"is_citizen":false}`),
			want:    person{IsCitizen: SetBool(false)},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value",
			data:    []byte(`{"is_citizen":null}`),
			want:    person{IsCitizen: Bool{}},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value (property missing)",
			data:    []byte(`{}`),
			want:    person{IsCitizen: Bool{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got person
			err := json.Unmarshal(tt.data, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
