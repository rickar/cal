// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package de provides holiday definitions for Germany.
package de

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Neujahr represents New Year's Day on 1-Jan
	Neujahr = aa.NewYear.Clone(&cal.Holiday{Name: "Neujahrstag", Type: cal.ObservancePublic})

	// HeiligeDreiKoenige represents Epiphany on 6-Jan
	HeiligeDreiKoenige = aa.Epiphany.Clone(&cal.Holiday{Name: "Heilige Drei Könige", Type: cal.ObservancePublic})

	// Frauentag represents Women's Day on 8-Mar
	Frauentag = &cal.Holiday{
		Name:  "Frauentag",
		Type:  cal.ObservancePublic,
		Month: time.March,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// Karfreitag represents Good Friday on the Friday before Easter
	Karfreitag = aa.GoodFriday.Clone(&cal.Holiday{Name: "Karfreitag", Type: cal.ObservancePublic})

	// Ostermontag represents Easter Monday on the day after Easter
	Ostermontag = aa.EasterMonday.Clone(&cal.Holiday{Name: "Ostermontag", Type: cal.ObservancePublic})

	// TagderArbeit represents Labour Day on 1-May
	TagderArbeit = aa.WorkersDay.Clone(&cal.Holiday{Name: "Tag der Arbeit", Type: cal.ObservancePublic})

	// ChristiHimmelfahrt represents Ascension Day on the 39th day after Easter
	ChristiHimmelfahrt = aa.AscensionDay.Clone(&cal.Holiday{Name: "Christi Himmelfahrt", Type: cal.ObservancePublic})

	// Pfingstmontag represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	Pfingstmontag = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Pfingstmontag", Type: cal.ObservancePublic})

	// Fronleichnam represents Corpus Christi on the 60th day after Easter
	Fronleichnam = aa.CorpusChristi.Clone(&cal.Holiday{Name: "Fronleichnam", Type: cal.ObservancePublic})

	// Friedensfest represents the Augsburger Hohes Friedensfest on 8-Aug
	Friedensfest = &cal.Holiday{
		Name:      "Friedensfest",
		Type:      cal.ObservancePublic,
		Month:     time.August,
		Day:       8,
		Func:      cal.CalcDayOfMonth,
		StartYear: 1950,
	}

	// MariaHimmelfahrt represents Assumption of Mary on 15-Aug
	MariaHimmelfahrt = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Mariä Himmelfahrt", Type: cal.ObservancePublic})

	// Weltkindertag represents World Children's Day on 20-Sep
	Weltkindertag = &cal.Holiday{
		Name:      "Weltkindertag",
		Type:      cal.ObservancePublic,
		Month:     time.September,
		Day:       20,
		Func:      cal.CalcDayOfMonth,
		StartYear: 2019,
	}

	// DeutschenEinheit represents German Unity Day on 3-Oct
	DeutschenEinheit = &cal.Holiday{
		Name:  "Tag der Deutschen Einheit",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   3,
		Func:  cal.CalcDayOfMonth,
	}

	// Reformationstag represents Reformation Day on 31-Oct
	Reformationstag = &cal.Holiday{
		Name:  "Reformationstag",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	// Allerheiligen represents All Saints' Day on 1-Nov
	Allerheiligen = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Allerheiligen", Type: cal.ObservancePublic})

	// BussUndBettag represents Repentance and Prayer Day on the first Wednesday between 16-22 Nov
	BussUndBettag = &cal.Holiday{
		Name:    "Buß- und Bettag",
		Type:    cal.ObservancePublic,
		Month:   time.November,
		Day:     16,
		Weekday: time.Wednesday,
		Offset:  1,
		Func:    cal.CalcWeekdayFrom,
	}

	// Heiligabend represents Christmas Eve on 24-Dec
	Heiligabend = &cal.Holiday{
		Name:  "Heiligabend",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// Weihnachtstag represents Christmas Day on 25-Dec
	Weihnachtstag = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Weihnachtstag", Type: cal.ObservancePublic})

	// ZweiterWeihnachtsfeiertag represents Boxing Day on 26-Dec
	ZweiterWeihnachtsfeiertag = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Zweiter Weihnachtsfeiertag", Type: cal.ObservancePublic})

	// Silvester represents NewYear Eve on 31 Dec
	Silvester = &cal.Holiday{
		Name:  "Silvester",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysBW provides a list of holidays in the Baden-Württemberg region
	HolidaysBW = []*cal.Holiday{
		Neujahr,
		HeiligeDreiKoenige,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		Fronleichnam,
		DeutschenEinheit,
		Allerheiligen,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysBY provides a list of holidays in the Bayern region
	HolidaysBY = []*cal.Holiday{
		Neujahr,
		HeiligeDreiKoenige,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		Fronleichnam,
		DeutschenEinheit,
		Allerheiligen,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysBE provides a list of holidays in the Berlin region
	HolidaysBE = []*cal.Holiday{
		Neujahr,
		Frauentag,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysBB provides a list of holidays in the Brandenburg region
	HolidaysBB = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Reformationstag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysHB provides a list of holidays in the Bremen region
	HolidaysHB = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Reformationstag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysHH provides a list of holidays in the Hamburg region
	HolidaysHH = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Reformationstag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysHE provides a list of holidays in the Hessen region
	HolidaysHE = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		Fronleichnam,
		DeutschenEinheit,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysMV provides a list of holidays in the Mecklenburg-Vorpommern region
	HolidaysMV = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Reformationstag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysNI provides a list of holidays in the Niedersachsen region
	HolidaysNI = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Reformationstag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysNW provides a list of holidays in the Nordrhein-Westfalen region
	HolidaysNW = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		Fronleichnam,
		DeutschenEinheit,
		Allerheiligen,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysRP provides a list of holidays in the Rheinland-Pfalz region
	HolidaysRP = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		Fronleichnam,
		DeutschenEinheit,
		Allerheiligen,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysSL provides a list of holidays in the Saarland region
	HolidaysSL = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		Fronleichnam,
		MariaHimmelfahrt,
		DeutschenEinheit,
		Allerheiligen,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysSN provides a list of holidays in the Sachsen region
	HolidaysSN = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Reformationstag,
		BussUndBettag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysST provides a list of holidays in the Sachsen-Anhalt region
	HolidaysST = []*cal.Holiday{
		Neujahr,
		HeiligeDreiKoenige,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Reformationstag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysSH provides a list of holidays in the Schleswig-Holstein region
	HolidaysSH = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		DeutschenEinheit,
		Reformationstag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}

	// HolidaysTH provides a list of holidays in the Thüringen region
	HolidaysTH = []*cal.Holiday{
		Neujahr,
		Karfreitag,
		Ostermontag,
		TagderArbeit,
		ChristiHimmelfahrt,
		Pfingstmontag,
		Weltkindertag,
		DeutschenEinheit,
		Reformationstag,
		Weihnachtstag,
		ZweiterWeihnachtsfeiertag,
	}
)
