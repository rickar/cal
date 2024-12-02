// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package mt provides holiday definitions for Malta.
package mt

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// L-ewwelTasSena represents New Year's Day on 1-Jan
	LEwwelTasSena = aa.NewYear.Clone(&cal.Holiday{Name: "L-ewwel tas-Sena", Type: cal.ObservancePublic})

	// NawfragjuSanPawl represents Feast of St. Paul's Shipwreck on 10-Feb
	NawfragjuSanPawl = &cal.Holiday{
		Name:  "Nawfraġju ta' San Pawl",
		Type:  cal.ObservancePublic,
		Month: time.February,
		Day:   10,
		Func:  cal.CalcDayOfMonth,
	}

	// SanGuzepp represents Feast of St. Joseph on 19-Mar
	SanGuzepp = &cal.Holiday{
		Name:  "San Ġużepp",
		Type:  cal.ObservancePublic,
		Month: time.March,
		Day:   19,
		Func:  cal.CalcDayOfMonth,
	}

	// Il-GimghaKbira represents Good Friday (movable)
	IlGimghaLKbira = aa.GoodFriday.Clone(&cal.Holiday{Name: "Il-Ġimgħa l-Kbira", Type: cal.ObservancePublic})

	// JumIl-Ħelsien represents Freedom Day on 31-Mar
	JumIlĦelsien = &cal.Holiday{
		Name:  "Jum il-Ħelsien",
		Type:  cal.ObservancePublic,
		Month: time.March,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	// JumIl-Ħaddiem represents Labour Day on 1-May
	JumIlĦaddiem = aa.WorkersDay.Clone(&cal.Holiday{Name: "Jum il-Ħaddiem", Type: cal.ObservancePublic})

	// SetteGiugno represents Sette Giugno on 7-Jun
	SetteGiugno = &cal.Holiday{
		Name:  "Sette Giugno",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   7,
		Func:  cal.CalcDayOfMonth,
	}

	// L-Imnarja represents Feast of St. Peter and St. Paul on 29-Jun
	LImnarja = &cal.Holiday{
		Name:  "L-Imnarja",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   29,
		Func:  cal.CalcDayOfMonth,
	}

	// SantaMarija represents Feast of the Assumption of Mary on 15-Aug
	SantaMarija = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Santa Marija", Type: cal.ObservancePublic})

	// JumIl-Vitorja represents Victory Day on 8-Sep
	JumIlVitorja = &cal.Holiday{
		Name:  "Jum il-Vitorja",
		Type:  cal.ObservancePublic,
		Month: time.September,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// JumL-Indipendenza represents Independence Day on 21-Sep
	JumLIndipendenza = &cal.Holiday{
		Name:  "Jum l-Indipendenza",
		Type:  cal.ObservancePublic,
		Month: time.September,
		Day:   21,
		Func:  cal.CalcDayOfMonth,
	}

	// Il-Kuncizzjoni represents Feast of the Immaculate Conception on 8-Dec
	IlKuncizzjoni = aa.ImmaculateConception.Clone(&cal.Holiday{Name: "Il-Kunċizzjoni", Type: cal.ObservancePublic})

	// JumIr-Repubblika represents Republic Day on 13-Dec
	JumIrRepubblika = &cal.Holiday{
		Name:  "Jum ir-Repubblika",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   13,
		Func:  cal.CalcDayOfMonth,
	}

	// Il-Milied represents Christmas Day on 25-Dec
	IlMilied = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Il-Milied", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		LEwwelTasSena,
		NawfragjuSanPawl,
		SanGuzepp,
		IlGimghaLKbira,
		JumIlĦelsien,
		JumIlĦaddiem,
		SetteGiugno,
		LImnarja,
		SantaMarija,
		JumIlVitorja,
		JumLIndipendenza,
		IlKuncizzjoni,
		JumIrRepubblika,
		IlMilied,
	}
)
