package aoc

import (
	"bufio"
	"os"
)

func CallForEachLineInFile(name string, fun func(string) error) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	// tokenize the buffered input line by line
	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)
	for s.Scan() {
		if err := fun(s.Text()); err != nil {
			return err
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}
