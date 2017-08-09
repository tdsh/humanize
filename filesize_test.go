package humanize

import (
	"testing"
)

func TestNaturalSize(t *testing.T) {
	var got string
	got = NaturalSize(123)
	if got != "123 Bytes" {
		t.Fatalf("want %q but %q", "123 Bytes", got)
	}
	got = NaturalSize(1978, GNU(true))
	if got != "2K" {
		t.Fatalf("want %q but %q", "2K", got)
	}
	got = NaturalSize(1000000000, Binary(true))
	if got != "953.7 MiB" {
		t.Fatalf("want %q but %q", "953.7 MiB", got)
	}
	got = NaturalSize(742617000027, GNU(true), Binary(true))
	if got != "742.6G" {
		t.Fatalf("want %q but %q", "742.6G", got)
	}
}
