package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	now := time.Now()
	return now.After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	weekday := t.Weekday()
	month := t.Month()
	hour, min, _ := t.Clock()
	day := t.Day()
	year := t.Year()
	res := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", weekday, month, day, year, hour, min)
	return res
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	layout := "2006-01-02"
	date := fmt.Sprintf("%d-09-15", year)
	t, _ := time.Parse(layout, date)
	return t
}
