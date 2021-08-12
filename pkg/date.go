package revcal

import (
	"fmt"
	"time"
)

type Date struct {
	year  int
	month int
	day   int
}

func (date *Date) Format() string {
	return fmt.Sprintf("%d %s %d", date.day, MonthNames[date.month-1], date.year)
}

// So we have arranged in the column of each month,
// the names of the real treasures of the rural economy.
// - Fabregi d'Églantine
func (date *Date) Symbol() string {
	dayNumber := 30*(date.month-1) + (date.day - 1)
	return Symbols[dayNumber]
}

func NewDate(date time.Time) Date {
	start := time.Date(1792, time.September, 22, 0, 0, 0, 0, time.UTC)
	days := int(date.Sub(start).Hours() / 24)
	year := 1
	yearLength := daysInRevolutionaryYear(year)
	for days >= yearLength {
		days -= yearLength
		year += 1
		yearLength = daysInRevolutionaryYear(year)
	}
	month := int(1 + (days / 30))
	day := int(1 + (days % 30))

	return Date{year: year, month: month, day: day}
}

// Historical leap years as 3, 7, 11, 15, and 20, and to
// treat subsequent years divisible by 4 that aren't divisible by 100
// (unless they are also divisible by 400), as per the Gregorian method
func isRepublicanLeapYear(year int) bool {
	return (year == 3 || year == 7 || year == 11 || year == 15 || year == 20) ||
		(year > 20 && (year%4 == 0 && (year%100 > 0) || year%400 == 0))
}

func daysInRevolutionaryYear(year int) int {
	if isRepublicanLeapYear(year) {
		return 366
	} else {
		return 365
	}
}
