// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package pt provides holiday definitions for Portugal.
package pt

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// AnoNovo represents New Year's Day on 1-Jan
	AnoNovo = aa.NewYear.Clone(&cal.Holiday{Name: "Ano Novo", Type: cal.ObservancePublic})

	// SextaFeiraSanta represents Good Friday on the Friday before Easter (movable)
	SextaFeiraSanta = aa.GoodFriday.Clone(&cal.Holiday{Name: "Sexta-feira Santa", Type: cal.ObservancePublic})

	// DomingoPascoa represents Easter Sunday (movable)
	DomingoPascoa = aa.Easter.Clone(&cal.Holiday{Name: "Domingo de Páscoa", Type: cal.ObservancePublic})

	// DiaDaLiberdade represents Freedom Day on 25-Apr
	DiaDaLiberdade = &cal.Holiday{
		Name:  "Dia da Liberdade",
		Type:  cal.ObservancePublic,
		Month: time.April,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// DiaDoTrabalhador represents Labour Day on 1-May
	DiaDoTrabalhador = aa.WorkersDay.Clone(&cal.Holiday{Name: "Dia do Trabalhador", Type: cal.ObservancePublic})

	// CorpoDeDeus represents Corpus Christi, 60 days after Easter (movable)
	CorpoDeDeus = &cal.Holiday{
		Name:   "Corpo de Deus",
		Type:   cal.ObservancePublic,
		Func:   cal.CalcEasterOffset,
		Offset: 60,
	}

	// DiaDePortugal represents Portugal Day on 10-Jun
	DiaDePortugal = &cal.Holiday{
		Name:  "Dia de Portugal",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   10,
		Func:  cal.CalcDayOfMonth,
	}

	// AssuncaoDeNossaSenhora represents Assumption Day on 15-Aug
	AssuncaoDeNossaSenhora = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Assunção de Nossa Senhora", Type: cal.ObservancePublic})

	// ImplantacaoDaRepublica represents Republic Day on 5-Oct
	ImplantacaoDaRepublica = &cal.Holiday{
		Name:  "Implantação da República",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   5,
		Func:  cal.CalcDayOfMonth,
	}

	// TodosOsSantos represents All Saints' Day on 1-Nov
	TodosOsSantos = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Dia de Todos os Santos", Type: cal.ObservancePublic})

	// RestauracaoDaIndependencia represents Restoration of Independence Day on 1-Dec
	RestauracaoDaIndependencia = &cal.Holiday{
		Name:  "Restauração da Independência",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// ImaculadaConceicao represents Immaculate Conception on 8-Dec
	ImaculadaConceicao = aa.ImmaculateConception.Clone(&cal.Holiday{Name: "Imaculada Conceição", Type: cal.ObservancePublic})

	// Natal represents Christmas Day on 25-Dec
	Natal = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Natal", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		AnoNovo,
		SextaFeiraSanta,
		DomingoPascoa,
		DiaDaLiberdade,
		DiaDoTrabalhador,
		CorpoDeDeus,
		DiaDePortugal,
		AssuncaoDeNossaSenhora,
		ImplantacaoDaRepublica,
		TodosOsSantos,
		RestauracaoDaIndependencia,
		ImaculadaConceicao,
		Natal,
	}
)

