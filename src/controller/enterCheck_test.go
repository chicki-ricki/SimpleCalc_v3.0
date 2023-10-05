package controller

import (
	"testing"
)

var (
	testCasesCheckBrackets = []struct {
		val    string
		expect bool
	}{
		{"({[]})", true},
		{"({}))", false},
		{"({)}", false},
	}

	testCasesStartCheck = []struct {
		val    string
		expect float64
	}{
		{"5 * 6 + (2 - 9) =", 23},
		{"6 + 3 * (1 + 4 * 5) * 2 =", 132},
		{"1 / 2", 0.5},
		{" 5 + (-2 + 3)", 6},
		{" 5 + (+2 + 3)", 10},
		{" 5 + ( -2 + 3)", 6},
		{" 1.5 + 1.5", 3},
		{"-1.5 + (-1.5)", -3},
		{"0.66E+4 + 2", 6602},
		{"0.66E+4 + 2e+2 + 300", 7100},
		{"1.1e+10 + 1.1e+10", 2.2e+10},
		{"2.4e+10 - 2.4e+10", 0},
		{"-2 + 3", 1},
		{"(5.2e+4 + sin(0.1) - 10) + (-0.2)", 51989.899833416646},
	}

	testCasesCheckUnary = []struct {
		enter  string
		expect string
	}{
		{"-2", "0 - 2"},
		{"    tan   (   1     )", "tan ( 1 ) "},
		{"0.66 E +     4 + 2e +  2 + 300", "0.66e+4 + 2e+2 + 300"},
	}
)

func TestCheckBrackets(t *testing.T) {
	for _, testCase := range testCasesCheckBrackets {
		actual := checkBrackets(testCase.val)
		if actual != testCase.expect {
			t.Errorf("Result was incorrect, expected: %v, actual: %v\n", testCase.expect, actual)
		}
	}
}

func TestStartCheck(t *testing.T) {
	for _, testCase := range testCasesStartCheck {
		actual := StartCheck(testCase.val)
		if actual != testCase.expect {
			t.Errorf("Result was incorrect, expected: %v, actual: %v\n", testCase.expect, actual)
		}
	}
}

func TestCheckUnary(t *testing.T) {
	for _, testCase := range testCasesCheckUnary {
		actual := checkUnary(testCase.enter)
		if actual != testCase.expect {
			t.Errorf("Result was incorrect, expected: %v, actual: %v\n", testCase.expect, actual)
		}
	}
}
