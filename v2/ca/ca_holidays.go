// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ca provides holiday definitions for Canada.
package ca

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard Canada weekend substitution rules:
	//   Saturdays move to Monday
	//   Sundays move to Monday
	// See: https://www.canada.ca/en/employment-social-development/services/labour-standards/reports/holidays.html#h2.6
	// Most employers in Canada, including the Federal Government, use the proceeding Monday.
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "New Year's Day", Type: cal.ObservancePublic,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: 2},
			{Day: time.Sunday, Offset: 1},
		}})

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Good Friday", Type: cal.ObservancePublic})

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Easter Monday", Type: cal.ObservancePublic})

	// VictoriaDay represents Victoria Day on the Monday before 25-May
	VictoriaDay = &cal.Holiday{
		Name:    "Victoria Day",
		Type:    cal.ObservancePublic,
		Month:   time.May,
		Weekday: time.Monday,
		Offset:  -2,
		Func:    cal.CalcWeekdayOffset,
	}

	// CanadaDay represents Canada Day on 1-July
	CanadaDay = &cal.Holiday{
		Name:     "Canada Day",
		Type:     cal.ObservancePublic,
		Month:    time.July,
		Day:      1,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// CivicDay represents Civic/Provincial Day on the first Monday of August
	CivicDay = &cal.Holiday{
		Name:    "Civic/Provincial Day",
		Type:    cal.ObservancePublic,
		Month:   time.August,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// LabourDay represents Labour Day on the first Monday of September
	LabourDay = &cal.Holiday{
		Name:    "Labour Day",
		Type:    cal.ObservancePublic,
		Month:   time.September,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayOffset,
	}

	// ThanksgivingDay represents ThanksgivingDay on the second Monday of October
	ThanksgivingDay = &cal.Holiday{
		Name:    "Thanksgiving Day",
		Type:    cal.ObservancePublic,
		Month:   time.October,
		Weekday: time.Monday,
		Offset:  2,
		Func:    cal.CalcWeekdayOffset,
	}

	// National Day for Truth and Reconciliation on 30-Sep
	NationalDayForTruthAndReconciliation = &cal.Holiday{
		StartYear: 2021,
		Name:      "National Day for Truth and Reconciliation",
		Month:     time.September,
		Day:       30,
		Func:      cal.CalcDayOfMonth,
		Observed:  weekendAlt,
	}

	// RemembranceDay represents Remembrance Day on 11-Nov
	RemembranceDay = aa.ArmisticeDay.Clone(&cal.Holiday{Name: "Remembrance Day", Type: cal.ObservancePublic, Observed: weekendAlt})

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Christmas Day", Type: cal.ObservanceBank,
		Observed: weekendAlt,
	})

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
		VictoriaDay,
		CanadaDay,
		CivicDay,
		LabourDay,
		ThanksgivingDay,
		NationalDayForTruthAndReconciliation,
		RemembranceDay,
		ChristmasDay,
		BoxingDay,
	}
)
