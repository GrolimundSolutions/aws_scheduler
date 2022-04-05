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
