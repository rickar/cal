// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package rs provides holiday definitions for Serbia.
package rs

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NovaGodina represents New Year's Day on 1-Jan
	NovaGodina = aa.NewYear.Clone(&cal.Holiday{Name: "Nova godina", Type: cal.ObservancePublic})

	// DrugiDanNoveGodine represents the second day of New Year's on 2-Jan
	DrugiDanNoveGodine = &cal.Holiday{
		Name:  "Drugi dan Nove godine",
		Type:  cal.ObservancePublic,
		Month: time.January,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// Bozic represents Orthodox Christmas Day on 7-Jan (Julian calendar)
	Bozic = &cal.Holiday{
		Name:   "Božić",
		Type:   cal.ObservancePublic,
		Month:  time.January,
		Day:    7,
		Julian: true,
		Func:   cal.CalcDayOfMonth,
	}

	// DanDrzavnosti represents Statehood Day on 15-Feb
	DanDrzavnosti = &cal.Holiday{
		Name:  "Dan državnosti",
		Type:  cal.ObservancePublic,
		Month: time.February,
		Day:   15,
		Func:  cal.CalcDayOfMonth,
	}

	// DrugiDanDrzavnosti represents the second day of Statehood Day on 16-Feb
	DrugiDanDrzavnosti = &cal.Holiday{
		Name:  "Drugi dan Dana državnosti",
		Type:  cal.ObservancePublic,
		Month: time.February,
		Day:   16,
		Func:  cal.CalcDayOfMonth,
	}

	// PraznikRada represents Labour Day on 1-May
	PraznikRada = aa.WorkersDay.Clone(&cal.Holiday{Name: "Praznik rada", Type: cal.ObservancePublic})

	// DrugiDanPraznikaRada represents the second day of Labour Day on 2-May
	DrugiDanPraznikaRada = &cal.Holiday{
		Name:  "Drugi dan Praznika rada",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// VelikiPetak represents Orthodox Good Friday (movable, 2 days before Orthodox Easter)
	VelikiPetak = &cal.Holiday{
		Name:   "Veliki petak",
		Type:   cal.ObservancePublic,
		Func:   cal.CalcEasterOffset,
		Julian: true,
		Offset: -2,
	}

	// Vaskrs represents Orthodox Easter Sunday (movable)
	Vaskrs = &cal.Holiday{
		Name:   "Vaskrs",
		Type:   cal.ObservancePublic,
		Func:   cal.CalcEasterOffset,
		Julian: true,
	}

	// VaskrsnjiPonedeljak represents Orthodox Easter Monday (movable, 1 day after Orthodox Easter)
	VaskrsnjiPonedeljak = &cal.Holiday{
		Name:   "Vaskršnji ponedeljak",
		Type:   cal.ObservancePublic,
		Func:   cal.CalcEasterOffset,
		Julian: true,
		Offset: 1,
	}

	// DanPrimirja represents Armistice Day on 11-Nov
	DanPrimirja = &cal.Holiday{
		Name:  "Dan primirja u Prvom svetskom ratu",
		Type:  cal.ObservancePublic,
		Month: time.November,
		Day:   11,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NovaGodina,
		DrugiDanNoveGodine,
		Bozic,
		DanDrzavnosti,
		DrugiDanDrzavnosti,
		PraznikRada,
		DrugiDanPraznikaRada,
		VelikiPetak,
		Vaskrs,
		VaskrsnjiPonedeljak,
		DanPrimirja,
	}
)
