// (c) Matías Blasi. Licensed under the BSD license (see LICENSE).

package ar

import (
	"time"

	"github.com/rickar/cal/v2"
)

var (

	// Standard AR weekend substitution rules:
	//   Saturdays move to Friday
	//   Sundays move to Monday
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: -1},
		{Day: time.Sunday, Offset: 1},
	}

	// Argentinian New Year is January 1st
	NewYear = &cal.Holiday{
		Name:  "Año nuevo",
		Type:  cal.ObservancePublic,
		Month: time.January,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// Argentinian Carnival Day 1 Fest is last February Monday
	CarnivalDay1 = &cal.Holiday{
		Name:   "Carnaval Día 1",
		Type:   cal.ObservancePublic,
		Offset: -48,
		Func:   cal.CalcEasterOffset,
	}

	// Argentinian Carnival Day 2 Fest is last February Monday
	CarnivalDay2 = &cal.Holiday{
		Name:   "Carnaval Día 2",
		Type:   cal.ObservancePublic,
		Offset: -47,
		Func:   cal.CalcEasterOffset,
	}

	// Argentinian conmemoration of last coup at 1976
	TruethDay = &cal.Holiday{
		Name:  "Día de la verdad y justicia",
		Type:  cal.ObservancePublic,
		Month: time.March,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// Argentinian conmemoration of Malvinas War Veterans
	MalvinasVeterans = &cal.Holiday{
		Name:  "Día de los Veteranos de la Guerra de Malvinas",
		Type:  cal.ObservancePublic,
		Month: time.April,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// Argentinian Holy Friday
	EasternsDay = &cal.Holiday{
		Name:   "Viernes Santo",
		Type:   cal.ObservancePublic,
		Offset: -3,
		Func:   cal.CalcEasterOffset,
	}

	// Argentinian Labor Day
	LaborDay = &cal.Holiday{
		Name:  "Día del trabajador",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// Argentinian conmemoration of popular revolution of May 1810
	RevolutionDay = &cal.Holiday{
		Name:  "Revolución de Mayo",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// Argentinian commemoration of the passage to the immortality of General Martín Miguel de Güemes.
	GuemesDay = &cal.Holiday{
		Name:     "Aniversario paso a la inmortalidad del General Martín Miguel de Güemes",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
		Month:    time.June,
		Day:      17,
		Func:     cal.CalcDayOfMonth,
	}

	// Argentinian commemoration of the passage to the immortality of General Manuel Belgrano.
	BelgranoDay = &cal.Holiday{
		Name:     "Aniversario paso a la inmortalidad del General Juan Manuel Belgrano",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
		Month:    time.June,
		Day:      20,
		Func:     cal.CalcDayOfMonth,
	}

	// Argentinian Independece Day
	IndependenceDay = &cal.Holiday{
		Name:  "Día de la independencia",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   9,
		Func:  cal.CalcDayOfMonth,
	}

	// Argentinian commemoration of the passage to the immortality of General José de San Martín.
	SanMartinDay = &cal.Holiday{
		Name:     "Aniversario paso a la inmortalidad del General José de San Martín",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
		Month:    time.August,
		Day:      17,
		Func:     cal.CalcDayOfMonth,
	}

	// Argentinian Respect for Cultural Diversity Day
	DiversityDay = &cal.Holiday{
		Name:     "Día del respeto a la diversidad cultural",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
		Month:    time.October,
		Day:      12,
		Func:     cal.CalcDayOfMonth,
	}

	// Argentinian National Sovereignty Day
	SovereigntyDay = &cal.Holiday{
		Name:     "Día de la Soberanía Nacional",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
		Month:    time.November,
		Day:      20,
		Func:     cal.CalcDayOfMonth,
	}

	// Argentinian commemoration of the Immaculate Conception of Mary
	VirgenDay = &cal.Holiday{
		Name:  "Día de la virgen María",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// Argentinian Christmas Day
	ChristmasDay = &cal.Holiday{
		Name:  "Navidad",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		CarnivalDay1,
		CarnivalDay2,
		TruethDay,
		MalvinasVeterans,
		EasternsDay,
		LaborDay,
		RevolutionDay,
		GuemesDay,
		BelgranoDay,
		IndependenceDay,
		SanMartinDay,
		DiversityDay,
		SovereigntyDay,
		VirgenDay,
		ChristmasDay,
	}
)
