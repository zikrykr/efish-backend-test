package helper

import "time"

func ParseStringToTime(dateStr string) time.Time {
	date, _ := time.Parse(time.RFC3339, dateStr)

	return date
}
