package humanize

import (
	"testing"
	"time"
)

var naturalDeltaTests = []struct {
	// Provide locale for i18n test.
	in     time.Duration
	rel    bool
	out    string
	locale string
}{
	{0, true, "a moment ago", ""},
	{time.Second * 15, false, "14 seconds", ""},
	{time.Hour * -8, true, "8 hours ago", ""},
	{time.Hour * 24 * 18, false, "17 days", ""},
	{time.Hour * 24 * 30 * -4, true, "4 months ago", ""},
	{time.Hour * 24 * 30 * 14, true, "1 year, 1 month from now", ""},
	{time.Hour * 24 * 30 * 14 * -6, true, "6 years ago", ""},
	{time.Hour * 24 * 30 * 14 * 18, false, "2 decades", ""},
	{time.Hour * 18, true, "17个小时后", "zh-cn"},
	{time.Hour * 24 * 30 * -14, true, "il y a 1 année, 1 mois", "fr-fr"},
}

func TestNaturalDelta(t *testing.T) {
	var got string
	now := time.Now()
	for _, tt := range naturalDeltaTests {
		if tt.locale != "" {
			SetLocale(tt.locale)
		}
		v := now.Add(tt.in)
		got = NaturalDelta(v, tt.rel)
		if got != tt.out {
			t.Errorf("NaturalDelta(%v, %v) = %q want %q", tt.in, tt.rel, got, tt.out)
		}
		if tt.locale != "" {
			SetLocale("en-us")
		}
	}
}

var naturalDayTests1 = []struct {
	// Provide locale for i18n test.
	in     time.Duration
	out    string
	locale string
}{
	{0, "today", ""},
	{time.Hour * 24, "tomorrow", ""},
	{time.Hour * -24, "yesterday", ""},
	{time.Hour * 24, "明天", "zh-cn"},
}

var naturalDayTests2 = []struct {
	// Specify absolute datetime as RFC822 layout.
	in     string
	out    string
	locale string
}{
	{"02 Jan 06 15:04 MST", "January 2", ""},
	{"16 Mar 01 00:00 MST", "March 16", "ja-jp"},
}

func TestNaturalDay(t *testing.T) {
	var got string
	now := time.Now()
	for _, tt := range naturalDayTests1 {
		if tt.locale != "" {
			SetLocale(tt.locale)
		}
		v := now.Add(tt.in)
		got = NaturalDay(v)
		if got != tt.out {
			t.Errorf("NaturalDay(%v) = %q want %q", tt.in, got, tt.out)
		}
		if tt.locale != "" {
			SetLocale("en-us")
		}
	}
	for _, tt := range naturalDayTests2 {
		if tt.locale != "" {
			SetLocale(tt.locale)
		}
		v, _ := time.Parse(time.RFC822, tt.in)
		got = NaturalDay(v)
		if got != tt.out {
			t.Errorf("NaturalDay(%v) = %q want %q", tt.in, got, tt.out)
		}
		if tt.locale != "" {
			SetLocale("en-us")
		}
	}
}
