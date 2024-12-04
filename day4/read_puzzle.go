package main

import (
	"bufio"
	"log"
	"os"
)

func newPuzzleFromInput() string {
	file, err := os.Open("day4/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := ""
	for scanner.Scan() {
		lines += (scanner.Text() + "\n")

	}
	return lines[:len(lines)-1]

}
