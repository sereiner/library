package util

import (
	"testing"
)

func TestHash(t *testing.T) {
	t.Log(Hash("hello12") % 4)
}

func TestGUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name:"1",want:""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GUID(); got != tt.want {
				t.Errorf("GUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
