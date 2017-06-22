// (c) 2017 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"time"
)

var (
	// Common holidays
	NewYear      = NewHoliday(time.January, 1)
	GoodFriday   = NewHolidayFunc(calculateGoodFriday)
	EasterMonday = NewHolidayFunc(calculateEasterMonday)
	Christmas    = NewHoliday(time.December, 25)
	Christmas2   = NewHoliday(time.December, 26)
)

var (
	// Holidays in Germany
	DE_Neujahr                = NewYear
	DE_KarFreitag             = GoodFriday
	DE_Ostermontag            = EasterMonday
	DE_TagderArbeit           = NewHoliday(time.May, 1)
	DE_Himmelfahrt            = NewHolidayFunc(calculateHimmelfahrt)
	DE_Pfingstmontag          = NewHolidayFunc(calculatePfingstMontag)
	DE_TagderDeutschenEinheit = NewHoliday(time.October, 3)
	DE_ErsterWeihnachtstag    = Christmas
	DE_ZweiterWeihnachtstag   = Christmas2
)

// AddGermanHolidays adds all German holidays to the Calendar
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

var (
	// European Central Bank Target2 holidays
	ECB_GoodFriday       = GoodFriday
	ECB_EasterMonday     = EasterMonday
	ECB_NewYearsDay      = NewYear
	ECB_LabourDay        = NewHoliday(time.May, 1)
	ECB_ChristmasDay     = Christmas
	ECB_ChristmasHoliday = Christmas2
)

// AddEcbHolidays adds all ECB Target2 holidays to the calendar
func AddEcbHolidays(c *Calendar) {
	c.AddHoliday(ECB_GoodFriday)
	c.AddHoliday(ECB_EasterMonday)
	c.AddHoliday(ECB_NewYearsDay)
	c.AddHoliday(ECB_LabourDay)
	c.AddHoliday(ECB_ChristmasDay)
	c.AddHoliday(ECB_ChristmasHoliday)
}

var (
	// Holidays in Great Britain
	GB_NewYear       = NewHolidayFunc(calculateNewYearsHoliday)
	GB_GoodFriday    = GoodFriday
	GB_EasterMonday  = EasterMonday
	GB_EarlyMay      = NewHolidayFloat(time.May, time.Monday, 1)
	GB_SpringHoliday = NewHolidayFloat(time.May, time.Monday, -1)
	GB_SummerHoliday = NewHolidayFloat(time.August, time.Monday, -1)
	GB_ChristmasDay  = Christmas
	GB_BoxingDay     = Christmas2
)

// AddBritishHolidays adds all British holidays to the Calender
func AddBritishHolidays(c *Calendar) {
	c.AddHoliday(GB_NewYear)
	c.AddHoliday(GB_GoodFriday)
	c.AddHoliday(GB_EasterMonday)
	c.AddHoliday(GB_EarlyMay)
	c.AddHoliday(GB_SpringHoliday)
	c.AddHoliday(GB_SummerHoliday)
	c.AddHoliday(GB_ChristmasDay)
	c.AddHoliday(GB_BoxingDay)
}

var (
	// Holidays in the Netherlands
	NL_Nieuwjaar       = NewYear
	NL_GoedeVrijdag    = GoodFriday
	NL_PaasMaandag     = EasterMonday
	NL_KoningsDag      = NewHolidayFunc(calculateKoningsDag)
	NL_BevrijdingsDag  = NewHoliday(time.May, 5)
	NL_Hemelvaart      = DE_Himmelfahrt
	NL_PinksterMaandag = DE_Pfingstmontag
	NL_EersteKerstdag  = Christmas
	NL_TweedeKerstdag  = Christmas2
)

// AddDutchHolidays adds all Dutch holidays to the Calendar
func AddDutchHolidays(c *Calendar) {
	c.AddHoliday(NL_Nieuwjaar)
	c.AddHoliday(NL_GoedeVrijdag)
	c.AddHoliday(NL_PaasMaandag)
	c.AddHoliday(NL_KoningsDag)
	c.AddHoliday(NL_BevrijdingsDag)
	c.AddHoliday(NL_Hemelvaart)
	c.AddHoliday(NL_PinksterMaandag)
	c.AddHoliday(NL_EersteKerstdag)
	c.AddHoliday(NL_TweedeKerstdag)
}

var (
	// United States holidays
	US_NewYear      = NewYear
	US_MLK          = NewHolidayFloat(time.January, time.Monday, 3)
	US_Presidents   = NewHolidayFloat(time.February, time.Monday, 3)
	US_Memorial     = NewHolidayFloat(time.May, time.Monday, -1)
	US_Independence = NewHoliday(time.July, 4)
	US_Labor        = NewHolidayFloat(time.September, time.Monday, 1)
	US_Columbus     = NewHolidayFloat(time.October, time.Monday, 2)
	US_Veterans     = NewHoliday(time.November, 11)
	US_Thanksgiving = NewHolidayFloat(time.November, time.Thursday, 4)
	US_Christmas    = Christmas
)

// AddUsHolidays adds all US holidays to the Calendar
func AddUsHolidays(cal *Calendar) {
	cal.AddHoliday(US_NewYear)
	cal.AddHoliday(US_MLK)
	cal.AddHoliday(US_Presidents)
	cal.AddHoliday(US_Memorial)
	cal.AddHoliday(US_Independence)
	cal.AddHoliday(US_Labor)
	cal.AddHoliday(US_Columbus)
	cal.AddHoliday(US_Veterans)
	cal.AddHoliday(US_Thanksgiving)
	cal.AddHoliday(US_Christmas)
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

func calculateGoodFriday(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// two days before Easter Sunday
	gf := easter.AddDate(0, 0, -2)
	return gf.Month(), gf.Day()
}

func calculateEasterMonday(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// the day after Easter Sunday
	em := easter.AddDate(0, 0, +1)
	return em.Month(), em.Day()
}

func calculateHimmelfahrt(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 39 days after Easter Sunday
	em := easter.AddDate(0, 0, +39)
	return em.Month(), em.Day()
}

func calculatePfingstMontag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +50)
	return em.Month(), em.Day()
}

//KoningsDag (kingsday) is April 27th, 26th if the 27th is a Sunday
func calculateKoningsDag(year int, loc *time.Location) (time.Month, int) {
	koningsDag := time.Date(year, time.April, 27, 0, 0, 0, 0, loc)
	if koningsDag.Weekday() == time.Sunday {
		koningsDag = koningsDag.AddDate(0, 0, -1)
	}
	return koningsDag.Month(), koningsDag.Day()
}

// NewYearsDay is the 1st of January unless the 1st is a Saturday or Sunday
// in which case it occurs on the following Monday.
func calculateNewYearsHoliday(year int, loc *time.Location) (time.Month, int) {
	day := time.Date(year, time.January, 1, 0, 0, 0, 0, loc)
	switch day.Weekday() {
	case time.Saturday:
		day = day.AddDate(0, 0, 2)
	case time.Sunday:
		day = day.AddDate(0, 0, 1)
	}
	return time.January, day.Day()
}
