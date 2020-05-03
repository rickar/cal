// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package at provides holiday definitions for Austria.
package at

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Neujahr represents New Year's Day on 1-Jan
	Neujahr = aa.NewYear.Clone("Neujahrstag", cal.ObservancePublic, nil)

	// HeiligeDreiKoenige represents Epiphany on 6-Jan
	HeiligeDreiKoenige = aa.Epiphany.Clone("Heilige Drei Könige", cal.ObservancePublic, nil)

	// Ostermontag represents Easter Monday on the day after Easter
	Ostermontag = aa.EasterMonday.Clone("Ostermontag", cal.ObservancePublic, nil)

	// TagderArbeit represents Labor Day on the first Monday in May
	TagderArbeit = aa.WorkersDay.Clone("Tag der Arbeit", cal.ObservancePublic, nil)

	// ChristiHimmelfahrt represents Ascension Day on the 39th day after Easter
	ChristiHimmelfahrt = aa.AscensionDay.Clone("Christi Himmelfahrt", cal.ObservancePublic, nil)

	// Pfingstmontag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	Pfingstmontag = aa.PentecostMonday.Clone("Pfingstmontag", cal.ObservancePublic, nil)

	// Fronleichnam represents Corpus Christi on the 60th day after Easter
	Fronleichnam = aa.CorpusChristi.Clone("Fronleichnam", cal.ObservancePublic, nil)

	// MariaHimmelfahrt represents Assumption of Mary on 15-Aug
	MariaHimmelfahrt = aa.AssumptionOfMary.Clone("Mariä Himmelfahrt", cal.ObservancePublic, nil)

	// Nationalfeiertag represents National Day on 26-Oct
	Nationalfeiertag = &cal.Holiday{
		Name:  "Nationalfeiertag",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   26,
		Func:  cal.CalcDayOfMonth,
	}

	// Allerheiligen represents All Saints' Day on 1-Nov
	Allerheiligen = aa.AllSaintsDay.Clone("Allerheiligen", cal.ObservancePublic, nil)

	// MariaEmpfaengnis represents Immaculate Conception on 8-Dec
	MariaEmpfaengnis = aa.ImmaculateConception.Clone("Mariä Empfängnis", cal.ObservancePublic, nil)

	// Christtag represents Christmas Day on 25-Dec
	Christtag = aa.ChristmasDay.Clone("Christtag", cal.ObservancePublic, nil)

	// Stefanitag represents St. Stephen's Day on 26-Dec
	Stefanitag = aa.ChristmasDay2.Clone("Stefanitag", cal.ObservancePublic, nil)

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
