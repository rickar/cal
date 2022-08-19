package gr

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// Xristougenna respresents New Year's Day on 1-Jan
	Protoxronia = aa.NewYear.Clone(&cal.Holiday{Name: "Xristougenna", Type: cal.ObservancePublic})

	// Theophania represents Epiphany on 6-Jan
	Theophania = aa.Epiphany.Clone(&cal.Holiday{
		Name: "Θεοφάνεια",
		Type: cal.ObservancePublic,
	})

	// Kathara Deftera represents the first day of the Lent
	KatharaDeftera = &cal.Holiday{
		Name:   "Καθαρά Δευτέρα",
		Type:   cal.ObservancePublic,
		Offset: -48,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// Ikosti Pempti Martiou (Independence Day) is the Anniversary of the declaration of the start of Greek War of Independence from the Ottoman Empire, in 1821.
	IkostiPemptiMartiou = &cal.Holiday{
		Name:  "Εικοστή Πέμπτη Μαρτίου",
		Type:  cal.ObservancePublic,
		Month: time.March,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// Megali Paraskevi represents Good Friday - two days before Easter
	MegaliParaskevi = &cal.Holiday{
		Name:   "Μεγάλη Παρασκευή",
		Type:   cal.ObservancePublic,
		Offset: -2,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// DefteraPascha represents Easter Monday on the day after Easter
	DefteraPascha = &cal.Holiday{
		Name:   "Δευτέρα του Πάσχα",
		Type:   cal.ObservancePublic,
		Offset: 1,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// Ergatiki Protomagia represents Labour Day on 1-May
	ErgatikiProtomagia = aa.WorkersDay.Clone(&cal.Holiday{
		Name: "Εργατική Πρωτομαγιά",
		Type: cal.ObservancePublic,
	})

	// Agiou Prevmatos represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	AgiouPrevmatos = &cal.Holiday{
		Name:   "Αγίου Πνεύματος",
		Type:   cal.ObservancePublic,
		Offset: 50,
		Julian: true,
		Func:   cal.CalcEasterOffset,
	}

	// Kimisi tis Theotokou represents Assumption of Mary on 15-Aug
	KimisiTisTheotokou = aa.AssumptionOfMary.Clone(&cal.Holiday{
		Name: "Κοίμηση της Θεοτόκου",
		Type: cal.ObservancePublic,
	})

	// Imera tou Ochi represents Celebration of the Greek refusal to the Italian ultimatum of 1940.
	ImeraTouOchi = &cal.Holiday{
		Name:  "Ημέρα του Όχι",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   28,
		Func:  cal.CalcDayOfMonth,
	}

	// Christougenna represents Christmas Day on 25-Dec
	Christougenna = aa.ChristmasDay.Clone(&cal.Holiday{
		Name: "Χριστούγεννα",
		Type: cal.ObservancePublic,
	})

	// Sínaxis Yperagías Theotókou Marías respresents the holiday to glorify the Theotokos
	SinaxisYperagiasTheotokou = aa.ChristmasDay2.Clone(&cal.Holiday{
		Name: "Σύναξις Υπεραγίας Θεοτόκου Μαρίας",
		Type: cal.ObservancePublic,
	})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Protoxronia,
		Theophania,
		KatharaDeftera,
		IkostiPemptiMartiou,
		MegaliParaskevi,
		DefteraPascha,
		ErgatikiProtomagia,
		AgiouPrevmatos,
		KimisiTisTheotokou,
		ImeraTouOchi,
		Christougenna,
		SinaxisYperagiasTheotokou,
	}
)
