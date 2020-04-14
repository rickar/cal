// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays in South Africa
// Reference https://en.wikipedia.org/wiki/Public_holidays_in_South_Africa
var (
	ZANewYear             = NewYear
	ZAHumanRights         = NewHoliday(time.March, 21)
	ZAGoodFriday          = GoodFriday
	ZAEasterMonday        = NewHolidayFunc(calculateEasterMonday)
	ZAFreedom             = NewHoliday(time.April, 27)
	ZAWorkers             = NewHoliday(time.May, 1)
	ZAYouth               = NewHoliday(time.June, 16)
	ZANationalWomens      = NewHoliday(time.August, 9)
	ZAHeritage            = NewHoliday(time.September, 24)
	ZADayOfReconciliation = NewHoliday(time.December, 16)
	ZAChristmas           = Christmas
	ZADayOfGoodwill       = Christmas2
)

// AddSouthAfricanHolidays adds all ZA holidays to the Calendar
func AddSouthAfricanHolidays(cal *Calendar) {
	cal.AddHoliday(
		ZANewYear,
		ZAHumanRights,
		ZAGoodFriday,
		ZAEasterMonday,
		ZAFreedom,
		ZAWorkers,
		ZAYouth,
		ZANationalWomens,
		ZAHeritage,
		ZADayOfReconciliation,
		ZAChristmas,
		ZADayOfGoodwill,
	)
}
