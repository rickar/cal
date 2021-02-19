// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).
// (c) Juan P. Garcia

// Package co provides holiday definitions for Colombia.
package co

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// AñoNuevo represents New Year's Day on 1-Jan
	AñoNuevo = aa.NewYear.Clone(&cal.Holiday{Name: "Año Nuevo", Type: cal.ObservancePublic})

	// Reyes represents Epiphany on 6-Jan moved to next Monday
	Reyes = &cal.Holiday{
		Name:    "Día de Reyes",
		Type:    cal.ObservancePublic,
		Month:   aa.Epiphany.Month,
		Day:     aa.Epiphany.Day,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayFrom,
	}

	// DomingoDeRamos represents Palm Sunday on the Sunday before Easter
	DomingoDeRamos = &cal.Holiday{
		Name:   "Domingo de Ramos",
		Offset: -7,
		Func:   cal.CalcEasterOffset,
		Type:   cal.ObservanceOther,
	}

	// JuevesSanto represents Maundy Thursday on the Thursday before Easter
	JuevesSanto = aa.MaundyThursday.Clone(&cal.Holiday{Name: "Jueves Santo", Type: cal.ObservancePublic})

	// ViernesSanto represents Good Friday on the Friday before Easter
	ViernesSanto = aa.GoodFriday.Clone(&cal.Holiday{Name: "Viernes Santo", Type: cal.ObservancePublic})

	// Pascua represents Easter
	Pascua = &cal.Holiday{
		Name:   "Pascua",
		Offset: 0,
		Func:   cal.CalcEasterOffset,
		Type:   cal.ObservanceOther,
	}

	// DíaAscension represents Day of Ascention on the Monday 43 days after Easter
	DíaAscension = &cal.Holiday{
		Name:   "Día de la Ascensión",
		Offset: 43,
		Func:   cal.CalcEasterOffset,
		Type:   cal.ObservancePublic,
	}

	// CorpusChristi represents Corpus Christi on the Monday 64 days after Easter
	CorpusChristi = &cal.Holiday{
		Name:   "Corpus Christi",
		Offset: 64,
		Func:   cal.CalcEasterOffset,
		Type:   cal.ObservancePublic,
	}

	// SagradoCorazon represents Sacred Heart on the Monday 71 days after Easter
	SagradoCorazon = &cal.Holiday{
		Name:   "Sagrado Corazón",
		Offset: 71,
		Func:   cal.CalcEasterOffset,
		Type:   cal.ObservancePublic,
	}

	// DíaMujer represents Women's Day on 8-Mar
	DíaMujer = &cal.Holiday{
		Name:  "Día de la Mujer",
		Type:  cal.ObservanceOther,
		Month: time.March,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// DíaSanJose represents Saint Joseph's Day on 19-Mar moved to next Monday
	DíaSanJose = &cal.Holiday{
		Name:    "Día de San José",
		Type:    cal.ObservancePublic,
		Month:   time.March,
		Day:     19,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayFrom,
	}

	// DíaIdioma represents Day of Language on 23-Mar
	DíaIdioma = &cal.Holiday{
		Name:  "Día del Idioma",
		Type:  cal.ObservanceOther,
		Month: time.April,
		Day:   23,
		Func:  cal.CalcDayOfMonth,
	}

	// DíaNino represents Children's Day on the last Saturday of April
	DíaNino = &cal.Holiday{
		Name:    "Día del Niño",
		Type:    cal.ObservanceOther,
		Month:   time.April,
		Weekday: time.Saturday,
		Offset:  -1,
		Func:    cal.CalcWeekdayOffset,
	}

	// Trabajo represents Labour Day on 1-May
	Trabajo = aa.WorkersDay.Clone(&cal.Holiday{Name: "Día del Trabajo", Type: cal.ObservancePublic})

	// DíaMadre represents Mother's Day on the second Sunday of May
	DíaMadre = &cal.Holiday{
		Name:    "Día de la Madre",
		Type:    cal.ObservanceOther,
		Month:   time.May,
		Weekday: time.Sunday,
		Offset:  2,
		Func:    cal.CalcWeekdayOffset,
	}

	// DíaPadre represents Father's Day on the third Sunday of June
	DíaPadre = &cal.Holiday{
		Name:    "Día del Padre",
		Type:    cal.ObservanceOther,
		Month:   time.June,
		Weekday: time.Sunday,
		Offset:  3,
		Func:    cal.CalcWeekdayOffset,
		Except:  []int{2018},
	}

	// DíaPadre2018 represents Father's Day on the fourth Sunday of June 2018
	DíaPadre2018 = &cal.Holiday{
		Name:      "Día del Padre",
		Type:      cal.ObservanceOther,
		Month:     time.June,
		Weekday:   time.Sunday,
		Offset:    4,
		Func:      cal.CalcWeekdayOffset,
		StartYear: 2018,
		EndYear:   2018,
	}

	// SanPedroSanPablo represents the Feast of Saint Peter and Saint Paul on 29-Jul moved to next Monday
	SanPedroSanPablo = &cal.Holiday{
		Name:    "San Pedro y San Pablo",
		Type:    cal.ObservancePublic,
		Month:   time.June,
		Day:     29,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayFrom,
	}

	// Independencia represents Independence Day on 20-Jul
	Independencia = &cal.Holiday{
		Name:  "Día de la Independencia",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   20,
		Func:  cal.CalcDayOfMonth,
	}

	// BatallaBoyaca represents Battle of Boyaca Day on 7-Aug
	BatallaBoyaca = &cal.Holiday{
		Name:  "Batalla de Boyacá",
		Type:  cal.ObservancePublic,
		Month: time.August,
		Day:   7,
		Func:  cal.CalcDayOfMonth,
	}

	// Asunción represents Assumption of Mary on 15-Aug moved to next Monday
	Asunción = &cal.Holiday{
		Name:    "Asunción de la Virgen",
		Type:    cal.ObservancePublic,
		Month:   aa.AssumptionOfMary.Month,
		Day:     aa.AssumptionOfMary.Day,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayFrom,
	}

	// DíaAmorAmistad represents Valentine's Day on the third Saturday of September
	DíaAmorAmistad = &cal.Holiday{
		Name:    "Día del Amor y la Amistad",
		Type:    cal.ObservanceOther,
		Month:   time.September,
		Weekday: time.Saturday,
		Offset:  3,
		Func:    cal.CalcWeekdayOffset,
	}

	// DíaRaza represents Columbus' Day on 15-Aug moved to next Monday
	DíaRaza = &cal.Holiday{
		Name:    "Día de la Raza",
		Month:   time.October,
		Day:     12,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayFrom,
	}

	// TodosLosSantos represents All Saints' Day on 1-Nov moved to next Monday
	TodosLosSantos = &cal.Holiday{
		Name:    "Día de todos los Santos",
		Type:    cal.ObservancePublic,
		Month:   aa.AllSaintsDay.Month,
		Day:     aa.AllSaintsDay.Day,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayFrom,
	}

	// IndependenciaCartagena represents Cartagena's Independence Day on 11-Nov moved to next Monday
	IndependenciaCartagena = &cal.Holiday{
		Name:    "Independencia de Cartagena",
		Type:    cal.ObservancePublic,
		Month:   time.November,
		Day:     11,
		Weekday: time.Monday,
		Offset:  1,
		Func:    cal.CalcWeekdayFrom,
	}

	// DíaMujerColombiana represents Colombian Women's Day on 14-Nov
	DíaMujerColombiana = &cal.Holiday{
		Name:  "Día de la Mujer Colombiana",
		Type:  cal.ObservanceOther,
		Month: time.November,
		Day:   14,
		Func:  cal.CalcDayOfMonth,
	}

	// VísperaInmaculadaConcepción represents Eve of Immaculate Conception on 7-Dec
	VísperaInmaculadaConcepción = &cal.Holiday{
		Name:  "Víspera de la Inmaculada Concepción",
		Type:  cal.ObservanceOther,
		Month: time.December,
		Day:   7,
		Func:  cal.CalcDayOfMonth,
	}

	// InmaculadaConcepción represents Immaculate Conception on 8-Dec
	InmaculadaConcepción = aa.ImmaculateConception.Clone(&cal.Holiday{Name: "Inmaculada Concepción", Type: cal.ObservancePublic})

	// Nochebuena represents Christmas' Eve on 24-Dec
	Nochebuena = &cal.Holiday{
		Name:  "Nochebuena",
		Type:  cal.ObservanceOther,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// Navidad represents Christmas Day on 25-Dec
	Navidad = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Navidad", Type: cal.ObservancePublic})

	// AñoViejo represents New Year's Eve on 31-Dec
	AñoViejo = &cal.Holiday{
		Name:  "Año Viejo",
		Type:  cal.ObservanceOther,
		Month: time.December,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		AñoNuevo,
		Reyes,
		DomingoDeRamos,
		JuevesSanto,
		ViernesSanto,
		Pascua,
		DíaAscension,
		CorpusChristi,
		SagradoCorazon,
		DíaMujer,
		DíaSanJose,
		DíaIdioma,
		DíaNino,
		Trabajo,
		DíaMadre,
		DíaPadre,
		DíaPadre2018,
		SanPedroSanPablo,
		Independencia,
		BatallaBoyaca,
		Asunción,
		DíaAmorAmistad,
		DíaRaza,
		TodosLosSantos,
		IndependenciaCartagena,
		DíaMujerColombiana,
		VísperaInmaculadaConcepción,
		InmaculadaConcepción,
		Nochebuena,
		Navidad,
		AñoViejo,
	}
)
