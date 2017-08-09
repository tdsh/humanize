package humanize

import (
	"bytes"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// Ordinal converts the integer n to the ordinal number as string.
func Ordinal(n int) string {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(n))
	switch n % 10 {
	case 1:
		b.WriteString("st")
	case 2:
		b.WriteString("nd")
	case 3:
		b.WriteString("rd")
	default:
		b.WriteString("th")
	}
	return b.String()
}

// IntComma converts the integer n to a string containing commas
// every three digits. It accepts octal or hex number.
// Returned value is decimal in those cases too.
// If the integer is octal or hex, decimal is also returned.
func IntComma(n int) string {
	orig := strconv.Itoa(n)
	if len(orig) <= 3 {
		return orig
	}
	var res []string
	for i := len(orig); i > 0; i -= 3 {
		if i <= 3 {
			res = append([]string{orig}, res...)
			break
		}
		res = append([]string{orig[i-3 : i]}, res...)
		orig = orig[:i-3]
	}
	return strings.Join(res, ",")
}

// IntWord converts a large integer n to a friendly text representation.
// For example, 1000000 becomes '1.0 million', and '1200000000' becomes
// '1.2 billion'. If n is negative, the number is just returned as string.
// It accepts octal or hex number. Returned value is decimal in those cases
// too.
func IntWord(n int64) string {
	orig := strconv.FormatInt(n, 10)
	if n < 0 {
		return orig
	}
	l := len(orig)
	if l <= 3 {
		return orig
	}
	getLocale()
	f := float64(n)
	f /= math.Pow(10, float64(l-1))
	f = math.Floor(f*10.0+.5) / 10.0
	var ret string
	switch {
	case l <= 6:
		ret = trans("scale_thousand", strconv.FormatFloat(f, 'f', -1, 64))
	case l <= 9:
		ret = trans("scale_million", strconv.FormatFloat(f, 'f', -1, 64))
	case l <= 12:
		ret = trans("scale_billion", strconv.FormatFloat(f, 'f', -1, 64))
	case l <= 15:
		ret = trans("scale_trillion", strconv.FormatFloat(f, 'f', -1, 64))
	case l <= 18:
		ret = trans("scale_quadrillion", strconv.FormatFloat(f, 'f', -1, 64))
	case l <= 21:
		ret = trans("scale_quintillion", strconv.FormatFloat(f, 'f', -1, 64))
	default:
		// never reach here.
		return orig
	}
	return ret
}

// APNumber returns the number spelled out for 1-9. Otherwise, just returns
// the number as string.
func APNumber(n int) string {
	if n < 1 || n > 9 {
		return strconv.Itoa(n)
	}
	getLocale()

	return []string{trans("number_one"), trans("number_two"), trans("number_three"), trans("number_four"), trans("number_five"), trans("number_six"), trans("number_seven"), trans("number_eight"), trans("number_nine")}[n-1]
}

// Fractional converts float64 value f to fractinal number
// as a string, in forms of fractions and mixed fractions.
func Fractional(f float64) string {
	var b bytes.Buffer
	i, frac := math.Modf(f)
	if frac == 0 {
		b.WriteString(strconv.FormatFloat(i, 'f', -1, 64))
		return b.String()
	}
	z := new(big.Rat).SetFloat64(frac)
	var p0, q0, p1, q1 int64 = 0, 1, 1, 0
	var a, q2 int64
	n := z.Num().Int64()
	d := z.Denom().Int64()
	if d <= 1000 {
		if i == 0 {
			b.WriteString(strconv.FormatInt(n, 10))
			b.WriteString("/")
			b.WriteString(strconv.FormatInt(d, 10))
		} else {
			b.WriteString(strconv.FormatFloat(i, 'f', -1, 64))
			b.WriteString(" ")
			b.WriteString(strconv.FormatInt(n, 10))
			b.WriteString("/")
			b.WriteString(strconv.FormatInt(d, 10))
		}
		return b.String()
	}
	// Let's make denominator not too big.
	// This logic to limit denominator has been borrowed from
	// Python fractions module implementation mainly.
	for {
		a = n / d
		q2 = q0 + a*q1
		if q2 > 1000 {
			break
		}
		p0, q0, p1, q1 = p1, q1, p0+a*p1, q2
		n, d = d, n-a*d
	}
	if i == 0 {
		b.WriteString(strconv.FormatInt(p1, 10))
		b.WriteString("/")
		b.WriteString(strconv.FormatInt(q1, 10))
	} else {
		b.WriteString(strconv.FormatFloat(i, 'f', -1, 64))
		b.WriteString(" ")
		b.WriteString(strconv.FormatInt(p1, 10))
		b.WriteString("/")
		b.WriteString(strconv.FormatInt(q1, 10))
	}
	return b.String()
}

// FormatPercent converts float64 value f to a percentage.
// It can accept value more than 1 or negative.
func FormatPercent(f float64) string {
	var b bytes.Buffer
	v := f * 100
	i, f := math.Modf(v)
	if f == 0 {
		b.WriteString(strconv.FormatFloat(i, 'f', -1, 64))
		b.WriteString("%")
	} else {
		// Round as "%.2f"
		f = math.Floor(f*100.0+.5) / 100.0
		b.WriteString(strconv.FormatFloat(i+f, 'f', -1, 64))
		b.WriteString("%")
	}
	return b.String()
}
