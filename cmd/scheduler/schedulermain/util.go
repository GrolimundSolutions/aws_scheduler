package schedulermain

import "time"

// getDayOfWeek returns the day of the week for now.
func getDayOfWeek(loc *time.Location) int {
	return int(time.Now().In(loc).Weekday())
}

// getActuallyHour returns the hour of the day for now.
func getActuallyHour(loc *time.Location) int {
	return time.Now().In(loc).Hour()
}
