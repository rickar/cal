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
	NewYear = aa.NewYear.Clone("Nowy Rok", cal.ObservancePublic, nil)

	// ThreeKings represents Epiphany on 6-Jan
	ThreeKings = aa.Epiphany.Clone("Święto Trzech Króli", cal.ObservancePublic, nil)

	// EasterMonday represents Easter Monday on the day after Easter
	EasterMonday = aa.EasterMonday.Clone("drugi dzień Wielkiej Nocy", cal.ObservancePublic, nil)

	// LabourDay represents Labor Day on 1-May
	LabourDay = aa.WorkersDay.Clone("Święto Państwowe", cal.ObservancePublic, nil)

	// ConstitutionDay represents Constitution Day on 3-May
	ConstitutionDay = &cal.Holiday{
		Name:  "Święto Narodowe Trzeciego Maja",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   3,
		Func:  cal.CalcDayOfMonth,
	}

	// CorpusChristi represents Corpus Christi on the 60th day after Easter
	CorpusChristi = aa.CorpusChristi.Clone("dzień Bożego Ciała", cal.ObservancePublic, nil)

	// AssumptionBlessedVirginMary represents Assumption of Mary on 15-Aug
	AssumptionBlessedVirginMary = aa.AssumptionOfMary.Clone("Wniebowzięcie Najświętszej Maryi Panny", cal.ObservancePublic, nil)

	// AllSaints represents All Saints' Day on 1-Nov
	AllSaints = aa.AllSaintsDay.Clone("Wszystkich Świętych", cal.ObservancePublic, nil)

	// NationalIndependenceDay represents National Independence Day on 11-Nov
	NationalIndependenceDay = aa.ArmisticeDay.Clone("Narodowe Święto Niepodległości", cal.ObservancePublic, nil)

	// ChristmasDayOne represents Christmas Day on 25-Dec
	ChristmasDayOne = aa.ChristmasDay.Clone("pierwszy dzień Bożego Narodzenia", cal.ObservancePublic, nil)

	// ChristmasDayTwo represents the second day of Christmas on 26-Dec
	ChristmasDayTwo = aa.ChristmasDay2.Clone("drugi dzień Bożego Narodzenia", cal.ObservancePublic, nil)

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
		ChristmasDayOne,
		ChristmasDayTwo,
	}
)
