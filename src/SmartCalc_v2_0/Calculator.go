package SmartCalc_v2_0

import (
	"fmt"
	"log"
)

func checkBrackets(str string) bool {
	var stack []string
	for _, char := range str {
		lenght := len(stack) - 1
		switch {
		case char == '(':
			stack = append(stack, ")")
		case char == '[':
			stack = append(stack, "]")
		case char == '{':
			stack = append(stack, "}")
		case char == ')' || char == '}' || char == ']':
			if len(stack) == 0 || stack[lenght] != string(char) {
				return false
			}
			stack[lenght] = ""
			stack = stack[:lenght]
		default:
			break
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func StartCalculate(str string) (rez float64) {
	rez = -1.0
	for {
		if checkBrackets(str) {
			// rez = calculate(str)
			fmt.Println("check OK")
			return rez
		} else {
			log.Println("Error of brackets")
			return rez
		}
	}
}
