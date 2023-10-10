package controller

import (
	"reflect"
	"testing"
)

var (
	testCasesGraphic = []struct {
		equation string
		start    int
		end      int
		pixels   int
		expect   []SCoordinates
	}{
		{"x     -4", 0, 9, 10, []SCoordinates{{0, -4}, {1, -3}, {2, -2}, {3, -1}, {4, 0}, {5, 1}, {6, 2}, {7, 3}, {8, 4}, {9, 5}}},
	}
)

func TestStartGraphic(t *testing.T) {
	for _, testCase := range testCasesGraphic {
		actual, _ := StartGraphic(testCase.equation, testCase.start, testCase.end, testCase.pixels)
		if !reflect.DeepEqual(actual, testCase.expect) {
			t.Errorf("Result was incorrect, expected: %v, actual: %v\n", testCase.expect, actual)
		}
	}
}
