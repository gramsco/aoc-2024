package main

import (
	"fmt"
	"testing"
)

func (m *Multiplication) isEqual(m2 Multiplication) bool {
	return m.a == m2.a && m.b == m2.b
}

func (m *Multiplication) toString() string {
	return fmt.Sprintf("mul(%d,%d)", m.a, m.b)
}

func TestParsing(t *testing.T) {
	parsed := Parser("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")

	if !parsed[0].isEqual(Multiplication{2, 4}) {
		t.Fatalf("expected mul(2,4), received %s", parsed[0].toString())
	}
}

func TestParsing2(t *testing.T) {
	parsed := ParserWithInstructions("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

	expectedResults := []Multiplication{{2, 4}, {8, 5}}
	for i, expectedResult := range expectedResults {
		if !parsed[i].isEqual(expectedResult) {
			t.Fatalf("expected %s, received %s", expectedResult.toString(), parsed[i].toString())
		}
	}

}

func TestMultiplication(t *testing.T) {
	m := getMultiplication("mul(2,4)")
	if !m.isEqual(Multiplication{2, 4}) {
		t.Fatalf("Expected mul(2,4), received %s", m.toString())
	}

	m2 := getMultiplication("mul(11,8)")
	if !m2.isEqual(Multiplication{11, 8}) {
		t.Fatalf("Expected mul(11,8), received %s", m.toString())
	}
}
