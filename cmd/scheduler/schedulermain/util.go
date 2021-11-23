package schedulermain

import "time"

func getDayOfWeek() int {
	loc, _ := time.LoadLocation("Europe/Zurich")
	return int(time.Now().In(loc).Weekday())
}
func getActuallyHour() int {
	loc, _ := time.LoadLocation("Europe/Zurich")
	return time.Now().In(loc).Hour()
}
