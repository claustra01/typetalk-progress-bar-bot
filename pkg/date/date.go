package date

import "time"

const (
	StartDate = "2024/06/15"
	EndDate   = "2025/03/09"
)

func StrToDate(str string) time.Time {
	t, _ := time.Parse("2006/01/02", str)
	return t
}

func GetProgress() float64 {
	now := time.Now()
	startDate := StrToDate(StartDate)
	endDate := StrToDate(EndDate)
	progress := float64(now.Sub(startDate)) / float64(endDate.Sub(startDate))
	return progress
}

func GetRemainingDays() int {
	now := time.Now()
	endDate := StrToDate(EndDate)
	remainingDays := int(endDate.Sub(now).Hours() / 24)
	return remainingDays
}
