package controller

import (
	"calc/calculator"
	"fmt"
	"log"
)

// проверка баланса скобок
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

// проверка на наличие унарных операций -2 => 0-2
func checkUnary(str string) string {

	return str
}

func StartCheck(str string) (rez float64) {
	rez = -1.0
	for {
		lenght := len(str) - 1
		if str[lenght] == '=' {
			str = str[:lenght]
		}
		if checkBrackets(str) {
			str = checkUnary(str) // to do this func !!!
			rez = calculator.StartCalculate(str)
			break
		} else {
			log.Println("Error of brackets, please enter new expression")
			fmt.Scan(&str)
		}
	}
	return
}
