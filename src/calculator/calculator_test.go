package calculator

import (
	"reflect"
	"testing"
)

var (
	testCasesPoland = []struct {
		val    []string
		expect []string
	}{
		{[]string{"654", "+", "3", "*", "(", "1", "+", "4", "*", "5", ")", "*", "2"}, []string{"654", "3", "1", "4", "5", "*", "+", "*", "2", "*", "+"}},
		{[]string{"5", "*", "6", "+", "(", "2", "-", "9", ")"}, []string{"5", "6", "*", "2", "9", "-", "+"}},
	}
)

func TestToPolandNotation(t *testing.T) {
	for _, testCase := range testCasesPoland {
		actual := toPolandNotation(testCase.val)
		// if actual != testCase.expect {
		if !reflect.DeepEqual(actual, testCase.expect) {
			t.Errorf("Result was incorrect, expected: %v, actual: %v", testCase.expect, actual)
		}
	}
}
