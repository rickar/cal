// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package si provides holiday definitions for Slovenia.
package si

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NovoLeto represents New Year's Day on 1-Jan.
	NovoLeto = aa.NewYear.Clone(&cal.Holiday{Name: "novo leto", Type: cal.ObservancePublic})

	// NovoLeto2 represents New Year's Day on 2-Jan.
	NovoLeto2 = &cal.Holiday{
		Name:  "novo leto",
		Month: time.January,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// PresernovDan represents Prešeren day on 8-Feb.
	PresernovDan = &cal.Holiday{
		Name:  "Prešernov dan",
		Month: time.February,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	// DanUporaProtiOkupatorju represents Day of Uprising Against Occupation on 27-Apr.
	DanUporaProtiOkupatorju = &cal.Holiday{
		Name:  "Dan upora proti okupatorju",
		Month: time.April,
		Day:   27,
		Func:  cal.CalcDayOfMonth,
	}

	// VelikaNoc represents Easter.
	VelikaNoc = &cal.Holiday{
		Name:   "velika noč",
		Offset: 0,
		Func:   cal.CalcEasterOffset,
		Type:   cal.ObservancePublic,
	}

	// VelikonocniPonedeljek represents Easter Monday on the day after Easter.
	VelikonocniPonedeljek = aa.EasterMonday.Clone(&cal.Holiday{Name: "velikonočni ponedeljek", Type: cal.ObservancePublic})

	// BinkostnaNedelja represents Pentecost Sunday (49 days after Easter).
	BinkostnaNedelja = &cal.Holiday{
		Name:   "binkošti",
		Offset: 49,
		Func:   cal.CalcEasterOffset,
	}

	// PrviMaj represents International Workers' Day on 1-May.
	PrviMaj = aa.WorkersDay.Clone(&cal.Holiday{Name: "praznik dela", Type: cal.ObservancePublic})

	// DrugiMaj represents International Workers' Day on 2-May.
	DrugiMaj = &cal.Holiday{
		Name:  "praznik dela",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// DanDrzavnosti represents National Day on 26-Oct.
	DanDrzavnosti = &cal.Holiday{
		Name:  "dan državnosti",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// MarijinoVnebovzetje represents Assumption of Mary on 15-Aug.
	MarijinoVnebovzetje = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Marijino vnebovzetje", Type: cal.ObservancePublic})

	// DanReformacije represents All Saints' Day on 31-Oct.
	DanReformacije = &cal.Holiday{
		Name:  "dan reformacije",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   31,
		Func:  cal.CalcDayOfMonth,
	}

	// DanSpominaNaMrtve represents All Saints' Day on 1-Nov.
	DanSpominaNaMrtve = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "dan spomina na mrtve", Type: cal.ObservancePublic})

	// Bozic represents Christmas Day on 25-Dec.
	Bozic = aa.ChristmasDay.Clone(&cal.Holiday{Name: "božič", Type: cal.ObservancePublic})

	// DanSamostojnostiInEnotnosti represents Independence Day on 26-Dec.
	DanSamostojnostiInEnotnosti = &cal.Holiday{
		Name:  "dan samostojnosti in enotnosti",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   26,
		Func:  cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays.
	Holidays = []*cal.Holiday{
		NovoLeto,
		NovoLeto2,
		PresernovDan,
		DanUporaProtiOkupatorju,
		VelikaNoc,
		VelikonocniPonedeljek,
		BinkostnaNedelja,
		PrviMaj,
		DrugiMaj,
		DanDrzavnosti,
		MarijinoVnebovzetje,
		DanReformacije,
		DanSpominaNaMrtve,
		Bozic,
		DanSamostojnostiInEnotnosti,
	}
)
