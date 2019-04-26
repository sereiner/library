package net

import "testing"

func TestGetLocalIP(t *testing.T) {
	type args struct {
		masks []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name:"1",args:args{masks:nil},want:"127.0.0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLocalIP(tt.args.masks...); got != tt.want {
				t.Errorf("GetLocalIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
