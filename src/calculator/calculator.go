package calculator

import (
	"fmt"
	"strings"
)

// func CheckBrackets(str string) bool {
// 	var stack []string
// 	for _, char := range str {
// 		lenght := len(stack) - 1
// 		switch {
// 		case char == '(':
// 			stack = append(stack, ")")
// 		case char == '[':
// 			stack = append(stack, "]")
// 		case char == '{':
// 			stack = append(stack, "}")
// 		case char == ')' || char == '}' || char == ']':
// 			if len(stack) == 0 || stack[lenght] != string(char) {
// 				return false
// 			}
// 			stack[lenght] = ""
// 			stack = stack[:lenght]
// 		default:
// 			break
// 		}
// 	}
// 	if len(stack) == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func strToArr(str string) []string {
	lenght := len(str)
	strArr := make([]string, lenght+2)
	num := 0
	for i := 0; i < lenght; i++ {
		fmt.Println("str[", i, "]:", string(str[i]))
		if str[i] >= 48 && str[i] <= 57 {
			strArr[num] += string(str[i])
		} else {
			if str[i] == '=' {
				break
			}
			num++
			strArr[num] += string(str[i])
			num++
		}
	}
	return strArr
}

func toPolandNotation(strArr []string) (expression []string) {
	var stack []string
	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"^": 1,
		"*": 2,
		"/": 2,
		"%": 2,
		"(": 3,
		"{": 3,
		"[": 3,
	}
	open_brackets := "({["
	close_brackets := ")}]"
	// add_op := "+-^"
	// mul_op := "*/%"
	operators := "+-^*/%"
	// unary_op := "+-"
	for _, char := range strArr {
		switch {
		case char != "" && char[0] >= '0' && char[0] <= '9':
			expression = append(expression, string(char))
		// case strings.Contains(open_brackets, char):
		// 	stack = append(stack, string(char))
		// case strings.Contains(mul_op, char):
		// 	stack = append(stack, char)
		// case strings.Contains(add_op, char) || strings.Contains(mul_op, char) || strings.Contains(open_brackets, char):
		case strings.Contains(operators, char):
			lenght := len(stack)
			fmt.Println("lenght of stack:", lenght, ", stack:", stack)
			if lenght == 0 || stack[lenght-1] == "(" {
				stack = append(stack, string(char))
			} else {
				for {
					fmt.Println("precedence of", char, "(char):", precedence[char])
					if lenght > 0 {
						fmt.Println("precedence of", stack[lenght-1], "(stack[lenght-1]):", precedence[stack[lenght-1]])
					}
					// fmt.Println("stack after precedence:", stack)
					if lenght > 0 && char != "" && precedence[char] <= precedence[stack[lenght-1]] && stack[lenght-1] != "(" {
						expression = append(expression, stack[lenght-1])
						stack[lenght-1] = ""
						stack = stack[:lenght-1]
						lenght = len(stack)
					} else {
						if char != "" && char != ")" {
							stack = append(stack, char)
						}
						fmt.Println("break, stack:", stack)
						break
					}
				}
			}
		case strings.Contains(open_brackets, char):
			stack = append(stack, string(char))
		case strings.Contains(close_brackets, char):
			lenght := len(stack)
			for {
				// if lenght > 0 {
				if strings.Contains(open_brackets, stack[lenght-1]) {
					break
				}
				expression = append(expression, stack[lenght-1])
				stack[lenght-1] = ""
				stack = stack[:lenght-1]
				lenght = len(stack)
			}
			if char == ")" {
				stack[lenght-1] = ""
				stack = stack[:lenght-1]
			}
			// }
		// case strings.Contains(unary_op, char):
		// 	stack = append(stack, char)
		default:
			break
		}
	}
	if len(stack) > 0 {
		lenght := len(stack)
		for {
			if lenght == 0 {
				break
			}
			expression = append(expression, stack[lenght-1])
			stack[lenght-1] = ""
			stack = stack[:lenght-1]
			lenght = len(stack)
		}
	}
	fmt.Println("**stack:", stack)
	fmt.Println("**expression:", expression)

	return
}

func StartCalculate(str string) (rez float64) {
	rez = -1.0
	strArr := strToArr(str)
	for i, val := range strArr {
		fmt.Println("strArr[", i, "]:", val, "|")
		if val == "" {
			fmt.Println("i with _:", i)
		}
	}
	notation := toPolandNotation(strArr)
	fmt.Println("Poland notation:", notation)
	// rez = calculate(notation)
	return rez
}
