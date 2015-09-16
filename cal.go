// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"time"
)

// IsWeekend reports whether the given date falls on a weekend.
func IsWeekend(date time.Time) bool {
	day := date.Weekday()
	return day == time.Saturday || day == time.Sunday
}

// IsWeekdayN reports whether the given date is the nth occurrence of the
// day in the month.
//
// The value of n affects the direction of counting:
//   n > 0: counting begins at the first day of the month.
//   n == 0: the result is always false.
//   n < 0: counting begins at the end of the month.
func IsWeekdayN(date time.Time, day time.Weekday, n int) bool {
	cday := date.Weekday()
	if cday != day || n == 0 {
		return false
	}

	if n > 0 {
		return (date.Day()-1)/7 == (n - 1)
	} else {
		last := LastOfTheMonth(date)

		for {
			if last.Weekday() == day {
				n++
			}
			if n == 0 {
				break
			}
			last = last.AddDate(0, 0, -1)
		}

		return last.Day() == date.Day()
	}
}

// Calendar represents a yearly calendar with a list of holidays.
type Calendar struct {
	holidays [13][]Holiday // 0 for offset based holidays, 1-12 for month based
	Observed ObservedRule
}

// NewCalendar creates a new Calendar with no holidays defined.
func NewCalendar() *Calendar {
	c := &Calendar{}
	for i := range c.holidays {
		c.holidays[i] = make([]Holiday, 0, 2)
	}
	return c
}

// AddHoliday adds a holiday to the calendar's list.
func (c *Calendar) AddHoliday(h Holiday) {
	c.holidays[h.Month] = append(c.holidays[h.Month], h)
}

// IsHoliday reports whether a given date is a holiday. It does not account
// for the observation of holidays on alternate days.
func (c *Calendar) IsHoliday(date time.Time) bool {
	idx := date.Month()
	for i := range c.holidays[idx] {
		if c.holidays[idx][i].Matches(date) {
			return true
		}
	}
	for i := range c.holidays[0] {
		if c.holidays[0][i].Matches(date) {
			return true
		}
	}
	return false
}

// IsWorkday reports whether a given date is a work day (business day).
func (c *Calendar) IsWorkday(date time.Time) bool {
	if IsWeekend(date) || c.IsHoliday(date) {
		return false
	}

	if c.Observed == ObservedExact {
		return true
	}

	day := date.Weekday()
	if c.Observed == ObservedMonday && day == time.Monday {
		sun := date.AddDate(0, 0, -1)
		sat := date.AddDate(0, 0, -2)
		return !c.IsHoliday(sat) && !c.IsHoliday(sun)
	} else if c.Observed == ObservedNearest {
		if day == time.Friday {
			sat := date.AddDate(0, 0, 1)
			return !c.IsHoliday(sat)
		} else if day == time.Monday {
			sun := date.AddDate(0, 0, -1)
			return !c.IsHoliday(sun)
		}
	}

	return true
}

// countWorkdays reports the number of workdays from the given date to the end
// of the month.
func (c *Calendar) countWorkdays(dt time.Time, month time.Month) int {
	n := 0
	for ; month == dt.Month(); dt = dt.AddDate(0, 0, 1) {
		if c.IsWorkday(dt) {
			n++
		}
	}
	return n
}

// Workdays reports the total number of workdays for the given year and month.
func (c *Calendar) Workdays(year int, month time.Month) int {
	return c.countWorkdays(time.Date(year, month, 1, 12, 0, 0, 0, time.UTC), month)
}

// WorkdaysRemain reports the total number of remaining workdays in the month
// for the given date.
func (c *Calendar) WorkdaysRemain(date time.Time) int {
	return c.countWorkdays(date.AddDate(0, 0, 1), date.Month())
}

// WorkdayN reports the day of the month that corresponds to the nth workday
// for the given year and month.
//
// The value of n affects the direction of counting:
//   n > 0: counting begins at the first day of the month.
//   n == 0: the result is always 0.
//   n < 0: counting begins at the end of the month.
func (c *Calendar) WorkdayN(year int, month time.Month, n int) int {
	var date time.Time
	var add int
	if n == 0 {
		return 0
	}

	if n > 0 {
		date = time.Date(year, month, 1, 12, 0, 0, 0, time.UTC)
		add = 1
	} else {
		date = time.Date(year, month+1, 1, 12, 0, 0, 0, time.UTC).AddDate(0, 0, -1)
		add = -1
		n = -n
	}

	ndays := 0
	for ; month == date.Month(); date = date.AddDate(0, 0, add) {
		if c.IsWorkday(date) {
			ndays++
			if ndays == n {
				return date.Day()
			}
		}
	}
	return 0
}

// FirstOfTheMonth returns a date that occurs on the first of the month of the
// supplied date. For example, if the supplied date is 6/13/2015 then this function
// would return 6/1/2015 (at the same time as the former)
func FirstOfTheMonth(d time.Time) time.Time {
	if d.IsZero() {
		return d
	}

	day := d.Day()
	return d.AddDate(0, 0, 1-day)
}

// LastOfTheMonth returns a date that occurs on the last day of the month of the
// supplied date. For example, if the supplied date is 6/13/2015 then this function
// would return 6/30/2015 (at the same time as the former)
func LastOfTheMonth(d time.Time) time.Time {
	if d.IsZero() {
		return d
	}

	day := FirstOfTheMonth(d)
	day = day.AddDate(0, 1, -1)
	return day
}
