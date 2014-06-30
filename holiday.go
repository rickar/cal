// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"time"
)

// ObservedRule represents a rule for observing a holiday that falls
// on a weekend.
type ObservedRule int

const (
	ObservedNearest ObservedRule = iota // nearest weekday (Friday or Monday)
	ObservedExact                       // the exact day only
	ObservedMonday                      // Monday always
)

var (
	// United States holidays
	US_NewYear      = NewHoliday(time.January, 1)
	US_MLK          = NewHolidayFloat(time.January, time.Monday, 3)
	US_Presidents   = NewHolidayFloat(time.February, time.Monday, 3)
	US_Memorial     = NewHolidayFloat(time.May, time.Monday, -1)
	US_Independence = NewHoliday(time.July, 4)
	US_Labor        = NewHolidayFloat(time.September, time.Monday, 1)
	US_Columbus     = NewHolidayFloat(time.October, time.Monday, 2)
	US_Veterans     = NewHoliday(time.November, 11)
	US_Thanksgiving = NewHolidayFloat(time.November, time.Thursday, 4)
	US_Christmas    = NewHoliday(time.December, 25)
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

// matches determines whether the given date is the one referred to by the
// Holiday.
func (h *Holiday) matches(date time.Time) bool {
	if h.Month > 0 {
		if date.Month() != h.Month {
			return false
		}
		if h.Day > 0 {
			return date.Day() == h.Day
		}
		if h.Weekday > 0 && h.Offset != 0 {
			return IsWeekdayN(date, h.Weekday, h.Offset)
		}
	} else if h.Offset > 0 {
		return date.YearDay() == h.Offset
	}
	return false
}
