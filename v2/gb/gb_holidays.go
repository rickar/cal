// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package gb provides holiday definitions for the United Kingdom.
package gb

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard UK weekend substitution rules:
	//   Saturdays move to Monday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "New Year's Day", Type: cal.ObservanceBank, Observed: weekendAlt})

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Good Friday", Type: cal.ObservanceBank})

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Easter Monday", Type: cal.ObservanceBank})

	// EarlyMay represents Early May on the first Monday of May
	EarlyMay = &cal.Holiday{
		Name:    "Early May",
		Type:    cal.ObservanceBank,
		Month:   time.May,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
		Except:  []int{2020},
	}

	// VEDay represents VE Day, the 75th anniversary of the end of WWII.
	VEDay = &cal.Holiday{
		Name:      "VE Day",
		Type:      cal.ObservanceBank,
		Month:     time.May,
		Day:       8,
		Func:      cal.CalcDayOfMonth,
		StartYear: 2020,
		EndYear:   2020,
	}

	// SpringHoliday represents Spring Bank Holiday on the last Monday of May
	SpringHoliday = &cal.Holiday{
		Name:    "Spring Bank Holiday",
		Type:    cal.ObservanceBank,
		Month:   time.May,
		Weekday: time.Monday,
		Offset:  -1,
		Func:    cal.CalcWeekdayOffset,
	}

	// SummerHolidayScotland represents Summer Bank Holiday in Scotland on the first Monday of August
	SummerHolidayScotland = &cal.Holiday{
		Name:    "Summer Bank Holiday",
		Type:    cal.ObservanceBank,
		Month:   time.August,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// SummerHoliday represents Summer Bank Holiday on the last Monday of August
	SummerHoliday = &cal.Holiday{
		Name:    "Summer Bank Holiday",
		Type:    cal.ObservanceBank,
		Month:   time.August,
		Weekday: time.Monday,
		Offset:  -1,
		Func:    cal.CalcWeekdayOffset,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Christmas Day", Type: cal.ObservanceBank, Observed: weekendAlt})

	// BoxingDay represents Boxing Day on 26-Dec
	BoxingDay = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Boxing Day", Type: cal.ObservanceBank,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: 2},
			{Day: time.Sunday, Offset: 2},
			{Day: time.Monday, Offset: 1}}})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		GoodFriday,
		EasterMonday,
		EarlyMay,
		VEDay,
		SpringHoliday,
		SummerHoliday,
		ChristmasDay,
		BoxingDay,
	}
)
