// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package sk provides holiday definitions for Slovakia.
package sk

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (

	// RepublicDay represents Republic Day on 1-Jan
	RepublicDay = aa.NewYear.Clone(&cal.Holiday{Name: "Deň vzniku Slovenskej republiky", Type: cal.ObservancePublic})

	// Epiphany represents Epiphany on 6-Jan
	Epiphany = aa.Epiphany.Clone(&cal.Holiday{Name: "Zjavenie Pána", Type: cal.ObservancePublic})

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Veľkonočný piatok", Type: cal.ObservancePublic})

	// EasterMonday represents Easter Monday - the day after Easter
	EasterMonday = aa.EasterMonday.Clone(&cal.Holiday{Name: "Veľkonočný pondelok", Type: cal.ObservancePublic})

	// LabourDay represents Labour Day on 1-May
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "Sviatok práce", Type: cal.ObservancePublic})

	// Liberation represents Liberation Day on 8-May
	Liberation = &cal.Holiday{
		Name:  "Deň víťazstva nad fašizmom",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// SaintsCyril represents St. Cyril and Methodius Day on 5-Jul
	SaintsCyril = &cal.Holiday{
		Name:  "Sviatok svätého Cyrila a Metoda",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   5,
		Func:  cal.CalcDayOfMonth,
	}

	// SNP represents Slovak National Uprising Anniversary on 29-Aug
	SNP = &cal.Holiday{
		Name:  "Výročie Slovenského národného povstania",
		Type:  cal.ObservancePublic,
		Month: time.August,
		Day:   29,
		Func:  cal.CalcDayOfMonth,
	}

	// Constitution represents Constitution Day on 1-Sep
	Constitution = &cal.Holiday{
		Name:    "Deň Ústavy Slovenskej republiky",
		Type:    cal.ObservancePublic,
		Month:   time.September,
		Day:     1,
		Func:    cal.CalcDayOfMonth,
		EndYear: 2024,
	}

	// LadyOfSorrows represents Day of Our Lady of the Seven Sorrows on 15-Sep
	LadyOfSorrows = &cal.Holiday{
		Name:  "Sviatok Panny Márie Sedembolestnej, patrónky Slovenska",
		Type:  cal.ObservancePublic,
		Month: time.September,
		Day:   15,
		Func:  cal.CalcDayOfMonth,
	}

	// AllSaints represents All Saints' Day on 1-Nov
	AllSaints = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Sviatok všetkých svätých", Type: cal.ObservancePublic})

	// Freedom represents Struggle for Freedom and Democracy Day on 17-Nov
	Freedom = &cal.Holiday{
		Name:  "Deň boja za slobodu a demokraciu",
		Type:  cal.ObservancePublic,
		Month: time.November,
		Day:   17,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasEve represents Christmas Eve on 24-Dec
	ChristmasEve = &cal.Holiday{
		Name:  "Štedrý deň",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Prvý sviatok vianočný", Type: cal.ObservancePublic})

	// SaintStephen represents Boxing Day on 26-Dec
	SaintStephen = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Druhý sviatok vianočný", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		RepublicDay,
		Epiphany,
		GoodFriday,
		EasterMonday,
		LabourDay,
		Liberation,
		SaintsCyril,
		SNP,
		Constitution,
		LadyOfSorrows,
		AllSaints,
		Freedom,
		ChristmasEve,
		ChristmasDay,
		SaintStephen,
	}
)
