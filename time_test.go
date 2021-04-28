package maybe

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func Test_SetTime(t *testing.T) {
	type args struct {
		value time.Time
	}
	tests := []struct {
		name string
		args args
		want Time
	}{
		{
			name: "Valid time",
			args: args{value: time.Date(2020, 04, 28, 18, 34, 52, 0, time.UTC)},
			want: Time{hasValue: true, value: time.Date(2020, 04, 28, 18, 34, 52, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetTime(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Time_Marshal(t *testing.T) {
	type person struct {
		CreatedAt Time `json:"created_at"`
	}
	tests := []struct {
		name    string
		data    person
		want    []byte
		wantErr bool
	}{
		{
			name:    "Property get serialised",
			data:    person{CreatedAt: SetTime(time.Date(2020, 04, 28, 18, 34, 52, 0, time.UTC))},
			want:    []byte(`{"created_at":"2020-04-28T18:34:52Z"}`),
			wantErr: false,
		},
		{
			name:    "Property does not get serialised",
			data:    person{CreatedAt: Time{}},
			want:    []byte(`{"created_at":null}`),
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

func Test_Time_Unmarshal(t *testing.T) {
	type person struct {
		CreatedAt Time `json:"created_at"`
	}
	tests := []struct {
		name    string
		data    []byte
		want    person
		wantErr bool
	}{
		{
			name:    "Unmarshal with value",
			data:    []byte(`{"created_at":"2020-04-28T18:34:52Z"}`),
			want:    person{CreatedAt: SetTime(time.Date(2020, 04, 28, 18, 34, 52, 0, time.UTC))},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value",
			data:    []byte(`{"created_at":null}`),
			want:    person{CreatedAt: Time{}},
			wantErr: false,
		},
		{
			name:    "Unmarshal without value (property missing)",
			data:    []byte(`{}`),
			want:    person{CreatedAt: Time{}},
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
