// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import "time"

var (
	// Holidays in Germany
	DE_Neujahr                = NewYear
	DE_HeiligeDreiKoenige     = NewHoliday(time.January, 6)
	DE_KarFreitag             = GoodFriday
	DE_Ostersonntag           = NewHolidayFunc(calculateOstersonntag)
	DE_Ostermontag            = EasterMonday
	DE_TagderArbeit           = NewHoliday(time.May, 1)
	DE_ChristiHimmelfahrt     = NewHolidayFunc(calculateHimmelfahrt)
	DE_Pfingstsonntag         = NewHolidayFunc(calculatePfingstSonntag)
	DE_Pfingstmontag          = NewHolidayFunc(calculatePfingstMontag)
	DE_Fronleichnam           = NewHolidayFunc(calculateFronleichnam)
	DE_MariaHimmelfahrt       = NewHoliday(time.August, 15)
	DE_TagderDeutschenEinheit = NewHoliday(time.October, 3)
	DE_Reformationstag        = NewHoliday(time.October, 31)
	DE_Allerheiligen          = NewHoliday(time.November, 1)
	DE_BußUndBettag           = NewHolidayFunc(calculateBußUndBettag)
	DE_ErsterWeihnachtstag    = Christmas
	DE_ZweiterWeihnachtstag   = Christmas2
)

// AddGermanHolidays adds all German holidays to the Calendar
func AddGermanHolidays(c *Calendar) {
	c.AddHoliday(DE_Neujahr)
	c.AddHoliday(DE_KarFreitag)
	c.AddHoliday(DE_Ostermontag)
	c.AddHoliday(DE_TagderArbeit)
	c.AddHoliday(DE_ChristiHimmelfahrt)
	c.AddHoliday(DE_Pfingstmontag)
	c.AddHoliday(DE_TagderDeutschenEinheit)
	c.AddHoliday(DE_ErsterWeihnachtstag)
	c.AddHoliday(DE_ZweiterWeihnachtstag)
}

func AddGermanyStateHolidays(c *Calendar, state string) {
	switch state {
	case "BB": // Brandenburg
		c.AddHoliday(DE_Ostersonntag)
		c.AddHoliday(DE_Pfingstsonntag)
		c.AddHoliday(DE_Reformationstag)
	case "BW": // Baden-Württemberg
		c.AddHoliday(DE_HeiligeDreiKoenige)
		c.AddHoliday(DE_Fronleichnam)
		c.AddHoliday(DE_Allerheiligen)
	case "BY": // Bayern
		c.AddHoliday(DE_HeiligeDreiKoenige)
		c.AddHoliday(DE_Fronleichnam)
		c.AddHoliday(DE_MariaHimmelfahrt)
		c.AddHoliday(DE_Allerheiligen)
	case "HE": // Hessen
		c.AddHoliday(DE_Fronleichnam)
	case "MV": // Mecklenburg-Vorpommern
		c.AddHoliday(DE_Reformationstag)
	case "NW": // Nordrhein-Westfalen
		c.AddHoliday(DE_Fronleichnam)
		c.AddHoliday(DE_Allerheiligen)
	case "RP": // Rheinland-Pfalz
		c.AddHoliday(DE_Fronleichnam)
		c.AddHoliday(DE_Allerheiligen)
	case "SA": // Sachsen
		c.AddHoliday(DE_Fronleichnam)
		c.AddHoliday(DE_Reformationstag)
		c.AddHoliday(DE_BußUndBettag)
	case "SL": // Saarland
		c.AddHoliday(DE_Fronleichnam)
		c.AddHoliday(DE_Allerheiligen)
		c.AddHoliday(DE_MariaHimmelfahrt)
	case "ST": // Sachen-Anhalt
		c.AddHoliday(DE_HeiligeDreiKoenige)
		c.AddHoliday(DE_Reformationstag)
	case "TH": // Thüringen
		c.AddHoliday(DE_Fronleichnam)
		c.AddHoliday(DE_Reformationstag)
	}
}

func calculateOstersonntag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	return easter.Month(), easter.Day()
}

func calculateHimmelfahrt(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 39 days after Easter Sunday
	em := easter.AddDate(0, 0, +39)
	return em.Month(), em.Day()
}

func calculatePfingstSonntag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +49)
	return em.Month(), em.Day()
}

func calculatePfingstMontag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +50)
	return em.Month(), em.Day()
}

func calculateFronleichnam(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +60)
	return em.Month(), em.Day()
}

func calculateBußUndBettag(year int, loc *time.Location) (time.Month, int) {
	t := time.Date(year, 11, 23, 0, 0, 0, 0, loc)

	for i := -1; i > -10; i-- {
		d := t.Add(time.Hour * 24 * time.Duration(i))
		if d.Weekday() == time.Wednesday {
			t = d
			break
		}
	}
	return t.Month(), t.Day()
}
