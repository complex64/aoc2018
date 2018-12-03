package main

import (
	"fmt"

	"github.com/complex64/aoc2018/internal/pkg/aoc"
)

const sideLength = 1000

type claim struct {
	id                  int
	x, y, width, height int // from the top left
}

func main() {
	m := make(map[uint64]int, sideLength*sideLength) // dovetailed keys from cantor pairing function
	with2OrMore := 0
	err := aoc.CallForEachLineInFile("./input", func(line string) error {
		c, err := readClaim(line)
		if err != nil {
			return err
		}

		for xi := c.x; xi < c.x+c.width; xi++ { // traverse row
			for yi := c.y; yi < c.y+c.height; yi++ { // traverse column
				k := cantorPair(xi, yi)
				m[k]++
				if m[k] == 2 { // all cells with 2 or more count once
					with2OrMore++
				}
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(with2OrMore)
}

func readClaim(s string) (*claim, error) {
	c := &claim{}
	_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &(c.id), &(c.x), &(c.y), &(c.width), &(c.height))
	if err != nil {
		return nil, err
	}
	return c, nil
}

// https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
func cantorPair(x, y int) uint64 {
	k1, k2 := uint64(x), uint64(y)
	return ((k1+k2)*(k1+k2+1))/2 + k2
}
