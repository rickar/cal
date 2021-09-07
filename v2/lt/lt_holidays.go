// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package lt provides holiday definitions for Lithuania.
// Holidays that occur during the weekends are omitted.

package lt

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "Naujieji metai", Type: cal.ObservancePublic})

	// StateRestorationDay represents Day of Restoration of the State of Lithuania on 16-Feb
	StateRestorationDay = &cal.Holiday{
		Name:  "Lietuvos valstybės atkūrimo diena",
		Type:  cal.ObservancePublic,
		Month: time.February,
		Day:   16,
		Func:  cal.CalcDayOfMonth,
	}

	// IndependenceDay represents Independence Restoration Day on 11-Mar
	IndependenceDay = &cal.Holiday{
		Name:  "Lietuvos nepriklausomybės atkūrimo diena",
		Type:  cal.ObservancePublic,
		Month: time.March,
		Day:   11,
		Func:  cal.CalcDayOfMonth,
	}

	// EasterMonday represents Easter Monday on the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Antroji šv. Velykų diena", Type: cal.ObservancePublic})

	// LabourDay represents Labor Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "Tarptautinė darbo diena", Type: cal.ObservancePublic})

	// SaintJohnsEve represents Saint John's Eve on 24-Jun
	SaintJohnsEve = &cal.Holiday{
		Name:  "Rasos ir Joninių diena",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// StatehoodDay represents Statehood Day on 06-Jul
	StatehoodDay = &cal.Holiday{
		Name:  "Valstybės (Lietuvos Karaliaus Mindaugo karūnavimo ir Tautiškos giesmės) diena",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   06,
		Func:  cal.CalcDayOfMonth,
	}

	// AssumptionDay represents Assumption of Mary on 15-Aug
	AssumptionDay = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Žolinė (Švč. Mergelės Marijos ėmimo į dangų diena)", Type: cal.ObservancePublic})

	// AllSaintsDay represents All Saints' Day on 1-Nov
	AllSaintsDay = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Visų šventųjų diena", Type: cal.ObservancePublic})

	// AllSouls represents All Souls' Day on 2-Nov
	AllSoulsDay = &cal.Holiday{
		Name:  "Mirusiųjų atminimo (Vėlinių) diena",
		Month: time.November,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasEve represents Christmas Eve on 24-Dec
	ChristmasEve = &cal.Holiday{
		Name:  "Šv. Kūčios",
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasDayOne represents Christmas Day on 25-Dec
	ChristmasDayOne = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Šv. Kalėdos", Type: cal.ObservancePublic})

	// ChristmasDayTwo represents the second day of Christmas on 26-Dec
	ChristmasDayTwo = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Šv. Kalėdos (antra diena)", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		StateRestorationDay,
		IndependenceDay,
		EasterMonday,
		LabourDay,
		SaintJohnsEve,
		StatehoodDay,
		AssumptionDay,
		AllSaintsDay,
		AllSoulsDay,
		ChristmasEve,
		ChristmasDayOne,
		ChristmasDayTwo,
	}
)
