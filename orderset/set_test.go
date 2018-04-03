package orderset

import (
	"reflect"
	"testing"
)

func TestGetSet(t *testing.T) {
	type args struct {
		array []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSet(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
