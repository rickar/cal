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

	// Target2 holidays
	ECB_GoodFriday       = CalculateGoodFriday(time.Now().Year(), time.Now().Location())
	ECB_EasterMonday     = CalculateEasterMonday(time.Now().Year(), time.Now().Location())
	ECB_NewYearsDay      = NewHoliday(time.January, 1)
	ECB_LabourDay        = NewHoliday(time.May, 1)
	ECB_ChristmasDay     = NewHoliday(time.December, 25)
	ECB_ChristmasHoliday = NewHoliday(time.December, 26)
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

func CalculateGoodFriday(year int, loc *time.Location) Holiday {
	easter := calculateEaster(year, loc)
	//Go the the day before  yesterday
	gf := easter.AddDate(0, 0, -2)
	return Holiday{Month: gf.Month(), Day: gf.Day()}
}

func CalculateEasterMonday(year int, loc *time.Location) Holiday {
	easter := calculateEaster(year, loc)
	//Go the the day after easter
	em := easter.AddDate(0, 0, +1)
	return Holiday{Month: em.Month(), Day: em.Day()}
}

func calculateEaster(year int, loc *time.Location) time.Time {
	y := year
	a := y % 19
	b := y / 100
	c := y % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451

	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
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
