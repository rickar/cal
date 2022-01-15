// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// DefaultLoc is the default time.Location to use in functions that do not
// require a full time.Time value.
var DefaultLoc = time.Local

// CacheMaxSize is the maximum number of items that can be stored in the cache
var CacheMaxSize = 365 * 3

// CacheEvictSize is the number of items to evict from cache when it is full
var CacheEvictSize = 30

// Calendar represents a basic yearly calendar with a list of holidays.
type Calendar struct {
	Name        string           // calendar short name
	Description string           // calendar description
	Locations   []*time.Location // locations where the calendar applies
	Holidays    []*Holiday       // applicable holidays for this calendar
	Cacheable   bool             // indicates that holiday calcs can be cached (don't change holiday defs while enabled)

	isHolCache map[holCacheKey]*holCacheEntry // cached results for IsHoliday
}

type holCacheKey struct {
	year  int
	month time.Month
	day   int
}

type holCacheEntry struct {
	act bool
	obs bool
	hol *Holiday
}

// shared entry to avoid repeated allocations for "false" entries
var holFalseEntry *holCacheEntry = &holCacheEntry{act: false, obs: false, hol: nil}

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

	c.Holidays = append(c.Holidays, h...)
}

// IsHoliday reports whether a given date is a holiday or an observation day.
func (c *Calendar) IsHoliday(date time.Time) (actual, observed bool, h *Holiday) {
	if c.Holidays == nil || !c.IsApplicable(date.Location()) {
		return false, false, nil
	}

	year, month, day := date.Date()

	if c.Cacheable {
		if c.isHolCache == nil {
			c.isHolCache = make(map[holCacheKey]*holCacheEntry)
		}
		if v, ok := c.isHolCache[holCacheKey{year: year, month: month, day: day}]; ok {
			return v.act, v.obs, v.hol
		}
	}

	for _, hol := range c.Holidays {
		act, obs := hol.Calc(year)

		actMatch := !act.IsZero()
		if actMatch {
			_, actMonth, actDay := act.Date()
			actMatch = actMonth == month && actDay == day
		}
		obsMatch := !obs.IsZero()
		if obsMatch {
			_, obsMonth, obsDay := obs.Date()
			obsMatch = obsMonth == month && obsDay == day
		}
		if actMatch || obsMatch {
			if c.Cacheable {
				c.evict()
				c.isHolCache[holCacheKey{year: year, month: month, day: day}] =
					&holCacheEntry{act: actMatch, obs: obsMatch, hol: hol}
			}
			return actMatch, obsMatch, hol
		}

		// handle holidays that wrap around to the next or previous year
		// e.g., New Year's Day on Saturday 1 Jan observed on Friday 31 Dec
		if hol.Observed != nil {
			actMonth := time.Month(0)
			if !act.IsZero() {
				actMonth = act.Month()
			}
			if actMonth == time.January {
				_, obs = hol.Calc(year + 1)
				obsMatch := !obs.IsZero()
				if obsMatch {
					obsYear, obsMonth, obsDay := obs.Date()
					obsMatch = obsYear == year && obsMonth == month && obsDay == day
				}
				if obsMatch {
					if c.Cacheable {
						c.evict()
						c.isHolCache[holCacheKey{year: year, month: month, day: day}] =
							&holCacheEntry{act: false, obs: obsMatch, hol: hol}
					}
					return false, obsMatch, hol
				}
			} else if actMonth == time.December {
				_, obs = hol.Calc(year - 1)
				obsMatch := !obs.IsZero()
				if obsMatch {
					obsYear, obsMonth, obsDay := obs.Date()
					obsMatch = obsYear == year && obsMonth == month && obsDay == day
				}
				if obsMatch {
					if c.Cacheable {
						c.evict()
						c.isHolCache[holCacheKey{year: year, month: month, day: day}] =
							&holCacheEntry{act: false, obs: obsMatch, hol: hol}
					}
					return false, obsMatch, hol
				}
			}
		}
	}

	if c.Cacheable {
		c.evict()
		c.isHolCache[holCacheKey{year: year, month: month, day: day}] = holFalseEntry
	}
	return false, false, nil
}

func (c *Calendar) evict() {
	if len(c.isHolCache) >= CacheMaxSize {
		n := 0
		for k := range c.isHolCache {
			delete(c.isHolCache, k)
			n++
			if n >= CacheEvictSize {
				break
			}
		}
	}
}
