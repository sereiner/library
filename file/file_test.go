package file

import (
	"os"
	"reflect"
	"testing"
)

func TestCreateFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantF   *os.File
		wantErr bool
	}{
		// TODO: Add test cases.
		{name:"1",args:args{path:"test.log"},wantF:nil,wantErr:true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, err := CreateFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotF, tt.wantF) {
				t.Errorf("CreateFile() = %v, want %v", gotF, tt.wantF)
			}
		})
	}
}
