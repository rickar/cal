package cal

import "time"

// Italian holidays
var (
	ITCapodanno             = NewYear
	ITEpifania              = NewHoliday(time.January, 6)
	ITPasquetta             = EasterMonday
	ITFestaDellaLiberazione = NewHoliday(time.April, 25)
	ITFestaDelLavoro        = NewHoliday(time.May, 1)
	ITFestaDellaRepubblica  = NewHoliday(time.June, 2)
	ITFerragosto            = NewHoliday(time.August, 15)
	ITTuttiISanti           = NewHoliday(time.November, 1)
	ITImmacolata            = NewHoliday(time.December, 8)
	ITNatale                = Christmas
	ITSantoStefano          = Christmas2
)

// AddItalianHolidays adds all Italian holidays to the Calendar
func AddItalianHolidays(c *Calendar) {
	c.AddHoliday(
		ITCapodanno,
		ITEpifania,
		ITPasquetta,
		ITFestaDellaLiberazione,
		ITFestaDelLavoro,
		ITFestaDellaRepubblica,
		ITFerragosto,
		ITTuttiISanti,
		ITImmacolata,
		ITNatale,
		ITSantoStefano,
	)
}
