package nz

import (
	"time"

	"github.com/rickar/cal/v2"
)

var DefaultLoc = time.Local

// CalcMatarikiOffset calculates the occurrence of a holiday that is determined
// by its relation to Matariki, the Māori New Year.
// See https://en.wikipedia.org/wiki/Matariki for more details.
// Dates 2022 to 2052 supported. Data taken from Matariki Advisory Committee report - Final Report V2 - 21 May 2021.
// https://www.mbie.govt.nz/assets/matariki-dates-2022-to-2052-matariki-advisory-group.pdf

var matarikiDates = map[int]time.Time{
	2022: time.Date(2022, time.June, 24, 0, 0, 0, 0, DefaultLoc),
	2023: time.Date(2023, time.July, 14, 0, 0, 0, 0, DefaultLoc),
	2024: time.Date(2024, time.June, 28, 0, 0, 0, 0, DefaultLoc),
	2025: time.Date(2025, time.June, 20, 0, 0, 0, 0, DefaultLoc),
	2026: time.Date(2026, time.July, 10, 0, 0, 0, 0, DefaultLoc),
	2027: time.Date(2027, time.June, 25, 0, 0, 0, 0, DefaultLoc),
	2028: time.Date(2028, time.July, 14, 0, 0, 0, 0, DefaultLoc),
	2029: time.Date(2029, time.July, 6, 0, 0, 0, 0, DefaultLoc),
	2030: time.Date(2030, time.June, 21, 0, 0, 0, 0, DefaultLoc),
	2031: time.Date(2031, time.July, 11, 0, 0, 0, 0, DefaultLoc),
	2032: time.Date(2032, time.July, 2, 0, 0, 0, 0, DefaultLoc),
	2033: time.Date(2033, time.June, 24, 0, 0, 0, 0, DefaultLoc),
	2034: time.Date(2034, time.July, 7, 0, 0, 0, 0, DefaultLoc),
	2035: time.Date(2035, time.June, 29, 0, 0, 0, 0, DefaultLoc),
	2036: time.Date(2036, time.July, 18, 0, 0, 0, 0, DefaultLoc),
	2037: time.Date(2037, time.July, 10, 0, 0, 0, 0, DefaultLoc),
	2038: time.Date(2038, time.June, 25, 0, 0, 0, 0, DefaultLoc),
	2039: time.Date(2039, time.July, 15, 0, 0, 0, 0, DefaultLoc),
	2040: time.Date(2040, time.July, 6, 0, 0, 0, 0, DefaultLoc),
	2041: time.Date(2041, time.July, 19, 0, 0, 0, 0, DefaultLoc),
	2042: time.Date(2042, time.July, 11, 0, 0, 0, 0, DefaultLoc),
	2043: time.Date(2043, time.July, 3, 0, 0, 0, 0, DefaultLoc),
	2044: time.Date(2044, time.June, 24, 0, 0, 0, 0, DefaultLoc),
	2045: time.Date(2045, time.July, 7, 0, 0, 0, 0, DefaultLoc),
	2046: time.Date(2046, time.June, 29, 0, 0, 0, 0, DefaultLoc),
	2047: time.Date(2047, time.July, 19, 0, 0, 0, 0, DefaultLoc),
	2048: time.Date(2048, time.July, 3, 0, 0, 0, 0, DefaultLoc),
	2049: time.Date(2049, time.June, 25, 0, 0, 0, 0, DefaultLoc),
	2050: time.Date(2050, time.July, 15, 0, 0, 0, 0, DefaultLoc),
	2051: time.Date(2051, time.June, 30, 0, 0, 0, 0, DefaultLoc),
	2052: time.Date(2052, time.June, 21, 0, 0, 0, 0, DefaultLoc),
}

func CalcMatarikiOffset(h *cal.Holiday, year int) time.Time {
	tm, ok := matarikiDates[year]
	if !ok {
		// return the first Matariki day
		tm, _ = matarikiDates[2022]
		return tm
	}
	return tm.Add(time.Duration(h.Offset) * 24 * time.Hour)
}
