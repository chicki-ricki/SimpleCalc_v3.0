package controller

import (
	"calc/calculator"
	"errors"
	"strings"
	"unicode"
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
	var retStr string
	s := strings.TrimSpace(str)
	s = strings.ToLower(str)
	for _, char := range s {
		if char == ' ' {
			continue
		}
		if char == ')' || char == '(' {
			retStr += " " + string(char) + " "
		} else if unicode.IsDigit(char) || unicode.IsLetter(char) || char == '.' || (len(retStr) > 1 && retStr[len(retStr)-1:] == "e") {
			retStr += string(char)
		} else if len(retStr) == 0 && (char == '-' || char == '+') {
			retStr += "0 " + string(char) + " "
		} else if (char == '+' || char == '-') && (len(retStr) > 1 && retStr[len(retStr)-2:len(retStr)-1] == "(") {
			retStr += " 0 " + string(char) + " "
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			retStr += " " + string(char) + " "
		} else {
			retStr += string(char) + " "
		}
	}
	return retStr
}

func StartCheck(str string) (rez float64, err error) {
	rez = -1.0
	err = nil
	for {
		lenght := len(str) - 1
		if str[lenght] == '=' {
			str = str[:lenght]
		}
		if checkBrackets(str) {
			str = checkUnary(str)
			rez, err = calculator.StartCalculate(str)
			break
		} else {
			// log.Println("Error of brackets, please enter new expression")
			// fmt.Scan(&str)
			err = errors.New("Error of brackets, please enter new expression")
		}
	}
	return
}
