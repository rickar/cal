// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ro provides holiday definitions for Romania.
package ro

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// AnulNou represents New Year's Day on 1-Jan
	AnulNou = aa.NewYear.Clone(&cal.Holiday{
		Name: "Anul Nou",
		Type: cal.ObservancePublic,
	})

	// AnulNou2 represents New Year's Second Day on 2-Jan
	AnulNou2 = &cal.Holiday{
		Name:  "A doua zi de Anul Nou",
		Type:  cal.ObservancePublic,
		Month: time.January,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// Boboteaza represents Epiphany on 6-Jan
	Boboteaza = &cal.Holiday{
		Name:      "Boboteaza",
		Type:      cal.ObservancePublic,
		Month:     time.January,
		Day:       6,
		StartYear: 2024,
		Func:      cal.CalcDayOfMonth,
	}

	// SfantulIon represents the celebration of Saint John the Baptist
	SfantulIon = &cal.Holiday{
		Name:      "Sfântul Ion",
		Type:      cal.ObservancePublic,
		Month:     1,
		Day:       7,
		StartYear: 2024,
		Func:      cal.CalcDayOfMonth,
	}

	// ZiuaUniriiPrincipatelorRomane represents the day when, in 1859,  the 2 Romanian principalities, Moldavia and Wallachia, united.
	ZiuaUniriiPrincipatelorRomane = &cal.Holiday{
		Name:  "Ziua Unirii Principatelor Române",
		Type:  cal.ObservancePublic,
		Month: time.January,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// VinereaMare represents Good Friday - two days before Easter
	VinereaMare = &cal.Holiday{
		Name:      "Vinerea Mare",
		Type:      cal.ObservancePublic,
		Offset:    -2,
		Julian:    true,
		Func:      cal.CalcEasterOffset,
		StartYear: 2018,
	}

	// Pastele represents the day of Easter
	Pastele = &cal.Holiday{
		Name:   "Paștele",
		Type:   cal.ObservancePublic,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// ADouaZiDePaste represents Easter Monday on the day after Easter
	ADouaZiDePaste = &cal.Holiday{
		Name:   "A doua zi de Paște",
		Type:   cal.ObservancePublic,
		Offset: 1,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// ZiuaMuncii represents Labour Day on 1-May
	ZiuaMuncii = aa.WorkersDay.Clone(&cal.Holiday{
		Name: "Ziua Muncii",
		Type: cal.ObservancePublic,
	})

	// ZiuaCopilului represents Children's Day on 1-June
	ZiuaCopilului = &cal.Holiday{
		Name:      "Ziua Copilului",
		Type:      cal.ObservancePublic,
		Month:     time.June,
		Day:       1,
		Func:      cal.CalcDayOfMonth,
		StartYear: 2017,
	}

	// Rusalii represents Pentecoast Sunday on the 49th day after Easter
	Rusalii = &cal.Holiday{
		Name:   "Rusalii",
		Type:   cal.ObservancePublic,
		Offset: 49,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// LuniDupaRusalii represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	LuniDupaRusalii = &cal.Holiday{
		Name:   "Luni după Rusalii",
		Type:   cal.ObservancePublic,
		Offset: 50,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// AdormireaMaiciiDomnului represents Assumption of Mary on 15-Aug
	AdormireaMaiciiDomnului = aa.AssumptionOfMary.Clone(&cal.Holiday{
		Name: "Adormirea Maicii Domnului",
		Type: cal.ObservancePublic,
	})

	// SfantulAndrei represents Saint Andrew's Day on 30-Nov
	SfantulAndrei = &cal.Holiday{
		Name:  "Sfântul Andrei",
		Type:  cal.ObservancePublic,
		Month: time.November,
		Day:   30,
		Func:  cal.CalcDayOfMonth,
	}

	// ZiuaNationala represents Great Union Day on 1-Dec
	ZiuaNationala = &cal.Holiday{
		Name:  "Ziua Națională a României",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// Craciunul represents Christmas Day on 25-Dec
	Craciunul = aa.ChristmasDay.Clone(&cal.Holiday{
		Name: "Crăciunul",
		Type: cal.ObservancePublic,
	})

	// Craciunul2 represents the day after Christmas (Boxing Day / St. Stephen's Day) on 26-Dec
	Craciunul2 = aa.ChristmasDay2.Clone(&cal.Holiday{
		Name: "A doua zi de Crăciun",
		Type: cal.ObservancePublic,
	})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		AnulNou,
		AnulNou2,
		Boboteaza,
		SfantulIon,
		ZiuaUniriiPrincipatelorRomane,
		VinereaMare,
		Pastele,
		ADouaZiDePaste,
		ZiuaMuncii,
		ZiuaCopilului,
		Rusalii,
		LuniDupaRusalii,
		AdormireaMaiciiDomnului,
		SfantulAndrei,
		ZiuaNationala,
		Craciunul,
		Craciunul2,
	}
)
