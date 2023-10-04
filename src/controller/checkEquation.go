package controller

import (
	"strconv"
	"strings"
)

type SCoordinates struct {
	x int
	y float64
}

// func convertToExpression(str string, i int) string {
// 	var ret string
// 	for _, val := range str {
// 		if val == 'x' {
// 			ret += strconv.Itoa(i)
// 		} else {
// 			ret += string(val)
// 		}
// 	}
// 	// fmt.Println("convertToExp|ret:", ret)
// 	return ret
// }

func StartEquation(str string, start int, end int) []SCoordinates {
	// count := 10 //need 600
	var ret []SCoordinates
	if checkBrackets(str) {
		for i := start; i < end; i++ {
			var elem SCoordinates
			x := i
			// convertStr := convertToExpression(str, i)
			convertStr := strings.ReplaceAll(str, "x", strconv.Itoa(i))
			y := StartCheck(convertStr)
			elem.x = x
			elem.y = y
			ret = append(ret, elem)
		}
	}
	return ret
}
