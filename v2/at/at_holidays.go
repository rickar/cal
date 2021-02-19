// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package at provides holiday definitions for Austria.
package at

import (
	"time"

	"github.com/Tamh/cal/v2"
	"github.com/Tamh/cal/v2/aa"
)

var (
	// Neujahr represents New Year's Day on 1-Jan
	Neujahr = aa.NewYear.Clone(&cal.Holiday{Name: "Neujahrstag", Type: cal.ObservancePublic})

	// HeiligeDreiKoenige represents Epiphany on 6-Jan
	HeiligeDreiKoenige = aa.Epiphany.Clone(&cal.Holiday{Name: "Heilige Drei Könige", Type: cal.ObservancePublic})

	// Ostermontag represents Easter Monday on the day after Easter
	Ostermontag = aa.EasterMonday.Clone(&cal.Holiday{Name: "Ostermontag", Type: cal.ObservancePublic})

	// TagderArbeit represents Labor Day on the first Monday in May
	TagderArbeit = aa.WorkersDay.Clone(&cal.Holiday{Name: "Tag der Arbeit", Type: cal.ObservancePublic})

	// ChristiHimmelfahrt represents Ascension Day on the 39th day after Easter
	ChristiHimmelfahrt = aa.AscensionDay.Clone(&cal.Holiday{Name: "Christi Himmelfahrt", Type: cal.ObservancePublic})

	// Pfingstmontag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	Pfingstmontag = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Pfingstmontag", Type: cal.ObservancePublic})

	// Fronleichnam represents Corpus Christi on the 60th day after Easter
	Fronleichnam = aa.CorpusChristi.Clone(&cal.Holiday{Name: "Fronleichnam", Type: cal.ObservancePublic})

	// MariaHimmelfahrt represents Assumption of Mary on 15-Aug
	MariaHimmelfahrt = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Mariä Himmelfahrt", Type: cal.ObservancePublic})

	// Nationalfeiertag represents National Day on 26-Oct
	Nationalfeiertag = &cal.Holiday{
		Name:  "Nationalfeiertag",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   26,
		Func:  cal.CalcDayOfMonth,
	}

	// Allerheiligen represents All Saints' Day on 1-Nov
	Allerheiligen = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Allerheiligen", Type: cal.ObservancePublic})

	// MariaEmpfaengnis represents Immaculate Conception on 8-Dec
	MariaEmpfaengnis = aa.ImmaculateConception.Clone(&cal.Holiday{Name: "Mariä Empfängnis", Type: cal.ObservancePublic})

	// Christtag represents Christmas Day on 25-Dec
	Christtag = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Christtag", Type: cal.ObservancePublic})

	// Stefanitag represents St. Stephen's Day on 26-Dec
	Stefanitag = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Stefanitag", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Neujahr,
		HeiligeDreiKoenige,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		Fronleichnam,
		MariaHimmelfahrt,
		Nationalfeiertag,
		Allerheiligen,
		MariaEmpfaengnis,
		Christtag,
		Stefanitag,
	}
)
