// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package es provides holiday definitions for Spain.
package es

import (
	"time"

	"github.com/Tamh/cal/v2"
	"github.com/Tamh/cal/v2/aa"
)

var (
	// AñoNuevo represents New Year's Day on 1-Jan
	AñoNuevo = aa.NewYear.Clone(&cal.Holiday{Name: "Año Nuevo", Type: cal.ObservancePublic})

	// Reyes represents Epiphany on 6-Jan
	Reyes = aa.Epiphany.Clone(&cal.Holiday{Name: "Día de Reyes", Type: cal.ObservancePublic})

	// ViernesSanto represents Good Friday on the Friday before Easter
	ViernesSanto = aa.GoodFriday.Clone(&cal.Holiday{Name: "Viernes Santo", Type: cal.ObservancePublic})

	// Trabajador represents Labour Day on 1-May
	Trabajador = aa.WorkersDay.Clone(&cal.Holiday{Name: "Día del Trabajador", Type: cal.ObservancePublic})

	// Asunción represents Assumption of Mary on 15-Aug
	Asunción = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Asunción", Type: cal.ObservancePublic})

	// FiestaNacionalDeEspaña represents Spanish National Day on 12-Oct
	FiestaNacionalDeEspaña = &cal.Holiday{
		Name:  "Fiesta Nacional de España",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   12,
		Func:  cal.CalcDayOfMonth,
	}

	// TodosLosSantos represents All Saints' Day on 1-Nov
	TodosLosSantos = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Día de todos los Santos", Type: cal.ObservancePublic})

	// Constitucion represents Spanish Constitution Day on 6-Dec
	Constitucion = &cal.Holiday{
		Name:  "Día de la Constitución",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   6,
		Func:  cal.CalcDayOfMonth,
	}

	// InmaculadaConcepcion represents Immaculate Conception on 8-Dec
	InmaculadaConcepcion = aa.ImmaculateConception.Clone(&cal.Holiday{Name: "Inmaculada Concepción", Type: cal.ObservancePublic})

	// Navidad represents Christmas Day on 25-Dec
	Navidad = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Navidad", Type: cal.ObservancePublic})

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
