// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package us provides holiday definitions for the United States of America.
package us

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard US weekend substitution rules:
	//   Saturdays move to Friday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: -1},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "New Year's Day", Type: cal.ObservancePublic, Observed: weekendAlt})

	// MlkDay represents Martin Luther King Jr. Day on the 3rd Monday in January
	MlkDay = &cal.Holiday{
		Name:    "Martin Luther King Jr. Day",
		Type:    cal.ObservancePublic,
		Month:   time.January,
		Weekday: time.Monday,
		Offset:  3,
		Func:    cal.CalcWeekdayOffset,
	}

	// PresidentsDay represents Presidents' Day on the 3rd Monday in February
	PresidentsDay = &cal.Holiday{
		Name:    "Presidents' Day",
		Type:    cal.ObservancePublic,
		Month:   time.February,
		Weekday: time.Monday,
		Offset:  3,
		Func:    cal.CalcWeekdayOffset,
	}

	// MemorialDay represents Memorial Day on the last Monday in May
	MemorialDay = &cal.Holiday{
		Name:    "Memorial Day",
		Type:    cal.ObservancePublic,
		Month:   time.May,
		Weekday: time.Monday,
		Offset:  -1,
		Func:    cal.CalcWeekdayOffset,
	}

	// Juneteenth represents Juneteenth on June 19th
	Juneteenth = &cal.Holiday{
		Name:      "Juneteenth",
		Type:      cal.ObservancePublic,
		Month:     time.June,
		Day:       19,
		Observed:  weekendAlt,
		Func:      cal.CalcDayOfMonth,
		StartYear: 2021,
	}

	// IndependenceDay represents Independence Day on 4-Jul
	IndependenceDay = &cal.Holiday{
		Name:     "Independence Day",
		Type:     cal.ObservancePublic,
		Month:    time.July,
		Day:      4,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// LaborDay represents Labor Day on the first Monday in September
	LaborDay = &cal.Holiday{
		Name:    "Labor Day",
		Type:    cal.ObservancePublic,
		Month:   time.September,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// ColumbusDay represents Columbus Day on the second Monday in October
	ColumbusDay = &cal.Holiday{
		Name:    "Columbus Day",
		Type:    cal.ObservancePublic,
		Month:   time.October,
		Weekday: time.Monday,
		Offset:  2,
		Func:    cal.CalcWeekdayOffset,
	}

	// VeteransDay represents Veterans Day on 11-Nov
	VeteransDay = &cal.Holiday{
		Name:     "Veterans Day",
		Type:     cal.ObservancePublic,
		Month:    time.November,
		Day:      11,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// ThanksgivingDay represents Thanksgiving Day on the fourth Thursday in November
	ThanksgivingDay = &cal.Holiday{
		Name:    "Thanksgiving Day",
		Type:    cal.ObservancePublic,
		Month:   time.November,
		Weekday: time.Thursday,
		Offset:  4,
		Func:    cal.CalcWeekdayOffset,
	}

	// DayAfterThanksgivingDay represents the day after Thanksgiving Day on the fourth Friday in November
	DayAfterThanksgivingDay = &cal.Holiday{
		Name:       "Day After Thanksgiving Day",
		Type:       cal.ObservancePublic,
		Month:      time.November,
		Weekday:    time.Thursday,
		Offset:     4,
		CalcOffset: 1,
		Func:       cal.CalcWeekdayOffset,
	}

	// ChristmasDay represents Christmas Day on the 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Christmas Day", Type: cal.ObservancePublic, Observed: weekendAlt})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		MlkDay,
		PresidentsDay,
		MemorialDay,
		Juneteenth,
		IndependenceDay,
		LaborDay,
		ColumbusDay,
		VeteransDay,
		ThanksgivingDay,
		ChristmasDay,
	}
)
