// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ch provides holiday definitions for Switzerland.
package ch

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Neujahr represents New Year's Day on 1-Jan
	Neujahr = aa.NewYear.Clone(&cal.Holiday{Name: "Neujahrstag", Type: cal.ObservancePublic})

	// Berchtoldstag represents an Alemannic holiday on 2-Jan
	Berchtoldstag = &cal.Holiday{
		Name:  "Berchtoldstag",
		Month: time.January,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// HeiligeDreiKoenige represents Epiphany on 6-Jan
	HeiligeDreiKoenige = aa.Epiphany.Clone(&cal.Holiday{Name: "Heilige Drei Könige", Type: cal.ObservancePublic})

	// Josefstag represents Feast of Saint Joseph on 19-Mar
	Josefstag = &cal.Holiday{
		Name:  "Josefstag",
		Month: time.March,
		Day:   19,
		Func:  cal.CalcDayOfMonth,
	}

	// Karfreitag represents Good Friday on the Friday before Easter
	Karfreitag = aa.GoodFriday.Clone(&cal.Holiday{Name: "Karfreitag", Type: cal.ObservancePublic})

	// Ostermontag represents Easter Monday on the day after Easter
	Ostermontag = aa.EasterMonday.Clone(&cal.Holiday{Name: "Ostermontag", Type: cal.ObservancePublic})

	// TagderArbeit represents Labour Day on 1-May
	TagderArbeit = aa.WorkersDay.Clone(&cal.Holiday{Name: "Tag der Arbeit", Type: cal.ObservancePublic})

	// Auffahrt represents Ascension Day on the 39th day after Easter
	Auffahrt = aa.AscensionDay.Clone(&cal.Holiday{Name: "Auffahrt", Type: cal.ObservancePublic})

	// Pfingstmontag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	Pfingstmontag = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Pfingstmontag", Type: cal.ObservancePublic})

	// Fronleichnam represents Corpus Christi on the 60th day after Easter
	Fronleichnam = aa.CorpusChristi.Clone(&cal.Holiday{Name: "Fronleichnam", Type: cal.ObservancePublic})

	// Bundesfeiertag represents the official national day of Switzerland on the 1-Aug
	Bundesfeiertag = &cal.Holiday{
		Name:  "Bundesfeiertag",
		Month: time.August,
		Day:   1,
		Type:  cal.ObservancePublic,
		Func:  cal.CalcDayOfMonth,
	}

	// MariaHimmelfahrt represents Assumption of Mary on 15-Aug
	MariaHimmelfahrt = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Mariä Himmelfahrt", Type: cal.ObservancePublic})

	// Allerheiligen represents All Saints' Day on 1-Nov
	Allerheiligen = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Allerheiligen", Type: cal.ObservancePublic})

	//MariaEmpfangnis represents Immaculate Conception
	MariaEmpfangnis = aa.ImmaculateConception.Clone(&cal.Holiday{Name: "Mariä Empfängnis"})

	// Weihnachtstag represents Christmas Day on 25-Dec
	Weihnachtstag = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Weihnachtstag", Type: cal.ObservancePublic})

	// ZweiterWeihnachtsfeiertag represents Boxing Day on 26-Dec
	ZweiterWeihnachtsfeiertag = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Zweiter Weihnachtsfeiertag", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Neujahr,
		Berchtoldstag,
		HeiligeDreiKoenige,
		Josefstag,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		MariaEmpfangnis,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysZH provides a list of holidays in the Canton of Zurich
	HolidaysZH = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysBE provides a list of holidays in the Canton of Bern
	HolidaysBE = []*cal.Holiday{
		Neujahr,
		Berchtoldstag,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysLU provides a list of holidays in the Canton of Lucerne
	HolidaysLU = []*cal.Holiday{
		Neujahr,
		Josefstag,
		Karfreitag,
		Auffahrt,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysUR provides a list of holidays in the Canton of Uri
	HolidaysUR = []*cal.Holiday{
		Neujahr,
		HeiligeDreiKoenige,
		Josefstag,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysSZ provides a list of holidays in the Canton of Schwyz
	HolidaysSZ = []*cal.Holiday{
		Neujahr,
		HeiligeDreiKoenige,
		Josefstag,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysOW provides a list of holidays in the Canton of Obwalden
	HolidaysOW = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Auffahrt,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
	}

	// HolidaysNW provides a list of holidays in the Canton of Nidwalden
	HolidaysNW = []*cal.Holiday{
		Neujahr,
		Josefstag,
		Karfreitag,
		Auffahrt,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
	}

	// HolidaysGL provides a list of holidays in the Canton of Glarus
	HolidaysGL = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Allerheiligen,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysZG provides a list of holidays in the Canton of ZG
	HolidaysZG = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Allerheiligen,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysFR provides a list of holidays in the Canton of Fribourg
	HolidaysFR = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		TagderArbeit,
		Auffahrt,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		Weihnachtstag,
	}

	// HolidaysSO provides a list of holidays in the Canton of Solothurn
	HolidaysSO = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		TagderArbeit,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysBS provides a list of holidays in the Canton of Basel Stadt
	HolidaysBS = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysBL provides a list of holidays in the Canton of Basel Land
	HolidaysBL = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysSH provides a list of holidays in the Canton of Schaffhausen
	HolidaysSH = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysAR provides a list of holidays in the Canton of Appenzell Ausserrhoden
	HolidaysAR = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysAI provides a list of holidays in the Canton of Appenzell Innerrhoden
	HolidaysAI = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysSG provides a list of holidays in the Canton of St. Gallen
	HolidaysSG = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Allerheiligen,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysGR provides a list of holidays in the Canton of Grisons
	HolidaysGR = []*cal.Holiday{
		Neujahr,
		HeiligeDreiKoenige,
		Josefstag,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysAG provides a list of holidays in the Canton of Aargau
	HolidaysAG = []*cal.Holiday{
		Neujahr,
		Berchtoldstag,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysTG provides a list of holidays in the Canton of Thurgau
	HolidaysTG = []*cal.Holiday{
		Neujahr,
		Berchtoldstag,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysVD provides a list of holidays in the Canton of Vaud
	HolidaysVD = []*cal.Holiday{
		Neujahr,
		Berchtoldstag,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Bundesfeiertag,
		Weihnachtstag,
	}

	// HolidaysTI provides a list of holidays in the Canton of Ticino
	HolidaysTI = []*cal.Holiday{
		Neujahr,
		HeiligeDreiKoenige,
		Josefstag,
		Ostermontag,
		TagderArbeit,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysVS provides a list of holidays in the Canton of Valais
	HolidaysVS = []*cal.Holiday{
		Neujahr,
		Josefstag,
		Auffahrt,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		MariaEmpfangnis,
		Weihnachtstag,
	}

	// HolidaysNE provides a list of holidays in the Canton of Neuchâtel
	HolidaysNE = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		TagderArbeit,
		Auffahrt,
		Fronleichnam,
		Bundesfeiertag,
		Weihnachtstag,
	}

	// HolidaysGE provides a list of holidays in the Canton of Geneva
	HolidaysGE = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		Weihnachtstag,
	}

	// HolidaysJU provides a list of holidays in the Canton of Jura
	HolidaysJU = []*cal.Holiday{
		Neujahr,
		Berchtoldstag,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		Auffahrt,
		Pfingstmontag,
		Fronleichnam,
		Bundesfeiertag,
		MariaHimmelfahrt,
		Allerheiligen,
		Weihnachtstag,
	}
)
