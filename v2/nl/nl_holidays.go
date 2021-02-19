// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package nl provides holiday definitions for the Netherlands.
package nl

import (
	"time"

	"github.com/Tamh/cal/v2"
	"github.com/Tamh/cal/v2/aa"
)

var (
	// Nieuwjaar represents New Year's Day on 1-Jan
	Nieuwjaar = aa.NewYear.Clone(&cal.Holiday{Name: "Nieuwjaarsdag", Type: cal.ObservancePublic})

	// GoedeVrijdag represents Good Friday on the Friday before Easter
	GoedeVrijdag = aa.GoodFriday.Clone(&cal.Holiday{Name: "Goede Vrijdag", Type: cal.ObservancePublic})

	// TweedePaasdag represents Easter Monday on the day after Easter
	TweedePaasdag = aa.EasterMonday.Clone(&cal.Holiday{Name: "Tweede Paasdag", Type: cal.ObservancePublic})

	// Koningsdag represents King's Day on 27-Apr
	Koningsdag = &cal.Holiday{
		Name:     "Koningsdag",
		Month:    time.April,
		Day:      27,
		Func:     cal.CalcDayOfMonth,
		Observed: []cal.AltDay{{Day: time.Sunday, Offset: -1}},
	}

	// BevrijdingsDag represents Liberation Day on 5-May
	BevrijdingsDag = &cal.Holiday{
		Name:  "Bevrijdingsdag",
		Month: time.May,
		Day:   5,
		Func:  cal.CalcDayOfMonth,
	}

	// Hemelvaart represents Ascension Day on the 39th day after Easter
	Hemelvaart = aa.AscensionDay.Clone(&cal.Holiday{Name: "Hemelvaartsdag", Type: cal.ObservancePublic})

	// TweedePinksterDag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	TweedePinksterDag = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Tweede Pinksterdag", Type: cal.ObservancePublic})

	// EersteKerstdag represents Christmas Day on 25-Dec
	EersteKerstdag = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Eerste Kerstdag", Type: cal.ObservancePublic})

	// TweedeKerstdag represents Boxing Day on 26-Dec
	TweedeKerstdag = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Tweede Kerstdag", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Nieuwjaar,
		GoedeVrijdag,
		TweedePaasdag,
		Koningsdag,
		BevrijdingsDag,
		Hemelvaart,
		TweedePinksterDag,
		EersteKerstdag,
		TweedeKerstdag,
	}
)
