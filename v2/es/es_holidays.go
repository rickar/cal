// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package es provides holiday definitions for Spain.
package es

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// AñoNuevo represents New Year's Day on 1-Jan
	AñoNuevo = aa.NewYear.Clone("Año Nuevo", cal.ObservancePublic, nil)

	// Reyes represents Epiphany on 6-Jan
	Reyes = aa.Epiphany.Clone("Día de Reyes", cal.ObservancePublic, nil)

	// ViernesSanto represents Good Friday on the Friday before Easter
	ViernesSanto = aa.GoodFriday.Clone("Viernes Santo", cal.ObservancePublic, nil)

	// Trabajador represents Labour Day on 1-May
	Trabajador = aa.WorkersDay.Clone("Día del Trabajador", cal.ObservancePublic, nil)

	// Asunción represents Assumption of Mary on 15-Aug
	Asunción = aa.AssumptionOfMary.Clone("Asunción", cal.ObservancePublic, nil)

	// FiestaNacionalDeEspaña represents Spanish National Day on 12-Oct
	FiestaNacionalDeEspaña = &cal.Holiday{
		Name:  "Fiesta Nacional de España",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   12,
		Func:  cal.CalcDayOfMonth,
	}

	// TodosLosSantos represents All Saints' Day on 1-Nov
	TodosLosSantos = aa.AllSaintsDay.Clone("Día de todos los Santos", cal.ObservancePublic, nil)

	// Constitucion represents Spanish Constitution Day on 6-Dec
	Constitucion = &cal.Holiday{
		Name:  "Día de la Constitución",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   6,
		Func:  cal.CalcDayOfMonth,
	}

	// InmaculadaConcepcion represents Immaculate Conception on 8-Dec
	InmaculadaConcepcion = aa.ImmaculateConception.Clone("Inmaculada Concepción", cal.ObservancePublic, nil)

	// Navidad represents Christmas Day on 25-Dec
	Navidad = aa.ChristmasDay.Clone("Navidad", cal.ObservancePublic, nil)

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		AñoNuevo,
		Reyes,
		ViernesSanto,
		Trabajador,
		Asunción,
		FiestaNacionalDeEspaña,
		TodosLosSantos,
		Constitucion,
		InmaculadaConcepcion,
		Navidad,
	}
)
