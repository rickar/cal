// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package cz provides holiday definitions for the Czech Republic.
package cz

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "Nový rok", Type: cal.ObservancePublic})

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Velký pátek", Type: cal.ObservancePublic})

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Velikonoční pondělí", Type: cal.ObservancePublic})

	// LabourDay represents Labour Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "Svátek práce", Type: cal.ObservancePublic})

	// LiberationDay represents Liberation Day on 8-May
	LiberationDay = &cal.Holiday{
		Name:  "Den osvobození",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// SaintsCyrilMethodius represents Saints Cyril and Methodius Day on 5-Jul
	SaintsCyrilMethodius = &cal.Holiday{
		Name:  "Den slovanských věrozvěstů Cyrila a Metoděje",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   5,
		Func:  cal.CalcDayOfMonth,
	}

	// JanHusDay represents Jan Hus Day on 6-Jul
	JanHusDay = &cal.Holiday{
		Name:  "Den upálení mistra Jana Husa",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   6,
		Func:  cal.CalcDayOfMonth,
	}

	// SaintWenceslasDay represents Saint Wenceslas Day on 28-Sep
	SaintWenceslasDay = &cal.Holiday{
		Name:  "Den české státnosti",
		Type:  cal.ObservancePublic,
		Month: time.September,
		Day:   28,
		Func:  cal.CalcDayOfMonth,
	}

	// IndependenceDay represents Independent Czechoslovak State Day on 28-Oct
	IndependenceDay = &cal.Holiday{
		Name:  "Den vzniku samostatného československého státu",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   28,
		Func:  cal.CalcDayOfMonth,
	}

	// FreedomDay represents Struggle for Freedom and Democracy Day on 17-Nov
	FreedomDay = &cal.Holiday{
		Name:  "Den boje za svobodu a demokracii",
		Type:  cal.ObservancePublic,
		Month: time.November,
		Day:   17,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasEve represents Christmas Eve 24-Dec
	ChristmasEve = &cal.Holiday{
		Name:  "Štědrý den",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "1. svátek vánoční", Type: cal.ObservancePublic})

	// SaintStephensDay represents Saints Stephen's Day on 26-Dec
	SaintStephensDay = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "2. svátek vánoční", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		GoodFriday,
		EasterMonday,
		LabourDay,
		LiberationDay,
		SaintsCyrilMethodius,
		JanHusDay,
		SaintWenceslasDay,
		IndependenceDay,
		FreedomDay,
		ChristmasEve,
		ChristmasDay,
		SaintStephensDay,
	}
)
