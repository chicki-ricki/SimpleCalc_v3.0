package controller

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type SCoordinates struct {
	x float64
	y float64
}

func StartGraphic(str string, start int, end int, pixels int) (ret []SCoordinates, err error) {
	var delta float64
	var deltaPixel float64
	err = nil
	maxVal := math.Max(float64(start), float64(end))
	minVal := math.Min(float64(start), float64(end))
	delta = maxVal - minVal
	deltaPixel = float64(delta) / float64(pixels-1)
	if checkBrackets(str) {
		for i := 0; i < pixels; i++ {
			var elem SCoordinates
			x := minVal + float64(i)*deltaPixel
			convertStr := strings.ReplaceAll(str, "x", strconv.Itoa(i))
			// y, err := StartCheck(convertStr)
			// if err == nil {
			// 	elem.x = x
			// 	elem.y = y
			// 	ret = append(ret, elem)
			// }
			y, _ := StartCheck(convertStr)
			elem.x = x
			elem.y = y
			ret = append(ret, elem)
		}
	} else {
		err = errors.New("Error of brackets, please enter new expression")
	}
	return ret, err
}
