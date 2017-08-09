package humanize

import (
	"testing"
)

var naturalSizeTests = []struct {
	in     int
	gnu    bool
	binary bool
	out    string
}{
	{123, false, false, "123 Bytes"},
	{1978, true, false, "2K"},
	{1000000000, false, true, "953.7 MiB"},
	{742617000027, true, true, "742.6G"},
}

func TestNaturalSize(t *testing.T) {
	var got string
	for _, tt := range naturalSizeTests {
		if tt.gnu == true {
			if tt.binary == true {
				got = NaturalSize(tt.in, GNU(true), Binary(true))
			} else {
				got = NaturalSize(tt.in, GNU(true))
			}
		} else {
			if tt.binary == true {
				got = NaturalSize(tt.in, Binary(true))
			} else {
				got = NaturalSize(tt.in)
			}
		}
		if got != tt.out {
			t.Errorf("NaturalSize(%d, GNU(%v), Binary(%v)) = %q want %q", tt.in, tt.gnu, tt.binary, got, tt.out)
		}
	}
}
