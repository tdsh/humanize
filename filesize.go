// Package humanize provides functions to convert number to various human
// readable format.
package humanize

import (
	"bytes"
	"math"
	"strconv"
)

// FormatOption is the options NaturalSize can take.
type FormatOption struct {
	binary bool
	gnu    bool
}

// SetOption is method definition to set FormatOption.
type SetOption func(*FormatOption)

// Binary enables binary suffixes (KiB, MiB...) for NaturalSize function.
// It also sets the base 2**10 instead of 10**3.
func Binary(on bool) SetOption {
	return func(ops *FormatOption) {
		ops.binary = on
	}
}

// GNU enables GNU-style prefixes (K, M...) for NaturalSize function.
func GNU(on bool) SetOption {
	return func(ops *FormatOption) {
		ops.gnu = on
	}
}

// decSuf is decimal suffix in default.
var decSuf = []string{" Bytes", " kB", " MB", " GB", " TB", " PB", " EB", " ZB", " YB"}

// binSuf is binary suffix when Binary(true) is passed.
var binSuf = []string{" Bytes", " KiB", " MiB", " GiB", " TiB", " PiB", " EiB", " ZiB", " YiB"}

// gnuSuf is GNU-style suffix when GNU(true) is passed.
var gnuSuf = []string{"B", "K", "M", "G", "T", "P", "E", "Z", "Y"}

// NaturalSize formats val as a number of byteslike a human readable
// filesize (eg. 10 kB).
// Returns decimal suffixes (kB, MB...) by default.
// If option Binary(true) is given, binary suffixes (KiB, MiB...) are used
// and the base will be 2**10 instead of 10**3.
// If option GNU(true) is given, GNU-style prefixes (K, M...) are used.
// GNU had higher priority than Binary so if both options are true, Binary
// is ignored.
//
// Here's some examples to call NaturalSize:
//     NaturalSize(123)
//     NaturalSize(1978, GNU(true))
//     NaturalSize(1000000000, Binary(true))
// Passing both GNU and Binary options is allowed. But in this case GNU is
// given priority and Binary is ignored.
//     NaturalSize(742617000027, GNU(true), Binary(true))
func NaturalSize(val int, options ...SetOption) string {
	var b bytes.Buffer
	opt := FormatOption{}
	for _, o := range options {
		o(&opt)
	}

	var suffix = decSuf
	base := 1000.0
	if opt.gnu == true {
		suffix = gnuSuf
	} else if opt.binary == true && opt.gnu == false {
		base = 1024.0
		suffix = binSuf
	}

	if val == 1 && opt.gnu == false {
		return "1 Byte"
	} else if val == 1 && opt.gnu == true {
		return "1B"
	}

	for i, suf := range suffix {
		d := math.Pow(base, float64(i+1))
		if val < int(d) {
			size := math.Floor(float64(val)*base/d*10.0+.5) / 10.0
			b.WriteString(strconv.FormatFloat(size, 'f', -1, 64))
			b.WriteString(suf)
			break
		}
	}
	return b.String()
}
