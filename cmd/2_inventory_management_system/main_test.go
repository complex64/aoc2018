package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// For example, if you see the following box IDs:
// - abcdef contains no letters that appear exactly two or three times.
// - bababc contains two a and three b, so it counts for both.
// - abbcde contains two b, but no letter appears exactly three times.
// - abcccd contains three c, but no letter appears exactly two times.
// - aabcdd contains two a and two d, but it only counts once.
// - abcdee contains two e.
// - ababab contains three a and three b, but it only counts once.
func TestGetDoubleAndTripleCount(t *testing.T) {
	assertDoubleAndTriple(t, "abcdef", 0, 0)
	assertDoubleAndTriple(t, "bababc", 1, 1)
	assertDoubleAndTriple(t, "abbcde", 1, 0)
	assertDoubleAndTriple(t, "abcccd", 0, 1)
	assertDoubleAndTriple(t, "aabcdd", 1, 0)
	assertDoubleAndTriple(t, "abcdee", 1, 0)
	assertDoubleAndTriple(t, "ababab", 0, 1)
}

func assertDoubleAndTriple(t *testing.T, line string, double, triple int) {
	d, tr := getDoubleAndTripleCount(line)
	assert.Equal(t, double, d)
	assert.Equal(t, triple, tr)
}
