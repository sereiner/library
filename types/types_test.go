package types

import "testing"

func TestGetString(t *testing.T) {
	type args struct {
		v   interface{}
		def []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name:"1",args:args{v:123},want:"123"},
		{name:"2",args:args{v:123.22},want:"123.22"},
		{name:"3",args:args{v:[]byte{46,47,48,49}},want:"./01"},
		{name:"4",args:args{v:"hello"},want:"hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.v, tt.args.def...); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		v   interface{}
		def []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name:"1",args:args{v:123,def:[]int{123}},want:123},
		{name:"2",args:args{v:"0123"},want:123},
		{name:"3",args:args{v:"123.12"},want:123},
		{name:"4",args:args{v:".12"},want:0},
		{name:"5",args:args{v:"1.2e+1",def:[]int{23}},want:12},
		{name:"6",args:args{v:"1.2E+",def:[]int{23}},want:23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt(tt.args.v, tt.args.def...); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt32(t *testing.T) {
	type args struct {
		v   interface{}
		def []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{name:"1",args:args{v:123},want:123},
		{name:"2",args:args{v:"1024"},want:1024},
		{name:"3",args:args{v:78.97},want:78},
		{name:"4",args:args{v:"78.97"},want:78},
		{name:"5",args:args{v:"0.97"},want:0},
		{name:"6",args:args{v:0.97},want:0},
		{name:"7",args:args{v:[]byte{49,50,51}},want:123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt32(tt.args.v, tt.args.def...); got != tt.want {
				t.Errorf("GetInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt64(t *testing.T) {
	type args struct {
		v   interface{}
		def []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{name:"1",args:args{v:123},want:123},
		{name:"2",args:args{v:"1024"},want:1024},
		{name:"3",args:args{v:78.97},want:78},
		{name:"4",args:args{v:"78.97"},want:78},
		{name:"5",args:args{v:"0.97"},want:0},
		{name:"6",args:args{v:0.97},want:0},
		{name:"7",args:args{v:[]byte{49,50,51}},want:123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt64(tt.args.v, tt.args.def...); got != tt.want {
				t.Errorf("GetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFloat32(t *testing.T) {
	type args struct {
		v   interface{}
		def []float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{name:"1",args:args{v:[]byte{50,46,51}},want:2.3},
		{name:"2",args:args{v:123.12},want:123.12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFloat32(tt.args.v, tt.args.def...); got != tt.want {
				t.Errorf("GetFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFloat64(t *testing.T) {
	type args struct {
		v   interface{}
		def []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{name:"1",args:args{v:[]byte{50,46,51}},want:2.3},
		{name:"2",args:args{v:123.12},want:123.12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFloat64(tt.args.v, tt.args.def...); got != tt.want {
				t.Errorf("GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUint32(t *testing.T) {
	type args struct {
		v   interface{}
		def []uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUint32(tt.args.v, tt.args.def...); got != tt.want {
				t.Errorf("GetUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUint64(t *testing.T) {
	type args struct {
		v   interface{}
		def []uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
		{name:"1",args:args{v:1},want:1},
		{name:"2",args:args{v:"123"},want:123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUint64(tt.args.v, tt.args.def...); got != tt.want {
				t.Errorf("GetUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
