// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package mw provides holiday definitions for the Republic of Malawi.
package mw

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard mw weekend substitution rules:
	//   Saturdays move to Friday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "New Year's Day", Type: cal.ObservancePublic, Observed: weekendAlt, Func: cal.CalcDayOfMonth})

	// ChilembweDay represents John Chilembwe Day on the 15th of January
	ChilembweDay = &cal.Holiday{
		Name:     "John Chilembwe Day",
		Type:     cal.ObservancePublic,
		Month:    time.January,
		Day:      15,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// MartyrsDay represents Martyrs' Day on the 3rd of March
	MartyrsDay = &cal.Holiday{
		Name:     "Martyrs' Day",
		Type:     cal.ObservancePublic,
		Month:    time.March,
		Day:      3,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// GoodFriday represents Good Friday
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Good Friday", Type: cal.ObservancePublic})

	// Easter represents Easter Monday
	Easter = aa.EasterMonday.Clone(&cal.Holiday{Name: "Easter Monday", Type: cal.ObservancePublic})

	// LabourDay represents Labour Day on the 1st of May
	LabourDay = &cal.Holiday{
		Name:     "Labour Day",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      1,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// KamuzuDay represents President Kamuzu Banda's Birthday on the 14th of May
	KamuzuDay = &cal.Holiday{
		Name:     "President Kamuzu Banda's Birthday",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      14,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// MothersDay represents Mother's Day on the 15th of October
	MothersDay = &cal.Holiday{
		Name:     "Mother's Day",
		Type:     cal.ObservancePublic,
		Month:    time.October,
		Day:      15,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// IndependenceDay represents Independence Day on the 6th of July
	IndependenceDay = &cal.Holiday{
		Name:     "Independence Day",
		Type:     cal.ObservancePublic,
		Month:    time.July,
		Day:      6,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// BoxingDay represents Christmas Boxing Day on the 26th of December
	BoxingDay = &cal.Holiday{
		Name:     "Christmas Boxing Day",
		Type:     cal.ObservancePublic,
		Month:    time.December,
		Day:      26,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on the 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Christmas Day", Type: cal.ObservancePublic, Observed: weekendAlt})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		MartyrsDay,
		MothersDay,
		ChilembweDay,
		KamuzuDay,
		IndependenceDay,
		LabourDay,
		Easter,
		GoodFriday,
		BoxingDay,
		ChristmasDay,
	}
)
