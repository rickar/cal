// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package it provides holiday definitions for Italy.
package it

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Capodanno represents New Year's Day on 1-Jan
	Capodanno = aa.NewYear.Clone("Capodanno", cal.ObservancePublic, nil)

	// Epifania represents Epipany on 6-Jan
	Epifania = aa.Epiphany.Clone("Epifania", cal.ObservancePublic, nil)

	// Pasquetta represents Easter Monday on the day after Easter
	Pasquetta = aa.EasterMonday.Clone("Pasquetta", cal.ObservancePublic, nil)

	// FestaDellaLiberazione represents Liberation Day on 25-Apr
	FestaDellaLiberazione = &cal.Holiday{
		Name:  "Festa della Liberazione",
		Type:  cal.ObservancePublic,
		Month: time.April,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// FestaDelLavoro represents Labour Day on 1-May
	FestaDelLavoro = aa.WorkersDay.Clone("Festa del Lavoro", cal.ObservancePublic, nil)

	// FestaDellaRepubblica represents Republic Day on 2-Jun
	FestaDellaRepubblica = &cal.Holiday{
		Name:  "Festa della Repubblica",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// Assunzione represents Assumption of Mary on 15-Aug
	Assunzione = aa.AssumptionOfMary.Clone("Assunzione", cal.ObservancePublic, nil)

	// TuttiISanti represents All Saints' Day on 1-Nov
	TuttiISanti = aa.AllSaintsDay.Clone("Tutti i santi", cal.ObservancePublic, nil)

	// Immacolata represents Immaculate Conception on 8-Dec
	Immacolata = aa.ImmaculateConception.Clone("Immacolata Concezione", cal.ObservancePublic, nil)

	// Natale represents Christmas Day on 25-Dec
	Natale = aa.ChristmasDay.Clone("Natale", cal.ObservancePublic, nil)

	// SantoStefano represents Saint Stephen's Day on 26-Dec
	SantoStefano = aa.ChristmasDay2.Clone("Santo Stefano", cal.ObservancePublic, nil)

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Capodanno,
		Epifania,
		Pasquetta,
		FestaDellaLiberazione,
		FestaDelLavoro,
		FestaDellaRepubblica,
		Assunzione,
		TuttiISanti,
		Immacolata,
		Natale,
		SantoStefano,
	}
)
