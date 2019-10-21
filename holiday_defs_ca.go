package cal

import "time"

// Canadian holidays
var (
	CANewYear        = NewYear
	CAFamilyDay      = NewHolidayFloat(time.February, time.Monday, 3)
	CAGoodFriday     = GoodFriday
	CAVictoriaDay    = NewHolidayFloat(time.May, time.Monday, 3)
	CACanadaDay      = NewHoliday(time.July, 1)
	CACivicHoliday   = NewHolidayFloat(time.August, time.Monday, 1)
	CALabourDay      = NewHolidayFloat(time.September, time.Monday, 1)
	CAThanksgiving   = NewHolidayFloat(time.October, time.Monday, 2)
	CARemembranceDay = NewHoliday(time.November, 11)
	CAChristmasDay   = Christmas
	CABoxingDay      = Christmas2
)

// AddCanadianHolidays adds all Canadian holidays to the Calender
func AddCanadianHolidays(c *Calendar) {
	c.AddHoliday(
		CANewYear,
		CAFamilyDay,
		CAGoodFriday,
		CAVictoriaDay,
		CACanadaDay,
		CACivicHoliday,
		CALabourDay,
		CAThanksgiving,
		CARemembranceDay,
		CAChristmasDay,
		CABoxingDay,
	)
}
