package main

import (
	"strings"
)

func solvePart2(s string) int {
	lines := strings.Split(s, "\n")
	count := 0

	for i, line := range lines {
		for j, char := range line {

			if char == 'A' {
				if i > 0 && i < len(lines)-1 && j > 0 && j < len(line)-1 {
					topLeftAndBottomRight := string(lines[i-1][j-1]) + string(lines[i+1][j+1])
					topRightAndBottomLeft := string(lines[i-1][j+1]) + string(lines[i+1][j-1])
					switch topLeftAndBottomRight + topRightAndBottomLeft {
					case "MSMS", "MSSM", "SMSM", "SMMS":
						count += 1
					}

				}
			}
		}
	}

	return count
}
