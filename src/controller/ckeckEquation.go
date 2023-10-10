package controller

import (
	"errors"
	"strconv"
	"strings"
)

func StartEquation(str string, x float64) (y float64, err error) {
	err = nil
	if checkBrackets(str) {
		convertStr := strings.ReplaceAll(str, "x", strconv.FormatFloat(x, 'f', 2, 64)) //Itoa(x))
		y, _ = StartCheck(convertStr)
	} else {
		err = errors.New("Error of brackets, please enter new expression")
	}

	return y, err
}
