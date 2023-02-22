// Package hr provides definitions for Croatian holidays
package hr

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Croatia doesn't have substitution rules
	weekendAlt []cal.AltDay

	// NovaGodina represents New Year's Day on 1-Jan
	NovaGodina = &cal.Holiday{
		Name:  "Nova godina",
		Month: time.January,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
		Type:  cal.ObservancePublic,
	}

	// SvetaTriKralja represents Epiphany on 6-Jan
	SvetaTriKralja = &cal.Holiday{
		Name:  "Sveta tri kralja",
		Month: time.January,
		Day:   6,
		Func:  cal.CalcDayOfMonth,
		Type:  cal.ObservancePublic,
	}

	// Uskrs represents Easter
	Uskrs = aa.Easter.Clone(&cal.Holiday{Name: "Uskrs", Type: cal.ObservancePublic})

	// UskrsnjiPonedjeljak represents Easter Monday
	UskrsnjiPonedjeljak = aa.EasterMonday.Clone(&cal.Holiday{Name: "Uskrsni ponedjeljak", Type: cal.ObservancePublic})

	// PraznikRada represents Workers day on 1-May
	PraznikRada = aa.WorkersDay.Clone(&cal.Holiday{Name: "Praznik rada", Type: cal.ObservancePublic})

	// DanDrzavnosti represents National day on 30-May
	DanDrzavnosti = &cal.Holiday{
		Name:  "Dan državnosti",
		Month: time.May,
		Day:   30,
		Func:  cal.CalcDayOfMonth,
		Type:  cal.ObservancePublic,
	}

	// Tijelovo represents Corpus Christi
	Tijelovo = aa.CorpusChristi.Clone(&cal.Holiday{Name: "Tijelovo", Type: cal.ObservancePublic})

	// DanAntifasistickeBorbe represents Anti-Fascist Struggle day on 22-June
	DanAntifasistickeBorbe = &cal.Holiday{
		Name:  "Dan antifašističke borbe",
		Month: time.June,
		Day:   22,
		Func:  cal.CalcDayOfMonth,
		Type:  cal.ObservancePublic,
	}

	// DanPobjedeIDomovinskeZahvalnosti represents Victory and Homeland Thanksgiving day on 5-August
	DanPobjedeIDomovinskeZahvalnosti = &cal.Holiday{
		Name:  "Dan pobjede i domovinske zahvalnosti",
		Month: time.August,
		Day:   5,
		Func:  cal.CalcDayOfMonth,
		Type:  cal.ObservancePublic,
	}

	// VelikaGospa represents Assumption of Mary on 1-August
	VelikaGospa = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Velika Gospa", Type: cal.ObservancePublic})

	// DanSvihSvetih represents AllSaints Day on 1-November
	DanSvihSvetih = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Dan svih svetih", Type: cal.ObservancePublic})

	// DanSjecanjaNaZrtveDomovinskogRata represents Day of Remembrance of the Victims of the Homeland War on 18-November
	DanSjecanjaNaZrtveDomovinskogRata = &cal.Holiday{
		Name:  "Dan sjećanja na žrtve Domovinskog rata",
		Month: time.November,
		Day:   18,
		Func:  cal.CalcDayOfMonth,
		Type:  cal.ObservancePublic,
	}

	// Bozic represents Christmas Day on 25-December
	Bozic = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Božić", Type: cal.ObservancePublic})

	// SvetiStjepan represents Saint Stephen's Day on 25-December
	SvetiStjepan = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Sveti Stjepan", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NovaGodina,
		SvetaTriKralja,
		Uskrs,
		UskrsnjiPonedjeljak,
		PraznikRada,
		DanDrzavnosti,
		Tijelovo,
		DanAntifasistickeBorbe,
		DanPobjedeIDomovinskeZahvalnosti,
		VelikaGospa,
		DanSvihSvetih,
		DanSjecanjaNaZrtveDomovinskogRata,
		Bozic,
		SvetiStjepan,
	}
)
