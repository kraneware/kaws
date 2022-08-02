package kaws_test

import (
	"github.com/kraneware/kaws"
	"testing"
)

func TestAwsConnectionValid(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kaws.AwsConnectionValid(); got != tt.want {
				t.Errorf("AwsConnectionValid() = %v, want %v", got, tt.want)
			}
		})
	}
}