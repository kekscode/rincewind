package timekeeper

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		label    string
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
		want *TimeSlice
	}{
		{
			"Test work slice",
			args{
				"Work",
				time.Duration(time.Minute * 25)},
			&TimeSlice{Label: "Work", Duration: time.Duration(time.Minute * 25)},
		},
		//{name: "Test work slice", args{label: "Work", duration: (time.Minute * 25)}, want: *TimeSlice},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.label, tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
