// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package fr provides holiday definitions for France.
package fr

import (
	"time"

	"github.com/Tamh/cal/v2"
	"github.com/Tamh/cal/v2/aa"
)

var (
	// NouvelAn represents New Year's Day on 1-Jan
	NouvelAn = aa.NewYear.Clone(&cal.Holiday{Name: "Nouvel an", Type: cal.ObservancePublic})

	// LundiDePâques represents Easter Monday on the day after Easter
	LundiDePâques = aa.EasterMonday.Clone(&cal.Holiday{Name: "Lundi de Pâques", Type: cal.ObservancePublic})

	// FêteDuTravail represents Labour Day on 1-May
	FêteDuTravail = aa.WorkersDay.Clone(&cal.Holiday{Name: "Fête du Travail", Type: cal.ObservancePublic})

	// FêteDeLaVictoire represents Victory in Europe Day on 8-May
	FêteDeLaVictoire = &cal.Holiday{
		Name:  "Fête de la Victoire",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// Ascension represents Ascension Day on the 39th day after Easter
	Ascension = aa.AscensionDay.Clone(&cal.Holiday{Name: "Ascension", Type: cal.ObservancePublic})

	// LundiDePentecôte represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	LundiDePentecôte = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Lundi de Pentecôte", Type: cal.ObservancePublic})

	// FêteNationale represents Bastille Day on 14-Jul
	FêteNationale = &cal.Holiday{
		Name:  "Fête Nationale",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   14,
		Func:  cal.CalcDayOfMonth,
	}

	// Assomption represents Assumption of Mary on 15-Aug
	Assomption = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Assomption", Type: cal.ObservancePublic})

	// Toussaint represents All Saints' Day on 1-Nov
	Toussaint = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Toussaint", Type: cal.ObservancePublic})

	// Armistice1918 represents Armistice Day on 11-Nov
	Armistice1918 = aa.ArmisticeDay.Clone(&cal.Holiday{Name: "Armistice de 1918", Type: cal.ObservancePublic})

	// Noël represents Christmas Day on 25-Dec
	Noël = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Noël", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NouvelAn,
		LundiDePâques,
		FêteDuTravail,
		FêteDeLaVictoire,
		Ascension,
		LundiDePentecôte,
		FêteNationale,
		Assomption,
		Toussaint,
		Armistice1918,
		Noël,
	}
)
