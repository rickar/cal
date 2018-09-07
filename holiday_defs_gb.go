package cal

import "time"

// British holidays
var (
	GBNewYear       = NewHolidayFunc(calculateNewYearsHoliday)
	GBGoodFriday    = GoodFriday
	GBEasterMonday  = EasterMonday
	GBEarlyMay      = NewHolidayFloat(time.May, time.Monday, 1)
	GBSpringHoliday = NewHolidayFloat(time.May, time.Monday, -1)
	GBSummerHoliday = NewHolidayFloat(time.August, time.Monday, -1)
	GBChristmasDay  = Christmas
	GBBoxingDay     = Christmas2
)

// AddBritishHolidays adds all British holidays to the Calender
func AddBritishHolidays(c *Calendar) {
	c.AddHoliday(
		GBNewYear,
		GBGoodFriday,
		GBEasterMonday,
		GBEarlyMay,
		GBSpringHoliday,
		GBSummerHoliday,
		GBChristmasDay,
		GBBoxingDay,
	)
}
