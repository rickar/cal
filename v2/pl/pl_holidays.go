// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package pl provides holiday definitions for Poland.
package pl

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "Nowy Rok", Type: cal.ObservancePublic})

	// ThreeKings represents Epiphany on 6-Jan
	ThreeKings = aa.Epiphany.Clone(&cal.Holiday{Name: "Święto Trzech Króli", Type: cal.ObservancePublic})

	// EasterMonday represents Easter Monday on the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "drugi dzień Wielkiej Nocy", Type: cal.ObservancePublic})

	// LabourDay represents Labor Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "Święto Państwowe", Type: cal.ObservancePublic})

	// ConstitutionDay represents Constitution Day on 3-May
	ConstitutionDay = &cal.Holiday{
		Name:  "Święto Narodowe Trzeciego Maja",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   3,
		Func:  cal.CalcDayOfMonth,
	}

	// CorpusChristi represents Corpus Christi on the 60th day after Easter
	CorpusChristi = aa.CorpusChristi.Clone(&cal.Holiday{Name: "dzień Bożego Ciała", Type: cal.ObservancePublic})

	// AssumptionBlessedVirginMary represents Assumption of Mary on 15-Aug
	AssumptionBlessedVirginMary = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Wniebowzięcie Najświętszej Maryi Panny", Type: cal.ObservancePublic})

	// AllSaints represents All Saints' Day on 1-Nov
	AllSaints = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Wszystkich Świętych", Type: cal.ObservancePublic})

	// NationalIndependenceDay represents National Independence Day on 11-Nov
	NationalIndependenceDay = aa.ArmisticeDay.Clone(&cal.Holiday{Name: "Narodowe Święto Niepodległości", Type: cal.ObservancePublic})

	// ChristmasEve represents Christmas Eve on 24-Dec
	ChristmasEve = &cal.Holiday{
		Name:      "Wigilia Bożego Narodzenia",
		Month:     time.December,
		Day:       24,
		StartYear: 2025,
		Func:      cal.CalcDayOfMonth,
	}

	// ChristmasDayOne represents Christmas Day on 25-Dec
	ChristmasDayOne = aa.ChristmasDay.Clone(&cal.Holiday{Name: "pierwszy dzień Bożego Narodzenia", Type: cal.ObservancePublic})

	// ChristmasDayTwo represents the second day of Christmas on 26-Dec
	ChristmasDayTwo = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "drugi dzień Bożego Narodzenia", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		ThreeKings,
		EasterMonday,
		LabourDay,
		ConstitutionDay,
		CorpusChristi,
		AssumptionBlessedVirginMary,
		AllSaints,
		NationalIndependenceDay,
		ChristmasEve,
		ChristmasDayOne,
		ChristmasDayTwo,
	}
)
