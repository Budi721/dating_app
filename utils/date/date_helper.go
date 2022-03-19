package date

import "time"

const (
	layoutISO = "2006-01-02"
)

func StringToDate(dateString string) time.Time {
	t, _ := time.Parse(layoutISO, dateString)
	return t
}
