package cal

import "time"

// British holidays
var (
	GBNewYear       = NewHolidayFunc(calculateNewYearsHoliday)
	GBGoodFriday    = GoodFriday
	GBEasterMonday  = EasterMonday
	GBEarlyMay      = makeGBEarlyMay()
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

// Early May holiday in UK is usually 1st Monday in May, but in 2020 this is
// replaced by VE day on Fri 8th May
func makeGBEarlyMay() Holiday {
	hol := NewHolidayFloat(time.May, time.Monday, 1)
	hol.Func = func(year int, loc *time.Location) (month time.Month, day int) {
		if year == 2020 {
			// In 2020 force VE day
			return time.May, 8
		}
		// Else defer to floating holiday calculation
		return time.May, 0
	}
	return hol
}
