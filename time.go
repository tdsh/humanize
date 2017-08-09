package humanize

import (
	"bytes"
	"math"
	"strconv"
	"time"
)

// NaturalDelta computes the natural representation of the amount
// of time elapsed from now. If relative is true, "ago" or "from now"
// is also returned depending on whether t is past or future.
func NaturalDelta(t time.Time, relative bool) string {
	var b bytes.Buffer
	getLocale()
	delta := time.Since(t).Seconds()
	dSecs := int(math.Abs(delta))
	dDays := dSecs / (60 * 60 * 24)
	dYears := dDays / 365
	dDays = dDays % 365
	dMonths := dDays / 30

	switch {
	case dYears == 0 && dDays == 0:
		switch {
		case dSecs == 0:
			b.WriteString(trans("delta_zero_sec"))
		case dSecs == 1:
			b.WriteString(trans("scale_second", 1))
		case dSecs < 60:
			b.WriteString(trans("scale_second", dSecs))
		case dSecs < 120:
			b.WriteString(trans("scale_minute", 1))
		case dSecs < 3600:
			dMins := dSecs / 60
			b.WriteString(trans("scale_minute", dMins))
		case dSecs < 7200:
			b.WriteString(trans("scale_hour", 1))
		default:
			dHours := dSecs / (60 * 60)
			b.WriteString(trans("scale_hour", dHours))
		}
	case dYears == 0:
		if dDays == 1 {
			b.WriteString(trans("delta_one_day"))
		} else {
			switch {
			case dMonths == 0:
				b.WriteString(trans("scale_day", dDays))
			case dMonths == 1:
				b.WriteString(trans("delta_one_month"))
			default:
				b.WriteString(trans("scale_month", dMonths))
			}
		}
	case dYears == 1:
		switch {
		case dMonths == 0 && dDays == 0:
			b.WriteString(trans("delta_one_year"))
		case dMonths == 0 && dDays == 1:
			b.WriteString(trans("scale_year", 1))
			b.WriteString(", ")
			b.WriteString(trans("scale_day", 1))
		case dMonths == 0 && dDays > 1:
			b.WriteString(trans("scale_year", 1))
			b.WriteString(", ")
			b.WriteString(trans("scale_day", dDays))
		case dMonths == 1:
			b.WriteString(trans("scale_year", 1))
			b.WriteString(", ")
			b.WriteString(trans("scale_month", 1))
		default:
			b.WriteString(trans("scale_year", 1))
			b.WriteString(", ")
			b.WriteString(trans("scale_month", dMonths))
		}
	default:
		decades := dYears / 10
		rem := dYears % 10
		if decades == 0 || rem > 0 {
			b.WriteString(trans("scale_year", dYears))
		} else if decades == 1 {
			b.WriteString(trans("scale_decade", 1))
		} else {
			b.WriteString(trans("scale_decade", decades))
		}
	}
	ret := b.String()

	if relative == true {
		if delta > 0 {
			ret = trans("suffix_before", ret)
		} else if delta < 0 {
			ret = trans("suffix_after", ret)
		}
	}
	return ret
}

// NaturalDay compares t to present day and returns tomorrow, today or yesterday
// if applicable. Otherwise, it returns date as a string.
func NaturalDay(t time.Time) string {
	getLocale()
	now := time.Now()
	if t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day() {
		return trans("day_today")
	}
	yesterday := now.Add(-24 * time.Hour)
	if t.Year() == yesterday.Year() && t.Month() == yesterday.Month() && t.Day() == yesterday.Day() {
		return trans("day_yesterday")
	}
	tomorrow := now.Add(24 * time.Hour)
	if t.Year() == tomorrow.Year() && t.Month() == tomorrow.Month() && t.Day() == tomorrow.Day() {
		return trans("day_tomorrow")
	}
	var b bytes.Buffer
	b.WriteString(t.Month().String())
	b.WriteString(" ")
	b.WriteString(strconv.Itoa(t.Day()))
	return b.String()
}
