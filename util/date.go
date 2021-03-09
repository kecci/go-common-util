package util

import "time"

func ConvertStringDateToTime(str string) (date time.Time) {
	date, _ = time.Parse(time.RFC3339, str)

	return date
}
