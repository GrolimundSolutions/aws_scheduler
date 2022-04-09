package schedulermain

import (
	"testing"
	"time"
)

func Test_getActuallyHour(t *testing.T) {
	type args struct {
		loc *time.Location
	}
	localLocation, _ := time.LoadLocation("Europe/Zurich")
	americanLocation, _ := time.LoadLocation("America/New_York")
	asianLocation, _ := time.LoadLocation("Asia/Shanghai")
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testLocal",
			args: args{
				loc: localLocation,
			},
			want: time.Now().In(localLocation).Hour(),
		},
		{
			name: "testAmerican",
			args: args{
				loc: americanLocation,
			},
			want: time.Now().In(americanLocation).Hour(),
		},
		{
			name: "testAsian",
			args: args{
				loc: asianLocation,
			},
			want: time.Now().In(asianLocation).Hour(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getActuallyHour(tt.args.loc); got != tt.want {
				t.Errorf("getActuallyHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDayOfWeek(t *testing.T) {
	type args struct {
		loc *time.Location
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testLocal",
			args: args{
				loc: time.Local,
			},
			want: int(time.Now().In(time.Local).Weekday()),
		},
		{
			name: "testUTC",
			args: args{
				loc: time.UTC,
			},
			want: int(time.Now().In(time.UTC).Weekday()),
		},
		{
			name: "testGMT",
			args: args{
				loc: time.FixedZone("GMT", 0),
			},
			want: int(time.Now().In(time.FixedZone("GMT", 0)).Weekday()),
		},
		{
			name: "testGMT+6",
			args: args{
				loc: time.FixedZone("GMT", 6),
			},
			want: int(time.Now().In(time.FixedZone("GMT", 6)).Weekday()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDayOfWeek(tt.args.loc); got != tt.want {
				t.Errorf("getDayOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
