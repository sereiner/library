package util

import (
	"testing"
)

func TestHash(t *testing.T) {
	t.Log(Hash("hello"))
}

func BenchmarkHash(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Hash("hello")
	}
}


func TestGUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "1", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GUID(); got != tt.want {
				t.Errorf("GUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash32(t *testing.T) {
	t.Log(Hash32("hello"))
}

func BenchmarkHash32(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Hash32("hello")
	}
}
