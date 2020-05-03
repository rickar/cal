// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// DefaultLoc is the default time.Location to use in functions that do not
// require a full time.Time value.
var DefaultLoc = time.Local

// Calendar represents a basic yearly calendar with a list of holidays.
type Calendar struct {
	Name        string           // calendar short name
	Description string           // calendar description
	Locations   []*time.Location // locations where the calendar applies
	Holidays    []*Holiday       // applicable holidays for this calendar
}

// IsApplicable reports whether the calendar is applicable for the given
// location.
//
// If no locations have been specified for the calendar, true is returned.
func (c *Calendar) IsApplicable(loc *time.Location) bool {
	if c.Locations == nil {
		return true
	}

	for _, l := range c.Locations {
		if l == loc {
			return true
		}
	}

	return false
}

// AddHoliday adds a holiday to the calendar's list.
func (c *Calendar) AddHoliday(h ...*Holiday) {
	if c.Holidays == nil {
		c.Holidays = make([]*Holiday, 0, 12)
	}

	for _, hd := range h {
		c.Holidays = append(c.Holidays, hd)
	}
}

// IsHoliday reports whether a given date is a holiday or an observation day.
func (c *Calendar) IsHoliday(date time.Time) (actual, observed bool, h *Holiday) {
	if c.Holidays == nil || !c.IsApplicable(date.Location()) {
		return false, false, nil
	}

	for _, hol := range c.Holidays {
		act, obs := hol.Calc(date.Year())
		actMatch := !act.IsZero() && act.Month() == date.Month() && act.Day() == date.Day()
		obsMatch := !obs.IsZero() && obs.Month() == date.Month() && obs.Day() == date.Day()
		if actMatch || obsMatch {
			return actMatch, obsMatch, hol
		}
	}

	return false, false, nil
}
