// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package hu provides holiday definitions for Hungary.
package hu

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Ujev represents New Year's Day on 1-Jan
	Ujev = aa.NewYear.Clone(&cal.Holiday{Name: "Újév", Type: cal.ObservancePublic})

	// NemzetiUnnepMarcius represents Revolution Day on 15-Mar
	NemzetiUnnepMarcius = &cal.Holiday{
		Name:  "Nemzeti ünnep",
		Type:  cal.ObservancePublic,
		Month: time.March,
		Day:   15,
		Func:  cal.CalcDayOfMonth,
	}

	// Nagypentek represents Good Friday (movable, Friday before Easter Sunday)
	Nagypentek = aa.GoodFriday.Clone(&cal.Holiday{Name: "Nagypéntek", Type: cal.ObservancePublic})

	// HusvetHetfo represents Easter Monday (movable, Monday after Easter Sunday)
	HusvetHetfo = aa.EasterMonday.Clone(&cal.Holiday{Name: "Húsvéthétfő", Type: cal.ObservancePublic})

	// AmunkaUnnepe represents Labour Day on 1-May
	AmunkaUnnepe = aa.WorkersDay.Clone(&cal.Holiday{Name: "A munka ünnepe", Type: cal.ObservancePublic})

	// PunkosdHetfo represents Whit Monday (movable, Monday after Pentecost)
	PunkosdHetfo = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Pünkösdhétfő", Type: cal.ObservancePublic})

	// SzentIstvanUnnepe represents State Foundation Day on 20-Aug
	SzentIstvanUnnepe = &cal.Holiday{
		Name:  "Az államalapítás ünnepe",
		Type:  cal.ObservancePublic,
		Month: time.August,
		Day:   20,
		Func:  cal.CalcDayOfMonth,
	}

	// NemzetiUnnepOkt represents Republic Day on 23-Oct
	NemzetiUnnepOkt = &cal.Holiday{
		Name:  "Nemzeti ünnep",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   23,
		Func:  cal.CalcDayOfMonth,
	}

	// Mindenszentek represents All Saints' Day on 1-Nov
	Mindenszentek = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Mindenszentek", Type: cal.ObservancePublic})

	// Karacsony represents Christmas Day on 25-Dec
	Karacsony = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Karácsony", Type: cal.ObservancePublic})

	// KaracsonyMasnapja represents Second Day of Christmas on 26-Dec
	KaracsonyMasnapja = &cal.Holiday{
		Name:  "Karácsony másnapja",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   26,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Ujev,
		NemzetiUnnepMarcius,
		Nagypentek,
		HusvetHetfo,
		AmunkaUnnepe,
		PunkosdHetfo,
		SzentIstvanUnnepe,
		NemzetiUnnepOkt,
		Mindenszentek,
		Karacsony,
		KaracsonyMasnapja,
	}
)
