// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ie provides holiday definitions for Brazil.
package br

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// AnoNovo represents New Year's Day on 1-Jan
	AnoNovo = aa.NewYear.Clone(&cal.Holiday{Name: "Ano Novo", Type: cal.ObservancePublic})

	// Tiradentes represents Tiradentes' Day on 21-Apr
	Tiradentes = &cal.Holiday{
		Name:  "Tiradentes",
		Month: time.April,
		Day:   21,
		Func:  cal.CalcDayOfMonth,
	}

	// Trabalhador represents Labor Day on 21-Apr
	Trabalhador = aa.WorkersDay.Clone(&cal.Holiday{Name: "Dia do Trabalhador", Type: cal.ObservancePublic})

	// Independencia represents Brazil Independence Day on 07-Sep
	Independencia = &cal.Holiday{
		Name:  "Independência do Brasil",
		Month: time.September,
		Day:   7,
		Func:  cal.CalcDayOfMonth,
	}

	// NossaSenhoraAparecida represents Our Lady of Aparecida Day - Patroness of Brazil on 12-Oct
	NossaSenhoraAparecida = &cal.Holiday{
		Name:  "Nossa Senhora Aparecida",
		Month: time.October,
		Day:   12,
		Func:  cal.CalcDayOfMonth,
	}

	// Finados represents Day of the Dead on 02-Nov
	Finados = &cal.Holiday{
		Name:  "Finados",
		Month: time.November,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// Republica represents Proclamation of the Republic on 15-Nov
	Republica = &cal.Holiday{
		Name:  "Proclamação da República",
		Month: time.November,
		Day:   15,
		Func:  cal.CalcDayOfMonth,
	}

	// CorpusChristi represents Corpus Christi on the 60th day after Easter
	CorpusChristi = aa.CorpusChristi.Clone(&cal.Holiday{Name: "Corpus Christi", Type: cal.ObservancePublic})

	// SextaFeiraSanta represents Good Friday - two days before Easter
	SextaFeiraSanta = aa.GoodFriday.Clone(&cal.Holiday{Name: "Sexta-feira Santa", Type: cal.ObservancePublic})

	// Carnaval represents Brazilian Carnival - 47 days before Easter
	Carnaval = &cal.Holiday{
		Name:   "Carnaval",
		Type:   cal.ObservancePublic,
		Offset: -47,
		Func:   cal.CalcEasterOffset,
	}

	// Natal represents Christmas Day on 25-Dec
	Natal = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Natal", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		AnoNovo,
		Tiradentes,
		Trabalhador,
		Independencia,
		NossaSenhoraAparecida,
		Finados,
		Republica,
		CorpusChristi,
		SextaFeiraSanta,
		Carnaval,
		Natal,
	}
)
