package utility

import (
	"fmt"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	type args struct {
		attempts int
		sleep    time.Duration
		fn       func() error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				attempts: 3,
				sleep:    1 * time.Second,
				fn: func() error {
					return fmt.Errorf("不可忽略错误")
				},
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				attempts: 3,
				sleep:    1 * time.Second,
				fn: func() error {
					return fmt.Errorf("这个错误可忽略 %w", CanStopErr)
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Retry(tt.args.attempts, tt.args.sleep, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Retry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
