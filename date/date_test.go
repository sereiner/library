package date

import (
	"testing"
)

func TestToday(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{name: "", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodayTimeStamp(); got != tt.want {
				t.Errorf("Today() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYesterdayTimeStamp(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{name: "", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YesterdayTimeStamp(); got != tt.want {
				t.Errorf("Yesterday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodayFormat(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "1", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodayFormat(); got != tt.want {
				t.Errorf("TodayFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNowFormat(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "1", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NowFormat(); got != tt.want {
				t.Errorf("NowFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodayTimeStamp(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodayTimeStamp(); got != tt.want {
				t.Errorf("TodayTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodayEndTimeStamp(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{name: "1", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodayEndTimeStamp(); got != tt.want {
				t.Errorf("TodayEndTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYesterdayEndTimeStamp(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{name: "1", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YesterdayEndTimeStamp(); got != tt.want {
				t.Errorf("YesterdayEndTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodayEndFormat(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "1", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TodayEndFormat(); got != tt.want {
				t.Errorf("TodayEndFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeStampToDate(t *testing.T) {
	type args struct {
		timeStamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{timeStamp: 1554566399}, want: "2019-04-06 23:59:59"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeStampToDate(tt.args.timeStamp); got != tt.want {
				t.Errorf("TimeStampToDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateToTimeStamp(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{name: "1", args: args{date: "2019-04-10T14:28:54+08:00"}, want: 1554566399},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateToTimeStamp(tt.args.date); got != tt.want {
				t.Errorf("DateToTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr2Stamp(t *testing.T) {
	type args struct {
		formatTimeStr string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{name: "1", args: args{formatTimeStr: "2019-04-10T14:28:54+08:00"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Str2Stamp(tt.args.formatTimeStr); got != tt.want {
				t.Errorf("Str2Stamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
