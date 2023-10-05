package controller

import (
	"math"
	"strconv"
	"strings"
)

type SCoordinates struct {
	x float64
	y float64
}

func StartEquation(str string, start int, end int) []SCoordinates {
	pixels := 10 //need 600
	var delta float64
	var deltaPixel float64
	maxVal := math.Max(float64(start), float64(end))
	minVal := math.Min(float64(start), float64(end))
	delta = maxVal - minVal
	deltaPixel = float64(delta) / float64(pixels-1)
	var ret []SCoordinates
	if checkBrackets(str) {
		for i := 0; i < pixels; i++ {
			var elem SCoordinates
			x := minVal + float64(i)*deltaPixel
			convertStr := strings.ReplaceAll(str, "x", strconv.Itoa(i))
			y := StartCheck(convertStr)
			elem.x = x
			elem.y = y
			ret = append(ret, elem)
		}
	}
	return ret
}
