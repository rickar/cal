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
	ECB_GoodFriday       = NewHolidayFunc(calculateGoodFriday)
	ECB_EasterMonday     = NewHolidayFunc(calculateEasterMonday)
	ECB_NewYearsDay      = NewHoliday(time.January, 1)
	ECB_LabourDay        = NewHoliday(time.May, 1)
	ECB_ChristmasDay     = NewHoliday(time.December, 25)
	ECB_ChristmasHoliday = NewHoliday(time.December, 26)

	// Holidays in Germany
	DE_Neujahr                = US_NewYear
	DE_KarFreitag             = NewHolidayFunc(calculateGoodFriday)
	DE_Ostermontag            = NewHolidayFunc(calculateEasterMonday)
	DE_TagderArbeit           = NewHoliday(time.May, 1)
	DE_Himmelfahrt            = NewHolidayFunc(calculateHimmelfahrt)
	DE_Pfingstmontag          = NewHolidayFunc(calculatePfingstMontag)
	DE_TagderDeutschenEinheit = NewHoliday(time.October, 3)
	DE_ErsterWeihnachtstag    = ECB_ChristmasDay
	DE_ZweiterWeihnachtstag   = ECB_ChristmasHoliday
)

// HolidayFn calculates the occurrence of a holiday for the given year.
// This is useful for holidays like Easter that depend on complex rules.
type HolidayFn func(year int, loc *time.Location) (month time.Month, day int)

// Holiday holds information about the yearly occurrence of a holiday.
//
// A valid Holiday consists of one of the following:
// - Month and Day (such as March 14 for Pi Day)
// - Month, Weekday, and Offset (such as the second Monday of October for Columbus Day)
// - Offset (such as the 183rd day of the year for the start of the second half)
// - Func (to calculate the holiday)
type Holiday struct {
	Month   time.Month
	Weekday time.Weekday
	Day     int
	Offset  int
	Func    HolidayFn

	// last values used to calculate month and day with Func
	lastYear int
	lastLoc  *time.Location
}

func calculateGoodFriday(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	//Go the the day before yesterday
	gf := easter.AddDate(0, 0, -2)
	return gf.Month(), gf.Day()
}

func calculateEasterMonday(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	//Go the the day after Easter
	em := easter.AddDate(0, 0, +1)
	return em.Month(), em.Day()
}

func calculateEaster(year int, loc *time.Location) time.Time {
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

	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
}

func calculateHimmelfahrt(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	//Go the the day after Easter
	em := easter.AddDate(0, 0, +39)
	return em.Month(), em.Day()
}

func calculatePfingstMontag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	//Go the the day after Easter
	em := easter.AddDate(0, 0, +50)
	return em.Month(), em.Day()
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

// NewHolidayFunc creates a new Holiday instance that uses a function to
// calculate the day and month.
func NewHolidayFunc(fn HolidayFn) Holiday {
	return Holiday{Func: fn}
}

// matches determines whether the given date is the one referred to by the
// Holiday.
func (h *Holiday) matches(date time.Time) bool {

	if h.Func != nil && (date.Year() != h.lastYear || date.Location() != h.lastLoc) {
		h.Month, h.Day = h.Func(date.Year(), date.Location())
		h.lastYear = date.Year()
		h.lastLoc = date.Location()
	}

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

func AddGermanHolidays(c *Calendar) {
	c.AddHoliday(DE_Neujahr)
	c.AddHoliday(DE_KarFreitag)
	c.AddHoliday(DE_Ostermontag)
	c.AddHoliday(DE_TagderArbeit)
	c.AddHoliday(DE_Himmelfahrt)
	c.AddHoliday(DE_Pfingstmontag)
	c.AddHoliday(DE_TagderDeutschenEinheit)
	c.AddHoliday(DE_ErsterWeihnachtstag)
	c.AddHoliday(DE_ZweiterWeihnachtstag)
}
