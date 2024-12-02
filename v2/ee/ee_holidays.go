// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ee provides holiday definitions for Estonia.
package ee

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Uusaasta represents New Year's Day on 1-Jan
	Uusaasta = aa.NewYear.Clone(&cal.Holiday{Name: "Uusaasta", Type: cal.ObservancePublic})

	// Iseseisvuspaev represents Independence Day on 24-Feb
	Iseseisvuspaev = &cal.Holiday{
		Name:  "Iseseisvuspäev",
		Type:  cal.ObservancePublic,
		Month: time.February,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// SuurReede represents Good Friday (movable, Friday before Easter Sunday)
	SuurReede = aa.GoodFriday.Clone(&cal.Holiday{Name: "Suur Reede", Type: cal.ObservancePublic})

	// Ulestousmispuhade represents Easter Sunday (movable)
	Ulestousmispuhade = aa.Easter.Clone(&cal.Holiday{Name: "Ülestõusmispühade 1. püha", Type: cal.ObservancePublic})

	// Kevadpuha represents Spring Day on 1-May
	Kevadpuha = aa.WorkersDay.Clone(&cal.Holiday{Name: "Kevadpüha", Type: cal.ObservancePublic})

	// Nelipuha represents Pentecost (movable, 49 days after Easter Sunday)
	Nelipuha = aa.Pentecost.Clone(&cal.Holiday{Name: "Nelipühade 1. püha", Type: cal.ObservancePublic})

	// Voidupuha represents Victory Day on 23-Jun
	Voidupuha = &cal.Holiday{
		Name:  "Võidupüha",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   23,
		Func:  cal.CalcDayOfMonth,
	}

	// Jaanipaev represents Midsummer Day on 24-Jun
	Jaanipaev = &cal.Holiday{
		Name:  "Jaanipäev",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// Taasiseseisvuspaev represents Day of Restoration of Independence on 20-Aug
	Taasiseseisvuspaev = &cal.Holiday{
		Name:  "Taasiseseisvumispäev",
		Type:  cal.ObservancePublic,
		Month: time.August,
		Day:   20,
		Func:  cal.CalcDayOfMonth,
	}

	// Joululaupaev represents Christmas Eve on 24-Dec
	Joululaupaev = &cal.Holiday{
		Name:  "Jõululaupäev",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   24,
		Func:  cal.CalcDayOfMonth,
	}

	// EsimeneJoulupuha represents Christmas Day on 25-Dec
	EsimeneJoulupuha = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Esimene jõulupüha", Type: cal.ObservancePublic})

	// TeineJoulupuha represents Boxing Day (Second Day of Christmas) on 26-Dec
	TeineJoulupuha = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Teine jõulupüha", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Uusaasta,
		Iseseisvuspaev,
		SuurReede,
		Ulestousmispuhade,
		Kevadpuha,
		Nelipuha,
		Voidupuha,
		Jaanipaev,
		Taasiseseisvuspaev,
		Joululaupaev,
		EsimeneJoulupuha,
		TeineJoulupuha,
	}
)
