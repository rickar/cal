// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"time"
)

// ObservedRule represents a rule for observing a holiday that falls
// on a weekend.
type ObservedRule int

// Observence
const (
	ObservedNearest ObservedRule = iota // nearest weekday (Friday or Monday)
	ObservedExact                       // the exact day only
	ObservedMonday                      // Monday always
)

// United States holidays
var (
	USNewYear      = NewHoliday(time.January, 1)
	USMLK          = NewHolidayFloat(time.January, time.Monday, 3)
	USPresidents   = NewHolidayFloat(time.February, time.Monday, 3)
	USMemorial     = NewHolidayFloat(time.May, time.Monday, -1)
	USIndependence = NewHoliday(time.July, 4)
	USLabor        = NewHolidayFloat(time.September, time.Monday, 1)
	USColumbus     = NewHolidayFloat(time.October, time.Monday, 2)
	USVeterans     = NewHoliday(time.November, 11)
	USThanksgiving = NewHolidayFloat(time.November, time.Thursday, 4)
	USChristmas    = NewHoliday(time.December, 25)
)

// Holiday holds information about the yearly occurrence of a holiday.
//
// A valid Holiday consists of:
// - Month and Day (such as March 14 for Pi Day)
// - Month, Weekday, and Offset (such as the second Monday of October for Columbus Day)
// - Offset (such as the 183rd day of the year for the start of the second half)
type Holiday struct {
	Month   time.Month
	Weekday time.Weekday
	Day     int
	Offset  int
}

// NewHoliday creates a new Holiday instance for an exact day of a month.
func NewHoliday(month time.Month, day int) Holiday {
	return Holiday{Month: month, Day: day}
}

// NewHolidayFloat creates a new Holiday instance for an offset-based day of
// a month.
func NewHolidayFloat(month time.Month, weekday time.Weekday, offset int) Holiday {
	return Holiday{Month: month, Weekday: weekday, Offset: offset}
}

// Matches determines whether the given date is the one referred to by the
// Holiday.
func (h *Holiday) Matches(date time.Time) bool {
	if h.Month > 0 {
		if date.Month() != h.Month {
			return false
		}
		if h.Day > 0 {
			return date.Day() == h.Day
		}
		if h.Weekday > 0 && h.Offset != 0 {
			r := IsWeekdayN(date, h.Weekday, h.Offset)
			return r
		}
	} else if h.Offset > 0 {
		return date.YearDay() == h.Offset
	}
	return false
}
