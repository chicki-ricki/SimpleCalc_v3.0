package controller

import (
	"testing"
)

var (
	testCases = []struct {
		val    string
		expect bool
	}{
		{"({[]})", true},
		{"({}))", false},
		{"({)}", false},
	}
)

func TestCheckBrackets(t *testing.T) {
	for _, testCase := range testCases {
		actual := CheckBrackets(testCase.val)
		if actual != testCase.expect {
			t.Errorf("Result was incorrect, expected: %v, actual: %v", testCase.expect, actual)
		}
	}
}
