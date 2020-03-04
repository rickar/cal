package cal

import (
	"time"
)

// Slovak holidays
var (
	SKNewYear       = NewYear
	SKEpiphany      = NewHoliday(time.January, 6)
	SKGoodFriday    = GoodFriday
	SKEasterMonday  = EasterMonday
	SKLabor         = NewHoliday(time.May, 1)
	SKLiberation    = NewHoliday(time.May, 8)
	SKSaintsCyril   = NewHoliday(time.July, 5)
	SKSNP           = NewHoliday(time.August, 29)
	SKConstitution  = NewHoliday(time.September, 1)
	SKLadyOfSorrows = NewHoliday(time.September, 15)
	SKAllSaints     = NewHoliday(time.November, 1)
	SKFreedom       = NewHoliday(time.November, 17)
	SKChristmasEve  = NewHoliday(time.December, 24)
	SKChristmas     = Christmas
	SKSaintStephen  = Christmas2
)

// AddSlovakHolidays adds all Slovak holidays to the Calendar
func AddSlovakHolidays(c *Calendar) {
	c.AddHoliday(
		SKNewYear,
		SKEpiphany,
		SKGoodFriday,
		SKEasterMonday,
		SKLabor,
		SKLiberation,
		SKSaintsCyril,
		SKSNP,
		SKConstitution,
		SKLadyOfSorrows,
		SKAllSaints,
		SKFreedom,
		SKChristmasEve,
		SKChristmas,
		SKSaintStephen,
	)
}
