// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ru provides holiday definitions for Russia.
package ru

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard RU weekend substitution rules:
	//   Saturdays move to Monday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{
		Name:     "Новый Год",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
	})

	// OrthodoxChristmas represents Orthodox Christmas on 7-Jan
	OrthodoxChristmas = &cal.Holiday{
		Name:     "Рождество Христово",
		Type:     cal.ObservancePublic,
		Month:    time.January,
		Day:      7,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// MilitaryDay represents Defender of the Fatherland Day 23-Feb
	MilitaryDay = &cal.Holiday{
		Name:     "День защитника Отечества",
		Type:     cal.ObservancePublic,
		Month:    time.February,
		Day:      23,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// WomensDay represents International Women's Day 8-Mar
	WomensDay = &cal.Holiday{
		Name:     "Международный женский день",
		Type:     cal.ObservancePublic,
		Month:    time.March,
		Day:      8,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// LabourDay represents Labour Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{
		Name:     "Праздник Весны и Труда",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt})

	// VictoryDay represents Victory Day on 9-May
	VictoryDay = &cal.Holiday{
		Name:     "День Победы советского народа в Великой Отечественной войне 1941-1945 годов",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      9,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// RussiasDay represents Russia's Day on 12-Jun
	RussiasDay = &cal.Holiday{
		Name:     "День России",
		Type:     cal.ObservancePublic,
		Month:    time.June,
		Day:      12,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// UnionDay represents National Union Day on 4-Nov
	UnionDay = &cal.Holiday{
		Name:     "День народного едиснтва России",
		Type:     cal.ObservancePublic,
		Month:    time.November,
		Day:      4,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		OrthodoxChristmas,
		MilitaryDay,
		WomensDay,
		LabourDay,
		VictoryDay,
		RussiasDay,
		UnionDay,
	}
)
