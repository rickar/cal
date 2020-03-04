package cal

import (
	"time"
)

// Czech holidays
var (
	CZNewYear        = NewYear
	CZGoodFriday     = GoodFriday
	CZEasterMonday   = EasterMonday
	CZLabor          = NewHoliday(time.May, 1)
	CZLiberation     = NewHoliday(time.May, 8)
	CZSaintsCyril    = NewHoliday(time.July, 5)
	CZJanHus         = NewHoliday(time.July, 6)
	CZSaintWenceslas = NewHoliday(time.September, 28)
	CZFoundationCZSK = NewHoliday(time.October, 28)
	CZFreedom        = NewHoliday(time.November, 17)
	CZChristmasEve   = NewHoliday(time.December, 24)
	CZChristmas      = Christmas
	CZSaintStephen   = Christmas2
)

// AddCzechHolidays adds all Czech holidays to the Calendar
func AddCzechHolidays(c *Calendar) {
	c.AddHoliday(
		CZNewYear,
		CZGoodFriday,
		CZEasterMonday,
		CZLabor,
		CZLiberation,
		CZSaintsCyril,
		CZJanHus,
		CZSaintWenceslas,
		CZFoundationCZSK,
		CZFreedom,
		CZChristmasEve,
		CZChristmas,
		CZSaintStephen,
	)
}
