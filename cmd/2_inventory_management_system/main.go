package main

import (
	"fmt"

	"github.com/complex64/aoc2018/internal/pkg/aoc"
)

// https://adventofcode.com/2018/day/2
func main() {
	wordsWithDouble, wordsWithTriple := 0, 0
	err := aoc.CallForEachLineInFile("./input", func(line string) error {
		d, t := getDoubleAndTripleCount(line)
		wordsWithDouble += d
		wordsWithTriple += t
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(wordsWithDouble * wordsWithTriple)
}

func getDoubleAndTripleCount(line string) (double int, triple int) {
	// map of (b in [a..z]) -> count(b) in line
	m := make(map[byte]int, 26)
	for i := 0; i < len(line); i++ {
		m[line[i]]++
	}
	for _, n := range m {
		if n == 2 && double == 0 { // doubles count once
			double++
		} else if n == 3 && triple == 0 { // triples count once
			triple++
		}
	}
	return double, triple
}
