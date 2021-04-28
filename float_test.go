package maybe

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_SetFloat(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Float
	}{
		{
			name: "Valid float",
			args: args{value: 72.4},
			want: Float{hasValue: true, value: 72.4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetFloat(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Float_Marshal(t *testing.T) {
	type person struct {
		Weight Float `json:"weight"`
	}
	tests := []struct {
		name    string
		data    person
		want    []byte
		wantErr bool
	}{
		{
			name:    "Property get serialised",
			data:    person{Weight: SetFloat(72.40)},
			want:    []byte(`{"weight":72.4}`),
			wantErr: false,
		},
		{
			name:    "Property does not get serialised",
			data:    person{Weight: Float{}},
			want:    []byte(`{"weight":null}`),
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

func Test_Float_Unmarshal(t *testing.T) {
	type person struct {
		Weight Float `json:"weight"`
	}
	tests := []struct {
		name    string
		data    []byte
		want    person
		wantErr bool
	}{
		{
			name:    "Unmarshal with value",
			data:    []byte(`{"weight":72.4}`),
			want:    person{Weight: SetFloat(72.40)},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value",
			data:    []byte(`{"weight":null}`),
			want:    person{Weight: Float{}},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value (property missing)",
			data:    []byte(`{}`),
			want:    person{Weight: Float{}},
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
