package helpers

import (
	"reflect"
	"testing"
)

func TestStringPtr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{name: "str to ptr", args: args{s: "test"}, want: StringPtr("test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringPtr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
