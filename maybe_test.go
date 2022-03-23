package maybe

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func Test_Maybe_Set(t *testing.T) {
	tests := []struct {
		name string
		args any
		want Maybe[any]
	}{
		{
			name: "maybe bool",
			args: true,
			want: Maybe[any]{hasValue: true, value: true},
		},
		{
			name: "maybe float",
			args: 72.4,
			want: Maybe[any]{hasValue: true, value: 72.4},
		},
		{
			name: "maybe int",
			args: 28,
			want: Maybe[any]{hasValue: true, value: 28},
		},
		{
			name: "maybe string",
			args: "foo",
			want: Maybe[any]{hasValue: true, value: "foo"},
		},
		{
			name: "maybe time",
			args: time.Date(2020, 04, 28, 18, 34, 52, 0, time.UTC),
			want: Maybe[any]{hasValue: true, value: time.Date(2020, 04, 28, 18, 34, 52, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Set(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Maybe_HasValue(t *testing.T) {
	tests := []struct {
		name string
		arg  Maybe[any]
		want any
	}{
		{
			name: "has value",
			arg:  Maybe[any]{hasValue: true},
			want: true,
		},
		{
			name: "hasn't value",
			arg:  Maybe[any]{},
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

func Test_Maybe_Value(t *testing.T) {
	tests := []struct {
		name string
		arg  Maybe[int]
		want any
	}{
		{
			name: "value is set",
			arg:  Set(24),
			want: 24,
		},
		{
			name: "Value is not set",
			arg:  Maybe[int]{},
			want: 0,
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

func Test_Maybe_Marshal(t *testing.T) {
	type person struct {
		IsCitizen Maybe[bool] `json:"is_citizen"`
	}
	tests := []struct {
		name    string
		data    person
		want    []byte
		wantErr bool
	}{
		{
			name:    "property is serialized",
			data:    person{IsCitizen: Set(false)},
			want:    []byte(`{"is_citizen":false}`),
			wantErr: false,
		},
		{
			name:    "Property isn't get serialized",
			data:    person{},
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

func Test_Maybe_Unmarshal(t *testing.T) {
	type person struct {
		IsCitizen Maybe[bool] `json:"is_citizen"`
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
			want:    person{IsCitizen: Set(false)},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value",
			data:    []byte(`{"is_citizen":null}`),
			want:    person{IsCitizen: Maybe[bool]{hasValue: false}},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value (property missing)",
			data:    []byte(`{}`),
			want:    person{IsCitizen: Maybe[bool]{hasValue: false}},
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
