package main

import (
	"testing"
)

func assert(t *testing.T, b any, a any) {
	if a != b {
		t.Fatalf("Expected %d, received %d", a, b)
	}

}

func TestCompute(t *testing.T) {

	assert(t, calculate([]int{1, 2, 3}, []string{"add", "add"}), 6)
	assert(t, calculate([]int{1, 2, 3}, []string{"add", "multiply"}), 9)
	assert(t, calculate([]int{1, 2, 3}, []string{"multiply", "add"}), 5)
	assert(t, calculate([]int{1, 2, 3}, []string{"multiply", "multiply"}), 6)

	assert(t, calculate([]int{3, 4, 5, 6}, []string{"add", "multiply", "add"}), 41)

	assert(t, calculate([]int{11, 6, 16, 20}, []string{"add", "multiply", "add"}), 292)

}

func TestExtractOperation(t *testing.T) {
	assert(t, ExtractOperation(0, 1)[0], "add")
	assert(t, ExtractOperation(7, 3)[0], "multiply")
	assert(t, ExtractOperation(7, 3)[1], "multiply")
	assert(t, ExtractOperation(7, 3)[2], "multiply")
}

func TestResult(t *testing.T) {

	assert(t, ResultOrZero(190, 10, 19), 190)
	assert(t, ResultOrZero(3267, 81, 40, 27), 3267)
	assert(t, ResultOrZero(83, 17, 5), 0)
	assert(t, ResultOrZero(156, 15, 6), 0)
	assert(t, ResultOrZero(7290, 6, 8, 6, 15), 0)
	assert(t, ResultOrZero(161011, 16, 10, 13), 0)
	assert(t, ResultOrZero(192, 17, 8, 14), 0)
	assert(t, ResultOrZero(21037, 9, 7, 18, 13), 0)
	assert(t, ResultOrZero(292, 11, 6, 16, 20), 292)
}

func TestBinaryStr(t *testing.T) {
	assert(t, binaryString(2, 4), "0010")
}

func TestParseStr(t *testing.T) {
	res := parseStr("190: 10 19")
	assert(t, res[0], 190)
	assert(t, res[1], 10)
	assert(t, res[2], 19)

	res2 := parseStr("1466166947230: 6 880 5 97 90 719 23")
	assert(t, res2[0], 1466166947230)
	assert(t, res2[len(res2)-1], 23)
	// assert(t, parseStr("3267: 81 40 27"), 3267)
	// assert(t, parseStr("83: 17 5"), 0)
	// assert(t, parseStr("156: 15 6"), 0)
	// assert(t, parseStr("7290: 6 8 6 15"), 0)
	// assert(t, parseStr("161011: 16 10 13"), 0)
	// assert(t, parseStr("192: 17 8 14"), 0)
	// assert(t, parseStr("21037: 9 7 18 13"), 0)
	// assert(t, parseStr("292: 11 6 16 20"), 292)

}
