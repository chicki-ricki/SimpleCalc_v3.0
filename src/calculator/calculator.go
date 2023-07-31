package calculator

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func strToArr(str string) []string {
	// fmt.Println("enter str in strToArr:", str)
	lenght := len(str)
	strArr := make([]string, lenght+2)
	num := 0
	for i := 0; i < lenght; i++ {
		if str[i] >= 48 && str[i] <= 57 || str[i] == '.' {
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
		// "%": 2,
		"(": 3,
		"{": 3,
		"[": 3,
	}
	open_brackets := "({["
	close_brackets := ")}]"
	operators := "+-^*/%"
	for _, char := range strArr {
		switch {
		case char != "" && char[0] >= '0' && char[0] <= '9':
			expression = append(expression, string(char))
		case strings.Contains(operators, char):
			lenght := len(stack)
			if lenght == 0 || stack[lenght-1] == "(" {
				stack = append(stack, string(char))
			} else {
				for {
					if lenght > 0 && char != "" && precedence[char] <= precedence[stack[lenght-1]] && stack[lenght-1] != "(" {
						expression = append(expression, stack[lenght-1])
						stack[lenght-1] = ""
						stack = stack[:lenght-1]
						lenght = len(stack)
					} else {
						if char != "" && char != ")" {
							stack = append(stack, char)
						}
						break
					}
				}
			}
		case strings.Contains(open_brackets, char):
			stack = append(stack, string(char))
		case strings.Contains(close_brackets, char):
			lenght := len(stack)
			for {
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
	return
}

func calculate(expression []string) float64 {
	operators := "+-^*/"
	var stack []float64

	for _, val := range expression {
		if strings.Contains(operators, val) {
			if len(stack) < 2 {
				log.Println("Too few arguments")
				break
			}
			lenght := len(stack)
			n1 := stack[lenght-1]
			// stack[lenght-1] = ""
			stack = stack[:lenght-1]
			lenght = len(stack)
			n2 := stack[lenght-1]
			// stack[lenght-1] = ""
			stack = stack[:lenght-1]
			var temp float64
			switch {
			case val == "+":
				temp = n2 + n1
			case val == "-":
				temp = n2 - n1
			case val == "^":
				temp = math.Pow(n2, n1)
			case val == "*":
				temp = n2 * n1
			case val == "/":
				if n1 == 0 {
					log.Println("Error: division by zero")
					break
				}
				temp = n2 / n1
			}
			stack = append(stack, temp)
		} else {
			if num, err := strconv.ParseFloat(val, 64); err == nil {
				stack = append(stack, num)
			} else {
				log.Println("Error in strconv:", err)
			}
		}
	}
	return stack[0]
}

func StartCalculate(str string) (rez float64) {
	rez = -1.0
	strArr := strToArr(str)
	// for _, val := range strArr {
	// 	fmt.Println("val: ", val)
	// }
	notation := toPolandNotation(strArr)
	rez = calculate(notation)
	return rez
}
