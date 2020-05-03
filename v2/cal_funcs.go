// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"time"
)

// IsWeekend reports whether the given time falls on a weekend.
func IsWeekend(t time.Time) bool {
	day := t.Weekday()
	return day == time.Saturday || day == time.Sunday
}

// WeekdayNFrom reports the nth occurrence of a weekday starting from the
// given time.
//
// The value of n affects the direction of counting:
//   n > 0: the result is the nth occurrence counting forwards from the
//     given time.
//   n == 0: the result is always the zero time.
//   n < 0: the result is the nth occurrence counting backwards from the
//     given time.
//
// The given time is considered an occurrence of the weekday; if n == 1 or -1
// and the given time matches the desired weekday, it will be returned
// unchanged.
//
// For non-zero results the time of day will be unchanged from the given time.
func WeekdayNFrom(t time.Time, day time.Weekday, n int) time.Time {
	if n == 0 {
		return time.Time{}
	} else if n < 0 {
		n++
	}

	wd := t.Weekday()
	if day < wd {
		return t.AddDate(0, 0, (7-int(wd-day))+((n-1)*7))
	} else if day > wd {
		return t.AddDate(0, 0, int(day-wd)+((n-1)*7))
	} else {
		if n <= 0 {
			n++
		}
		return t.AddDate(0, 0, (n-1)*7)
	}
}

// WeekdayN reports the nth occurrence of a weekday starting in the given year
// and month.
//
// The value of n is handled the same way as WeekdayNFrom.
//
// If n is larger than the number of occurrences of the weekday in the month,
// counting will continue into the following month(s) until the nth occurrence
// is found. For example, n=52 would report the first occurrence a year away.
//
// All non-zero results will be at noon in the DefaultLoc location/timezone.
func WeekdayN(year int, month time.Month, day time.Weekday, n int) time.Time {
	if n > 0 {
		return WeekdayNFrom(time.Date(year, month, 1, 12, 0, 0, 0, DefaultLoc), day, n)
	} else if n == 0 {
		return time.Time{}
	} else {
		return WeekdayNFrom(time.Date(year, month+1, 0, 12, 0, 0, 0, DefaultLoc), day, n)
	}
}

// IsWeekdayN reports whether the given date is the nth occurrence of the
// day in the month.
//
// The value of n is handled the same way as WeekdayNFrom.
//
// If n is larger than the number of occurrences of the weekday in the month,
// the result will be false.
func IsWeekdayN(t time.Time, day time.Weekday, n int) bool {
	if n == 0 || t.Weekday() != day {
		return false
	}

	if n > 0 {
		return (t.Day()-1)/7 == (n - 1)
	}

	want := WeekdayN(t.Year(), t.Month(), day, n)
	return want.Year() == t.Year() && want.Month() == t.Month() && want.Day() == t.Day()
}

// DayStart reports the start of the day in t (sets time fields zero).
func DayStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// DayEnd reports the end of the day in t (sets time fields to maximum).
func DayEnd(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// MonthStart reports the starting day of the month in t. The time portion is
// unchanged.
func MonthStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, t.Hour(), t.Minute(), t.Second(),
		t.Nanosecond(), t.Location())
}

// MonthEnd reports the ending day of the month in t. The time portion is
// unchanged.
func MonthEnd(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, t.Hour(), t.Minute(),
		t.Second(), t.Nanosecond(), t.Location())
}

// ReplaceLocation returns a copy of the given time in the given location
// without changing the time of day.
func ReplaceLocation(t time.Time, loc *time.Location) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
}

// JulianDayNumber reports the Julian Day Number for t. Note that Julian days
// start at 12:00 UTC.
func JulianDayNumber(t time.Time) int {
	// algorithm from http://www.tondering.dk/claus/cal/julperiod.php#formula
	utc := t.UTC()
	a := (14 - int(utc.Month())) / 12
	y := utc.Year() + 4800 - a
	m := int(utc.Month()) + 12*a - 3

	jdn := utc.Day() + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
	if utc.Hour() < 12 {
		jdn--
	}
	return jdn
}

// ModifiedJulianDayNumber reports the modified Julian Day Number for t. Note
// that modified Julian days start at 00:00 UTC.
func ModifiedJulianDayNumber(t time.Time) int {
	// algorithm from http://www.tondering.dk/claus/cal/julperiod.php#formula
	mjd := JulianDate(t)
	mjd -= 2400000.5
	return int(mjd)
}

// JulianDate reports the Julian Date (which includes time as a fraction) for t.
func JulianDate(t time.Time) float64 {
	utc := t.UTC()
	jdn := JulianDayNumber(t)
	if utc.Hour() < 12 {
		jdn++
	}
	return float64(jdn) + (float64(utc.Hour())-12.0)/24.0 +
		float64(utc.Minute())/1440.0 + float64(utc.Second())/86400.0
}

// ModifiedJulianDate reports the modified Julian Date Number for t. Note
// that modified Julian days start at 00:00 UTC.
func ModifiedJulianDate(t time.Time) float64 {
	// algorithm from http://www.tondering.dk/claus/cal/julperiod.php#formula
	mjd := JulianDate(t)
	mjd -= 2400000.5
	return mjd
}

// MaxTime reports the maximum time in the given list. If the list is empty,
// the zero value is returned.
func MaxTime(ts ...time.Time) time.Time {
	r := time.Time{}
	for _, t := range ts {
		if t.After(r) {
			r = t
		}
	}
	return r
}

// MinTime reports the minimum time in the given list. If the list is empty,
// the zero value is returned.
func MinTime(ts ...time.Time) time.Time {
	if len(ts) == 0 {
		return time.Time{}
	}
	r := ts[0]
	for _, t := range ts[1:] {
		if t.Before(r) {
			r = t
		}
	}
	return r
}
