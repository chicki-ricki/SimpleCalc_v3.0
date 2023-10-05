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

	testCasesCalculate = []struct {
		enter  []string
		expect float64
	}{
		{[]string{"0", "cos"}, 1},
		{[]string{"0", "sin"}, 0},
		{[]string{"1", "asin"}, 1.5707963267948966},
		{[]string{"0", "acos"}, 1.5707963267948966},
		{[]string{"25", "sqrt"}, 5},
		{[]string{"1", "tan"}, 1.557407724654902},
		{[]string{"1", "atan"}, 0.7853981633974483},
		{[]string{"7.38905609893065", "ln"}, 2},
		{[]string{"100", "log"}, 2},
		{[]string{"1", "2", "+"}, 3},
		{[]string{"3", "1", "-"}, 2},
		{[]string{"4", "5", "*"}, 20},
		{[]string{"24", "6", "/"}, 4},
		{[]string{"5", "3", "mod"}, 2},
		{[]string{"2", "8", "^"}, 256},
	}
)

func TestToPolandNotation(t *testing.T) {
	for _, testCase := range testCasesPoland {
		actual := toPolandNotation(testCase.val)
		// if actual != testCase.expect {
		if !reflect.DeepEqual(actual, testCase.expect) {
			t.Errorf("Result was incorrect, expected: %v, actual: %v\n", testCase.expect, actual)
		}
	}
}

func TestCalculate(t *testing.T) {
	for _, testCase := range testCasesCalculate {
		actual := calculate(testCase.enter)
		if actual != testCase.expect {
			t.Errorf("Result was incorrect, exected: %v, actual: %v\n", testCase.expect, actual)
		}
	}
}
