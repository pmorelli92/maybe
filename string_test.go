package maybe

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_SetString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "Valid string",
			args: args{value: "Pablo"},
			want: String{hasValue: true, value: "Pablo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Marshal(t *testing.T) {
	type person struct {
		Name String `json:"name"`
	}
	tests := []struct {
		name    string
		data    person
		want    []byte
		wantErr bool
	}{
		{
			name:    "Property get serialised",
			data:    person{Name: SetString("Pablo")},
			want:    []byte(`{"name":"Pablo"}`),
			wantErr: false,
		},
		{
			name:    "Property does not get serialised",
			data:    person{Name: String{}},
			want:    []byte(`{"name":null}`),
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

func Test_String_Unmarshal(t *testing.T) {
	type person struct {
		Name String `json:"name"`
	}
	tests := []struct {
		name    string
		data    []byte
		want    person
		wantErr bool
	}{
		{
			name:    "Unmarshal with value",
			data:    []byte(`{"name":"Pablo"}`),
			want:    person{Name: SetString("Pablo")},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value",
			data:    []byte(`{"name":null}`),
			want:    person{Name: String{}},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value (property missing)",
			data:    []byte(`{}`),
			want:    person{Name: String{}},
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
