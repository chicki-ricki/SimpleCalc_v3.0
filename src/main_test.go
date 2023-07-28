package main

import (
	"calc/controller"
	"testing"
)

var (
	testCases = []struct {
		val    string
		expect float64
	}{
		{"5*6+(2-9)=", 23},
		{"6+3*(1+4*5)*2=", 132},
		{"1/2", 0.5},
	}
)

func TestMain(t *testing.T) {
	for _, testCase := range testCases {
		actual := controller.StartCheck(testCase.val)
		if actual != testCase.expect {
			t.Errorf("Result was incorrect, expected: %v, actual: %v", testCase.expect, actual)
		}
	}
}
