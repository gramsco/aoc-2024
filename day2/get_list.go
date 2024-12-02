package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getListOfReports() [][]int {
	file, err := os.Open("day2/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reports := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		level := []int{}
		txt := scanner.Text()
		splitted := strings.Split(txt, " ")

		for _, value := range splitted {
			success, err := strconv.Atoi(value)

			if err != nil {
				panic(err)
			}

			level = append(level, success)
		}
		reports = append(reports, level)
	}

	return reports

}
