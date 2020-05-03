// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"time"
)

// ObservanceType represents the type of holiday or special day being observed.
type ObservanceType uint8

// Allowed values for ObservanceType
const (
	ObservancePublic    ObservanceType = iota // public / national / regional holiday
	ObservanceBank                            // bank holiday
	ObservanceReligious                       // religious holiday
	ObservanceOther                           // all other holidays (school, work, etc.)
)

// HolidayFn calculates the expected occurrence of a holiday for the given year.
// Returned times are the start of the day in the system location.
//
// Note that implementations of this function should always return the day
// that the holiday is expected to occur. Additional rules for substitution days
// and exception years will be applied by Holiday.Calc().
type HolidayFn func(h *Holiday, year int) time.Time

// AltDay represents an alternative day to observe a holiday.
type AltDay struct {
	Day    time.Weekday // the day to match
	Offset int          // the number of days to move the observance forward (positive) or backward (negative)
}

// Holiday holds information about the type and occurrence of a holiday.
type Holiday struct {
	Name        string         // name in local language
	Description string         // further details/notes
	Type        ObservanceType // type of day being observed
	StartYear   int            // the first year the holiday is observed
	EndYear     int            // the last year the holiday is observed
	Except      []int          // years where the holiday doesn't apply

	// calculation fields; required fields depend on rule being followed
	Month    time.Month   // the month the holiday occurs
	Day      int          // the day the holiday occurs
	Weekday  time.Weekday // the weekday the holiday occurs
	Offset   int          // the weekday or start date offset the holiday occurs
	Julian   bool         // the holiday is based on a Julian calendar
	Observed []AltDay     // the substitution days for the holiday
	Func     HolidayFn    // logic used to determine occurrences
}

// Clone returns a copy of the Holiday with the given name, type, and
// observance rules.
func (h *Holiday) Clone(name string, typ ObservanceType, observed []AltDay) *Holiday {
	return &Holiday{
		Name:        name,
		Description: h.Description,
		Type:        typ,
		StartYear:   h.StartYear,
		EndYear:     h.EndYear,
		Except:      h.Except,
		Month:       h.Month,
		Day:         h.Day,
		Weekday:     h.Weekday,
		Offset:      h.Offset,
		Julian:      h.Julian,
		Observed:    observed,
		Func:        h.Func,
	}
}

// Calc reports the actual and observed dates of a holiday for the given year.
// If the holiday is not observed in the given year, the zero time is returned.
//
// Returned times are the start of the day in the system location.
func (h *Holiday) Calc(year int) (actual, observed time.Time) {
	if (h.StartYear > 0 && year < h.StartYear) ||
		(h.EndYear > 0 && year > h.EndYear) || h.Func == nil {
		return time.Time{}, time.Time{}
	}
	if len(h.Except) > 0 {
		for _, ex := range h.Except {
			if year == ex {
				return time.Time{}, time.Time{}
			}
		}
	}
	actual = h.Func(h, year)

	if h.Observed == nil {
		return actual, actual
	}

	day := actual.Weekday()
	for _, o := range h.Observed {
		if o.Day == day {
			return actual, actual.AddDate(0, 0, o.Offset)
		}
	}
	return actual, actual
}

// CalcDayOfMonth calculates the occurrence of a holiday that is always a
// specific day of the month such as the 5th of November.
func CalcDayOfMonth(h *Holiday, year int) time.Time {
	return time.Date(year, h.Month, h.Day, 0, 0, 0, 0, DefaultLoc)
}

// CalcWeekdayOffset calculates the occurrence of a holiday that falls on the
// nth occurrence of a weekday in a month, such as the third wednesday of July.
func CalcWeekdayOffset(h *Holiday, year int) time.Time {
	return DayStart(WeekdayN(year, h.Month, h.Weekday, h.Offset))
}

// CalcWeekdayFrom calculates the occurrence of a holiday that falls on a
// specific day of the week following a starting date.
func CalcWeekdayFrom(h *Holiday, year int) time.Time {
	return DayStart(WeekdayNFrom(time.Date(year, h.Month, h.Day, 12, 0, 0, 0, DefaultLoc), h.Weekday, h.Offset))
}

// CalcEasterOffset calculates the occurrence of a holiday that is determined
// by its relation to the Easter holiday.
func CalcEasterOffset(h *Holiday, year int) time.Time {
	var month, day int
	if h.Julian {
		// Meeus algorithm
		a := year % 4
		b := year % 7
		c := year % 19
		d := (19*c + 15) % 30
		e := (2*a + 4*b - d + 34) % 7

		month = int(d+e+114) / 31
		day = ((d + e + 114) % 31) + 1
		day += 13
	} else {
		// Meeus/Jones/Butcher algorithm
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

		month = (h + l - 7*m + 114) / 31
		day = ((h + l - 7*m + 114) % 31) + 1
	}

	return time.Date(year, time.Month(month), day+h.Offset, 0, 0, 0, 0, DefaultLoc)
}
