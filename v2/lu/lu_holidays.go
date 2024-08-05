// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package be provides holiday definitions for Luxembourg.
package lu

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NeitJoer represents New Year's Day on 1-Jan
	NeitJoer = aa.NewYear.Clone(&cal.Holiday{Name: "Neit Joer", Type: cal.ObservancePublic})

	// Ouschterméindeg represents Easter Monday on the day after Easter
	Ouschtermeindeg = aa.EasterMonday.Clone(&cal.Holiday{Name: "Ouschterméindeg", Type: cal.ObservancePublic})

	// DagVunAarbecht represents Labor Day on the first Monday in May
	DagVunAarbecht = aa.WorkersDay.Clone(&cal.Holiday{Name: "Dag vun der Aarbecht", Type: cal.ObservancePublic})

	// ChristiHimmelfaart represents Ascension Day on the 39th day after Easter
	ChristiHimmelfaart = aa.AscensionDay.Clone(&cal.Holiday{Name: "Christi Himmelfaart", Type: cal.ObservancePublic})

	// Pengschtméindeg represents Pentecost Monday on the day after Pentecost (50 days after Easter)
	Pengschtméindeg = aa.PentecostMonday.Clone(&cal.Holiday{Name: "Péngschtméindeg", Type: cal.ObservancePublic})

	// Nationalfeierdag represents Luxembourg National Day on 23-Jul
	Nationalfeierdag = &cal.Holiday{
		Name:  "Nationalfeierdag",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   23,
		Func:  cal.CalcDayOfMonth,
	}

	// MariesHimmelfaart represents Assumption of Mary on 15-Aug
	MariesHimmelfaart = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Maries Himmelfaart", Type: cal.ObservancePublic})

	// Allerhellgen represents All Saints' Day on 1-Nov
	Allerhellgen = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Allerhellgen", Type: cal.ObservancePublic})

	// Chrëschtdag represents Christmas Day on 25-Dec
	Chreschtdag = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Chrëschtdag", Type: cal.ObservancePublic})

	// ZweetenDagChrëschtdag represents second Christmas Day Day on 26-Dec
	ZweetenDagChrëschtdag = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Zweeten Dag vum Chrëschtdag", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NeitJoer,
		Ouschtermeindeg,
		DagVunAarbecht,
		ChristiHimmelfaart,
		Pengschtméindeg,
		Nationalfeierdag,
		MariesHimmelfaart,
		Allerhellgen,
		Chreschtdag,
		ZweetenDagChrëschtdag,
	}
)
