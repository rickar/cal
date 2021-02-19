// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ua provides holiday definitions for Ukraine.
package ua

import (
	"time"

	"github.com/Tamh/cal/v2"
	"github.com/Tamh/cal/v2/aa"
)

var (
	// Standard UA weekend substitution rules:
	//   Saturdays move to Monday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "Новий Рік", Type: cal.ObservancePublic, Observed: weekendAlt})

	// OrthodoxChristmas represents Orthodox Christmas on 7-Jan
	OrthodoxChristmas = &cal.Holiday{
		Name:     "Різдво",
		Type:     cal.ObservancePublic,
		Month:    time.January,
		Day:      7,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// WomensDay represents International Women's Day 8-Mar
	WomensDay = &cal.Holiday{
		Name:     "Міжнародний жіночий день",
		Type:     cal.ObservancePublic,
		Month:    time.March,
		Day:      8,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// OrthodoxEasterMonday represents Orthodox Easter Monday on the day after Julian Easter
	OrthodoxEasterMonday = &cal.Holiday{
		Name:   "Великдень",
		Type:   cal.ObservancePublic,
		Offset: 1,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// LabourDay represents Labour Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "День праці", Type: cal.ObservancePublic, Observed: weekendAlt})

	// LabourDay2 represents second Labour Day on 2-May
	LabourDay2 = &cal.Holiday{
		Name:     "День праці",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      2,
		Observed: weekendAlt,
		EndYear:  2017,
		Func:     cal.CalcDayOfMonth,
	}

	// VictoryDay represents Victory Day on 9-May
	VictoryDay = &cal.Holiday{
		Name:     "День перемоги над нацизмом у Другій світовій війні",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      9,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// OrthodoxPentecostMonday represents the day after Orthodox Pentecost, 49 days after Easter
	OrthodoxPentecostMonday = &cal.Holiday{
		Name:   "Трійця",
		Type:   cal.ObservancePublic,
		Offset: 49,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// ConstitutionDay represents Constitution Day on 9-May
	ConstitutionDay = &cal.Holiday{
		Name:     "День Конституції",
		Type:     cal.ObservancePublic,
		Month:    time.June,
		Day:      28,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// IndependenceDay represents Independence Day on 24-Aug
	IndependenceDay = &cal.Holiday{
		Name:     "День Незалежності",
		Type:     cal.ObservancePublic,
		Month:    time.August,
		Day:      24,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// DefenderOfUkraineDay represents Defender of Ukraine Day on 14-Oct
	DefenderOfUkraineDay = &cal.Holiday{
		Name:      "День захисника України",
		Type:      cal.ObservancePublic,
		StartYear: 2015,
		Month:     time.October,
		Day:       14,
		Observed:  weekendAlt,
		Func:      cal.CalcDayOfMonth,
	}

	// CatholicChristmas represents Christmas Day on 25-Dec
	CatholicChristmas = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Різдво", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		OrthodoxChristmas,
		WomensDay,
		OrthodoxEasterMonday,
		LabourDay,
		LabourDay2,
		VictoryDay,
		OrthodoxPentecostMonday,
		ConstitutionDay,
		IndependenceDay,
		DefenderOfUkraineDay,
		CatholicChristmas,
	}
)
