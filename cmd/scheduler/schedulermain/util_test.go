package schedulermain

import (
	"testing"
	"time"
)

func Test_getActuallyHour(t *testing.T) {
	type args struct {
		loc *time.Location
	}
	loca, _ := time.LoadLocation("Europe/Zurich")
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				loc: loca,
			},
			want: time.Now().In(loca).Hour(),
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
			name: "test1",
			args: args{
				loc: time.Local,
			},
			want: int(time.Now().In(time.Local).Weekday()),
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
