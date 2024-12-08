package main

import (
	"testing"
)

func assert(t *testing.T, b any, a any) {
	if a != b {
		t.Fatalf("Expected %d, received %d", a, b)
	}

}

func TestPart2(t *testing.T) {
	assert(t, 156, reduceCompute(156, []int{15, 6}, true))
}

func TestParseStr(t *testing.T) {
	res := parseStr("190: 10 19")
	assert(t, res[0], 190)
	assert(t, res[1], 10)
	assert(t, res[2], 19)

	res2 := parseStr("1466166947230: 6 880 5 97 90 719 23")
	assert(t, res2[0], 1466166947230)
	assert(t, res2[len(res2)-1], 23)

}
