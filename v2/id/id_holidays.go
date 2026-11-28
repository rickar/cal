// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package id provides holiday definitions for Indonesia.
//
// National holidays (hari libur nasional) are set each year by a joint decree
// of three ministers (SKB 3 Menteri). Holidays falling on a weekend are not
// moved.
//
// Keputusan Presiden Nomor 8 Tahun 2024 renamed Good Friday and Ascension Day
// from "Isa Almasih" to "Yesus Kristus" and made Easter Sunday a national
// holiday from 2024.
//
// Chinese New Year (Tahun Baru Imlek), Isra Mikraj, Nyepi, Eid al-Fitr, Vesak,
// Eid al-Adha, Islamic New Year and the Prophet's Birthday are national
// holidays too, but are not included: they follow the Chinese, Islamic,
// Balinese Saka and Buddhist calendars, and the government sets their dates
// each year in the SKB. Joint leave days (cuti bersama) are taken from annual
// leave and are not holidays.
package id

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "Tahun Baru Masehi", Type: cal.ObservancePublic})

	// GoodFridayOld represents Good Friday, named Wafat Isa Almasih until 2023
	GoodFridayOld = aa.GoodFriday.Clone(&cal.Holiday{
		Name:    "Wafat Isa Almasih",
		Type:    cal.ObservancePublic,
		EndYear: 2023,
	})

	// GoodFriday represents Good Friday, named Wafat Yesus Kristus from 2024
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{
		Name:      "Wafat Yesus Kristus",
		Type:      cal.ObservancePublic,
		StartYear: 2024,
	})

	// Easter represents Easter Sunday, a national holiday from 2024
	Easter = aa.Easter.Clone(&cal.Holiday{
		Name:      "Kebangkitan Yesus Kristus (Paskah)",
		Type:      cal.ObservancePublic,
		StartYear: 2024,
	})

	// LabourDay represents International Workers' Day on 1-May, a national
	// holiday from 2014 (Keputusan Presiden Nomor 24 Tahun 2013)
	LabourDay = aa.WorkersDay.Clone(&cal.Holiday{
		Name:      "Hari Buruh Internasional",
		Type:      cal.ObservancePublic,
		StartYear: 2014,
	})

	// AscensionDayOld represents Ascension Day, named Kenaikan Isa Almasih
	// until 2023
	AscensionDayOld = aa.AscensionDay.Clone(&cal.Holiday{
		Name:    "Kenaikan Isa Almasih",
		Type:    cal.ObservancePublic,
		EndYear: 2023,
	})

	// AscensionDay represents Ascension Day, named Kenaikan Yesus Kristus from
	// 2024
	AscensionDay = aa.AscensionDay.Clone(&cal.Holiday{
		Name:      "Kenaikan Yesus Kristus",
		Type:      cal.ObservancePublic,
		StartYear: 2024,
	})

	// PancasilaDay represents Pancasila Day on 1-Jun, a national holiday from
	// 2017 (Keputusan Presiden Nomor 24 Tahun 2016)
	PancasilaDay = &cal.Holiday{
		Name:      "Hari Lahir Pancasila",
		Type:      cal.ObservancePublic,
		Month:     time.June,
		Day:       1,
		StartYear: 2017,
		Func:      cal.CalcDayOfMonth,
	}

	// IndependenceDay represents Independence Day on 17-Aug
	IndependenceDay = &cal.Holiday{
		Name:  "Hari Kemerdekaan Republik Indonesia",
		Type:  cal.ObservancePublic,
		Month: time.August,
		Day:   17,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Hari Raya Natal", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		GoodFridayOld,
		GoodFriday,
		Easter,
		LabourDay,
		AscensionDayOld,
		AscensionDay,
		PancasilaDay,
		IndependenceDay,
		ChristmasDay,
	}
)
