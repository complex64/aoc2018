package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f is a file of integers to sum up
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// tokenize the buffered input line by line
	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)
	var sum, i int32
	for ; s.Scan(); {
		// each line holds a single sign-prefixed integer
		if _, err = fmt.Sscanf(s.Text(), "%d", &i); err != nil {
			panic(err)
		}
		sum += i
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	fmt.Println(sum)
}
