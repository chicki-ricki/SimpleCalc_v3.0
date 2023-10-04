package controller

import (
	"calc/calculator"
	"fmt"
	"log"
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

	// fmt.Printf("checkUnary|s=%s|\n", s)
	// for i, char := range s {
	for _, char := range s {
		// if len(retStr) > 0 {
		// 	fmt.Println("checkUnary|char:", string(char))
		// 	fmt.Println("checkUnary|retStr[len(retStr):-2]:", retStr[len(retStr)-2:], "len:", len(retStr))
		// } else {
		// 	fmt.Println("checkUnary|char:", string(char))
		// }
		if char == ' ' {
			continue
		}
		if char == ')' || char == '(' {
			retStr += " " + string(char) + " "
		} else if unicode.IsDigit(char) || unicode.IsLetter(char) || char == '.' || (len(retStr) > 1 && retStr[len(retStr)-1:] == "e") {
			retStr += string(char)
		} else if len(retStr) == 0 && (char == '-' || char == '+') {
			retStr += "0 " + string(char) + " "
			// fmt.Println("retStr += \"0 \" + string(char)|", retStr)
		} else if (char == '+' || char == '-') && (len(retStr) > 1 && retStr[len(retStr)-2:len(retStr)-1] == "(") {
			retStr += " 0 " + string(char) + " "
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			retStr += " " + string(char) + " "
		} else { //if char != ' ' {
			fmt.Println("checkUnary|else|char:", string(char))
			retStr += string(char) + " "
		}
	}
	// fmt.Println("checkUnary|str return: ", retStr)
	return retStr
}

func StartCheck(str string) (rez float64) {
	rez = -1.0
	for {
		lenght := len(str) - 1
		if str[lenght] == '=' {
			str = str[:lenght]
		}
		if checkBrackets(str) {
			str = checkUnary(str)
			// fmt.Printf("StartCheck|str after checkUnary=%s|\n", str)
			rez = calculator.StartCalculate(str)
			break
		} else {
			log.Println("Error of brackets, please enter new expression")
			fmt.Scan(&str)
		}
	}
	return
}
