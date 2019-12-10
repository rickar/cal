package cal

import (
	"time"
)

// Poland holidays
var (
	PLNewYear                     = NewYear
	PLThreeKings                  = NewHoliday(time.January, 6)
	PLEasterMonday                = EasterMonday
	PLEasterSunday                = NewHolidayFunc(calculateEasterSunday)
	PLLabourDay                   = NewHoliday(time.May, 1)
	PLNationalDay                 = NewHoliday(time.May, 3)
	PLPentecostDay                = NewHolidayFunc(calculatePentecostDay)
	PLCorpusChristi               = NewHolidayFunc(calculateCorpusChristiDay)
	PLAssumptionBlessedVirginMary = NewHoliday(time.August, 15)
	PLAllSaints                   = NewHoliday(time.November, 1)
	PLNationalIndependenceDay     = NewHoliday(time.November, 11)
	PLChristmasDayOne             = Christmas
	PLChristmasDayTwo             = Christmas2
)

// AddPolandHolidays adds all Poland holidays to the Calendar
func AddPolandHolidays(c *Calendar) {
	c.AddHoliday(
		PLNewYear,
		PLThreeKings,
		PLEasterMonday,
		PLEasterSunday,
		PLLabourDay,
		PLNationalDay,
		PLPentecostDay,
		PLCorpusChristi,
		PLAssumptionBlessedVirginMary,
		PLAllSaints,
		PLNationalIndependenceDay,
		PLChristmasDayOne,
		PLChristmasDayTwo,
	)
}

func calculateEasterSunday(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	return easter.Month(), easter.Day()
}

func calculatePentecostDay(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +49)
	return em.Month(), em.Day()
}

func calculateCorpusChristiDay(year int, loc *time.Location) (time.Month, int) {

	easter := calculateEaster(year, loc)
	// 60 days after Easter Sunday
	em := easter.AddDate(0, 0, +60)
	return em.Month(), em.Day()
}
