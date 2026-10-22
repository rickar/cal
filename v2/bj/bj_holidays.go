// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package bj provides holiday definitions for Benin.
//
// The legal holidays are fixed by Law 90-019 of 27 July 1990, which makes
// them days off with pay. The annual traditional religions festival was
// added by Law 97-031 of 20 August 1997 and moved by Law 2024-32 of
// 2 September 2024. Holidays falling on a weekend are not moved.
//
// The Islamic holidays Jour du Ramadan (Eid al-Fitr), Jour de la Tabaski
// (Eid al-Adha) and Journée Maouloud (Prophet's Birthday) are legal holidays
// too, but are not included: their dates depend on the sighting of the moon
// and are announced by the government shortly in advance.
package bj

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "Fête du Nouvel An", Type: cal.ObservancePublic})

	// TraditionalReligionsDayEve represents the Thursday before the annual
	// traditional religions festival, a day off since 2025 (Law 2024-32)
	TraditionalReligionsDayEve = &cal.Holiday{
		Name:       "Fête annuelle des religions traditionnelles",
		Type:       cal.ObservancePublic,
		Month:      time.January,
		Weekday:    time.Friday,
		Offset:     2,
		CalcOffset: -1,
		StartYear:  2025,
		Func:       cal.CalcWeekdayOffset,
	}

	// TraditionalReligionsDay represents the annual traditional religions
	// festival (Vodoun Festival): 10-Jan from 1998 (Law 97-031), the second
	// Friday of January from 2025 (Law 2024-32)
	TraditionalReligionsDay = &cal.Holiday{
		Name:      "Fête annuelle des religions traditionnelles",
		Type:      cal.ObservancePublic,
		Month:     time.January,
		Day:       10,
		Weekday:   time.Friday,
		Offset:    2,
		StartYear: 1998,
		Func:      calcTraditionalReligionsDay,
	}

	// RemembranceDay represents Remembrance Day on 16-Jan, a National Day
	// commemorated without a day off
	RemembranceDay = &cal.Holiday{
		Name:  "Journée de Souvenir",
		Type:  cal.ObservanceOther,
		Month: time.January,
		Day:   16,
		Func:  cal.CalcDayOfMonth,
	}

	// PeoplesSovereigntyDay represents People's Sovereignty Day on 28-Feb, a
	// National Day commemorated without a day off
	PeoplesSovereigntyDay = &cal.Holiday{
		Name:  "Journée de la Souveraineté du Peuple",
		Type:  cal.ObservanceOther,
		Month: time.February,
		Day:   28,
		Func:  cal.CalcDayOfMonth,
	}

	// WomensDay represents Women's Day on 8-Mar, a National Day commemorated
	// without a day off
	WomensDay = &cal.Holiday{
		Name:  "Journée de la Femme",
		Type:  cal.ObservanceOther,
		Month: time.March,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// EasterMonday represents Easter Monday on the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Lundi de Pâques", Type: cal.ObservancePublic})

	// LabourDay represents Labour Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "Fête du Travail", Type: cal.ObservancePublic})

	// AscensionDay represents Ascension Day on the 39th day after Easter
	AscensionDay = aa.AscensionDay.Clone(&cal.Holiday{Name: "Jour de l'Ascension", Type: cal.ObservancePublic})

	// PentecostMonday represents Pentecost Monday on the 50th day after Easter
	PentecostMonday = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Lundi de Pentecôte", Type: cal.ObservancePublic})

	// NationalDay represents National Day on 1-Aug, celebrating independence
	NationalDay = &cal.Holiday{
		Name:  "Fête Nationale",
		Type:  cal.ObservancePublic,
		Month: time.August,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// AssumptionDay represents Assumption Day on 15-Aug
	AssumptionDay = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Jour de l'Assomption", Type: cal.ObservancePublic})

	// AllSaintsDay represents All Saints' Day on 1-Nov
	AllSaintsDay = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Jour de la Toussaint", Type: cal.ObservancePublic})

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Jour de la Noël", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays. The National
	// Days (RemembranceDay, PeoplesSovereigntyDay and WomensDay) are not days
	// off and are not included.
	Holidays = []*cal.Holiday{
		NewYear,
		TraditionalReligionsDayEve,
		TraditionalReligionsDay,
		EasterMonday,
		LabourDay,
		AscensionDay,
		PentecostMonday,
		NationalDay,
		AssumptionDay,
		AllSaintsDay,
		ChristmasDay,
	}
)

// calcTraditionalReligionsDay returns 10-Jan until 2024 and the second Friday
// of January from 2025.
func calcTraditionalReligionsDay(h *cal.Holiday, year int) time.Time {
	if year >= 2025 {
		return cal.CalcWeekdayOffset(h, year)
	}
	return cal.CalcDayOfMonth(h, year)
}
