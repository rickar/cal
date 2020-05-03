// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package nz provides holiday definitions for New Zealand.
package nz

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard New Zealand weekend substitution rules:
	//   Saturdays move to Monday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone("New Year's Day", cal.ObservancePublic, weekendAlt)

	// DayAfterNewYear represents Day after New Year's Day on 2-Jan
	DayAfterNewYear = &cal.Holiday{
		Name:  "Day after New Year's Day",
		Type:  cal.ObservancePublic,
		Month: time.January,
		Day:   2,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: 2},
			{Day: time.Sunday, Offset: 2},
			{Day: time.Monday, Offset: 1}},
		Func: cal.CalcDayOfMonth,
	}

	// WaitangiDay represents Waitangi Day on 6-Feb
	WaitangiDay = &cal.Holiday{
		Name:     "Waitangi Day",
		Type:     cal.ObservancePublic,
		Month:    time.February,
		Day:      6,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone("Good Friday", cal.ObservancePublic, nil)

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone("Easter Monday", cal.ObservancePublic, nil)

	// AnzacDay represents ANZAC Day on 25-Apr
	AnzacDay = &cal.Holiday{
		Name:     "ANZAC Day",
		Type:     cal.ObservancePublic,
		Month:    time.April,
		Day:      25,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// QueensBirthday represents Queen's Birthday on the first Monday in June
	QueensBirthday = &cal.Holiday{
		Name:    "Queen's Birthday",
		Type:    cal.ObservancePublic,
		Month:   time.June,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// LabourDay represents Labour Day on the fourth Monday in October
	LabourDay = &cal.Holiday{
		Name:    "Labour Day",
		Type:    cal.ObservancePublic,
		Month:   time.October,
		Weekday: time.Monday,
		Offset:  4,
		Func:    cal.CalcWeekdayOffset,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone("Christmas Day", cal.ObservanceBank, weekendAlt)

	// BoxingDay represents Boxing Day on 26-Dec
	BoxingDay = aa.ChristmasDay2.Clone("Boxing Day", cal.ObservanceBank,
		[]cal.AltDay{
			{Day: time.Saturday, Offset: 2},
			{Day: time.Sunday, Offset: 2},
			{Day: time.Monday, Offset: 1}})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		DayAfterNewYear,
		WaitangiDay,
		GoodFriday,
		EasterMonday,
		AnzacDay,
		QueensBirthday,
		LabourDay,
		ChristmasDay,
		BoxingDay,
	}
)
