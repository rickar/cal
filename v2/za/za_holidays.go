// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package za provides holiday definitions for South Africa.
package za

import (
	"time"

	"github.com/Tamh/cal/v2"
	"github.com/Tamh/cal/v2/aa"
)

var (
	// Standard ZA weekend substitution rules:
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "New Year's Day", Type: cal.ObservancePublic, Observed: weekendAlt})

	// HumanRightsDay represents Human Rights Day on 21-Mar
	HumanRightsDay = &cal.Holiday{
		Name:     "Human Rights Day",
		Type:     cal.ObservancePublic,
		Month:    time.March,
		Day:      21,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Good Friday", Type: cal.ObservancePublic})

	// FamilyDay represents Family Day on the Monday after Easter
	FamilyDay = aa.EasterMonday.Clone(&cal.Holiday{Name: "Family Day", Type: cal.ObservancePublic})

	// FreedomDay represents Freedom Day on 27-Apr
	FreedomDay = &cal.Holiday{
		Name:     "Freedom Day",
		Type:     cal.ObservancePublic,
		Month:    time.April,
		Day:      27,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// WorkersDay represents Workers' Day on 1-May
	WorkersDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "Workers' Day", Type: cal.ObservancePublic, Observed: weekendAlt})

	// YouthDay represents Youth Day on 16-June
	YouthDay = &cal.Holiday{
		Name:     "Youth Day",
		Type:     cal.ObservancePublic,
		Month:    time.June,
		Day:      16,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// WomensDay represents National Women's Day on 9-Aug
	WomensDay = &cal.Holiday{
		Name:     "National Women's Day",
		Type:     cal.ObservancePublic,
		Month:    time.August,
		Day:      9,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// HeritageDay represents Heritage Day on 24-Sep
	HeritageDay = &cal.Holiday{
		Name:     "Heritage Day",
		Type:     cal.ObservancePublic,
		Month:    time.September,
		Day:      24,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// ReconciliationDay represents Day of Reconciliation on 16-Dec
	ReconciliationDay = &cal.Holiday{
		Name:     "Reconciliation Day",
		Type:     cal.ObservancePublic,
		Month:    time.December,
		Day:      16,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Christmas Day", Type: cal.ObservanceBank, Observed: weekendAlt})

	// GoodwillDay represents Day of Goodwill on 26-Dec
	GoodwillDay = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Day of Goodwill", Type: cal.ObservanceBank,
		Observed: []cal.AltDay{
			{Day: time.Sunday, Offset: 1},
			{Day: time.Monday, Offset: 1},
		}})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		HumanRightsDay,
		GoodFriday,
		FamilyDay,
		FreedomDay,
		WorkersDay,
		YouthDay,
		WomensDay,
		HeritageDay,
		ReconciliationDay,
		ChristmasDay,
		GoodwillDay,
	}
)
