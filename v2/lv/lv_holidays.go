// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package lv provides holiday definitions for Latvia.
// Holidays that occur during the weekends are omitted.

package lv

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Standard weekend substitution rules:
	//   Saturdays move to Monday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "Jaunais Gads", Type: cal.ObservancePublic})

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Lielā Piektdiena", Type: cal.ObservancePublic})

	// Easter represents Easter
	Easter = aa.Easter.Clone(&cal.Holiday{Name: "Pirmās Lieldienas", Type: cal.ObservancePublic})

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Otrās Lieldienas", Type: cal.ObservancePublic})

	// LabourDay represents International Workers' Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "Darba svētki, Latvijas Republikas Satversmes sapulces sasaukšanas diena", Type: cal.ObservancePublic})

	// StateRestorationDay represents Day of Restoration of the State of Latvia on 4th May
	StateRestorationDay = &cal.Holiday{
		Name:     "Latvijas Republikas Neatkarības deklarācijas pasludināšanas diena",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      4,
		Func:     cal.CalcDayOfMonth,
		Observed: weekendAlt,
	}

	// MidsummerEve represents evening on the summer solstice - 23th of June
	MidsummerEve = &cal.Holiday{
		Name:  "Līgo diena",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   23,
		Func:  cal.CalcDayOfMonth,
	}

	// MidsummeDay represents day after  the summer solstice - 24th of June
	MidsummeDay = &cal.Holiday{
		Name:  "Jāņu diena (vasaras saulgrieži)",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// StateProclamationDay represents Proclamation Day of the Republic of Latvia on 18th-Nov
	StateProclamationDay = &cal.Holiday{
		Name:     "Latvijas Republikas proklamēšanas diena",
		Type:     cal.ObservancePublic,
		Month:    time.November,
		Day:      18,
		Func:     cal.CalcDayOfMonth,
		Observed: weekendAlt,
	}

	// ChristmasEve represents Christmas Eve 24-Dec
	ChristmasEve = &cal.Holiday{
		Name:  "Ziemassvētku vakars (ziemas saulgrieži)",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Pirmie Ziemassvētki (ziemas saulgrieži)", Type: cal.ObservancePublic})

	// ChristmasDay2 represents Christmas Second Dat on 26-Dec
	ChristmasDay2 = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Otrie Ziemassvētki (ziemas saulgrieži)", Type: cal.ObservancePublic})

	// NewYearEve represents New Year's Eve on 31-Dec
	NewYearEve = &cal.Holiday{
		Name:  "Vecgada vakars",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		GoodFriday,
		Easter,
		EasterMonday,
		LabourDay,
		StateRestorationDay,
		MidsummerEve,
		MidsummeDay,
		StateProclamationDay,
		ChristmasEve,
		ChristmasDay,
		ChristmasDay2,
		NewYearEve,
	}
)
