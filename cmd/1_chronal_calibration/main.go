package main

import (
	"fmt"

	"github.com/complex64/aoc2018/internal/pkg/aoc"
)

func main() {
	var sum, i int32
	err := aoc.CallForEachLineInFile("./input", func(line string) error {
		// each line holds a single sign-prefixed integer
		if _, err := fmt.Sscanf(line, "%d", &i); err != nil {
			return err
		}
		sum += i
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(sum)
}
