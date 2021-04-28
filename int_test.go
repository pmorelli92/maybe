package maybe

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_SetInt(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want Int
	}{
		{
			name: "Valid int",
			args: args{value: 28},
			want: Int{hasValue: true, value: 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetInt(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Int_Marshal(t *testing.T) {
	type person struct {
		Age Int `json:"age"`
	}
	tests := []struct {
		name    string
		data    person
		want    []byte
		wantErr bool
	}{
		{
			name:    "Property get serialised",
			data:    person{Age: SetInt(28)},
			want:    []byte(`{"age":28}`),
			wantErr: false,
		},
		{
			name:    "Property does not get serialised",
			data:    person{Age: Int{}},
			want:    []byte(`{"age":null}`),
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

func Test_Int_Unmarshal(t *testing.T) {
	type person struct {
		Age Int `json:"age"`
	}
	tests := []struct {
		name    string
		data    []byte
		want    person
		wantErr bool
	}{
		{
			name:    "Unmarshal with value",
			data:    []byte(`{"age":28}`),
			want:    person{Age: SetInt(28)},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value",
			data:    []byte(`{"age":null}`),
			want:    person{Age: Int{}},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value (property missing)",
			data:    []byte(`{}`),
			want:    person{Age: Int{}},
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
