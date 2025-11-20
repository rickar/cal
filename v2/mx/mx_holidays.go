// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package mx provides holiday definitions for Mexico.
package mx

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard MX weekend substitution rules:
	//   Sundays move to Monday
	//   Saturdays move to Friday
	weekendAlt = []cal.AltDay{
		{Day: time.Sunday, Offset: 1},
		{Day: time.Saturday, Offset: -1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "New Year's Day", Type: cal.ObservancePublic, Observed: weekendAlt})

	// ConstitutionDay represents Constitution Day on 5-Feb
	ConstitutionDay = &cal.Holiday{
		Name:     "Constitution Day",
		Type:     cal.ObservancePublic,
		Month:    time.February,
		Day:      5,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// BenitoJuarezDay represents Benito Juárez's Birthday Day on 21-Mar
	BenitoJuarezDay = &cal.Holiday{
		Name:     "Benito Juárez's Birthday",
		Type:     cal.ObservancePublic,
		Month:    time.March,
		Day:      21,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// LabourDay represents Labour Day on 1-May
	LabourDay = &cal.Holiday{
		Name:     "Labour Day",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      1,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// IndependenceDay represents Independence Day on 16-Sep
	IndependenceDay = &cal.Holiday{
		Name:     "Independence Day",
		Type:     cal.ObservancePublic,
		Month:    time.September,
		Day:      16,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// RevolutionDay represents Revolution Day on the 3rd Monday in November
	RevolutionDay = &cal.Holiday{
		Name:    "Revolution Day",
		Type:    cal.ObservancePublic,
		Month:   time.November,
		Offset:  3,
		Weekday: time.Monday,
		Func:    cal.CalcWeekdayOffset,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Christmas Day", Type: cal.ObservanceBank, Observed: weekendAlt})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		ConstitutionDay,
		BenitoJuarezDay,
		LabourDay,
		IndependenceDay,
		RevolutionDay,
		ChristmasDay,
	}
)
