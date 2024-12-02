// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package cy provides holiday definitions for Cyprus.
package cy

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Protochronia represents New Year's Day on 1-Jan
	Protochronia = aa.NewYear.Clone(&cal.Holiday{Name: "Πρωτοχρονιά", Type: cal.ObservancePublic})

	// Theofania represents Epiphany on 6-Jan
	Theofania = aa.Epiphany.Clone(&cal.Holiday{Name: "Θεοφάνεια", Type: cal.ObservancePublic})

	// KatharaDeftera represents Green Monday (movable, 48 days before Orthodox Easter Sunday)
	KatharaDeftera = &cal.Holiday{
		Name:   "Καθαρά Δευτέρα",
		Type:   cal.ObservancePublic,
		Func:   cal.CalcEasterOffset,
		Julian: true,
		Offset: -48,
	}

	// EllinikiEpanastasi represents Greek Independence Day on 25-Mar
	EllinikiEpanastasi = &cal.Holiday{
		Name:  "Ελληνική Επανάσταση",
		Type:  cal.ObservancePublic,
		Month: time.March,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// EthnikiEpetios represents National Day (EOKA Day) on 1-Apr
	EthnikiEpetios = &cal.Holiday{
		Name:  "Εθνική Επέτειος",
		Type:  cal.ObservancePublic,
		Month: time.April,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// ErgatikoiProtomagia represents Labour Day on 1-May
	ErgatikoiProtomagia = aa.WorkersDay.Clone(&cal.Holiday{Name: "Εργατική Πρωτομαγιά", Type: cal.ObservancePublic})

	// MegaliParaskevi represents Orthodox Good Friday (movable Julian calendar-based)
	MegaliParaskevi = &cal.Holiday{
		Name: "Μεγάλη Παρασκευή",
		Type: cal.ObservancePublic,
		Func: func(h *cal.Holiday, year int) time.Time {
			// Orthodox Good Friday is 2 days before Orthodox Easter Sunday
			return cal.CalcEasterOffset(&cal.Holiday{Offset: -2, Julian: true}, year)
		},
	}

	// DeuteraTouPascha represents Orthodox Easter Monday (Julian calendar-based)
	DeuteraTouPascha = &cal.Holiday{
		Name: "Δευτέρα του Πάσχα",
		Type: cal.ObservancePublic,
		Func: func(h *cal.Holiday, year int) time.Time {
			// Orthodox Easter Monday is 1 day after Orthodox Easter Sunday
			return cal.CalcEasterOffset(&cal.Holiday{Offset: 1, Julian: true}, year)
		},
	}

	// AgiouPnevmatos represents Orthodox Whit Monday (movable, 50 days after Orthodox Easter Sunday)
	AgiouPnevmatos = &cal.Holiday{
		Name:   "Αγίου Πνεύματος",
		Type:   cal.ObservancePublic,
		Func:   cal.CalcEasterOffset,
		Julian: true,
		Offset: 50,
	}

	// KoimisiTisTheotokou represents Assumption Day on 15-Aug
	KoimisiTisTheotokou = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Κοίμηση της Θεοτόκου", Type: cal.ObservancePublic})

	// Anexartisia represents Cyprus Independence Day on 1-Oct
	Anexartisia = &cal.Holiday{
		Name:  "Ανεξαρτησία της Κύπρου",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// EpeteiosTouOchi represents Ochi Day on 28-Oct
	EpeteiosTouOchi = &cal.Holiday{
		Name:  "Επέτειος του Όχι",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   28,
		Func:  cal.CalcDayOfMonth,
	}

	// Christougenna represents Christmas Day on 25-Dec
	Christougenna = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Χριστούγεννα", Type: cal.ObservancePublic})

	// DeyteriMeraTonChristougennon represents Boxing Day (Second Day of Christmas) on 26-Dec
	DeyteriMeraTonChristougennon = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Δεύτερη Μέρα των Χριστουγέννων", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Protochronia,
		Theofania,
		KatharaDeftera,
		EllinikiEpanastasi,
		EthnikiEpetios,
		ErgatikoiProtomagia,
		MegaliParaskevi,
		DeuteraTouPascha,
		AgiouPnevmatos,
		KoimisiTisTheotokou,
		Anexartisia,
		EpeteiosTouOchi,
		Christougenna,
		DeyteriMeraTonChristougennon,
	}
)
