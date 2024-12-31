package helper

import "time"

func StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func EndOfDay(date time.Time) time.Time {
	return StartOfDay(date).AddDate(0, 0, 1).Add(-time.Nanosecond)
}

func StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

func EndOfMonth(date time.Time) time.Time {
	return StartOfMonth(date).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func StartOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), 1, 1, 0, 0, 0, 0, date.Location())
}

func EndOfYear(date time.Time) time.Time {
	return StartOfYear(date).AddDate(1, 0, 0).Add(-time.Nanosecond)
}
