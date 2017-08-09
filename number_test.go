package humanize

import (
	"testing"
)

var ordinalTests = []struct {
	in  int
	out string
}{
	{2, "2nd"},
	{21, "21st"},
	{32, "32nd"},
	{103, "103rd"},
}

func TestOrdinal(t *testing.T) {
	var got string
	for _, tt := range ordinalTests {
		got = Ordinal(tt.in)
		if got != tt.out {
			t.Errorf("Ordinal(%v) = %q want %q", tt.in, got, tt.out)
		}
	}
}

var intCommaTests = []struct {
	in  int
	out string
}{
	{123, "123"},
	{12345, "12,345"},
	{1234567890, "1,234,567,890"},
	{0123, "83"},
	{0xdeadbeaf, "3,735,928,495"},
}

func TestIntComma(t *testing.T) {
	var got string
	for _, tt := range intCommaTests {
		got = IntComma(tt.in)
		if got != tt.out {
			t.Errorf("IntComma(%v) = %q want %q", tt.in, got, tt.out)
		}
	}
}

var intWordTests = []struct {
	// Provide locale for i18n test.
	in     int64
	out    string
	locale string
}{
	{1200, "1.2 thousand", ""},
	{-1192, "-1192", ""},
	{1200000000, "1.2 billion", ""},
	{9990000, "10 million", ""},
	{9223372036854775807, "9.2 quintillion", ""},
	{123456789, "1.2 百万", "ja-jp"},
	{77777777777777, "7.8 trillones", "es-es"},
}

func TestIntWord(t *testing.T) {
	var got string
	for _, tt := range intWordTests {
		if tt.locale != "" {
			SetLocale(tt.locale)
		}
		got = IntWord(tt.in)
		if got != tt.out {
			t.Errorf("IntWord(%v) = %q want %q", tt.in, got, tt.out)
		}
		if tt.locale != "" {
			SetLocale("en-us")
		}
	}
}

var apNumberTests = []struct {
	// Provide locale for i18n test.
	in     int
	out    string
	locale string
}{
	{5, "five", ""},
	{100, "100", ""},
	{3, "trois", "fr-fr"},
}

func TestAPNumber(t *testing.T) {
	var got string
	for _, tt := range apNumberTests {
		if tt.locale != "" {
			SetLocale(tt.locale)
		}
		got = APNumber(tt.in)
		if got != tt.out {
			t.Errorf("APNumber(%v) = %q want %q", tt.in, got, tt.out)
		}
		if tt.locale != "" {
			SetLocale("en-us")
		}
	}
}

var fractionalTests = []struct {
	in  float64
	out string
}{
	{0.3, "3/10"},
	{1.3, "1 3/10"},
	{1, "1"},
	{0.2, "1/5"},
	{0.6789843274, "294/433"},
	{0.54321, "44/81"},
}

func TestFractional(t *testing.T) {
	var got string
	for _, tt := range fractionalTests {
		got = Fractional(tt.in)
		if got != tt.out {
			t.Errorf("Fractional(%v) = %q want %q", tt.in, got, tt.out)
		}
	}
}

var formatPercentTests = []struct {
	in  float64
	out string
}{
	{0.3, "30%"},
	{0.87654, "87.65%"},
	{5.0012345, "500.12%"},
	{-0.45, "-45%"},
}

func TestFormatPercent(t *testing.T) {
	var got string
	for _, tt := range formatPercentTests {
		got = FormatPercent(tt.in)
		if got != tt.out {
			t.Errorf("FormatPercent(%v) = %q want %q", tt.in, got, tt.out)
		}
	}
}
