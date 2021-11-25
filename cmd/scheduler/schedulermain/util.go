package schedulermain

import "time"

func getDayOfWeek(loc *time.Location) int {
	return int(time.Now().In(loc).Weekday())
}
func getActuallyHour(loc *time.Location) int {
	return time.Now().In(loc).Hour()
}
