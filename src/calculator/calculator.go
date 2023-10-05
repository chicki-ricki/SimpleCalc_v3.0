package calculator

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func toPolandNotation(strArr []string) (expression []string) {
	var stack []string
	precedence := map[string]int{
		"+":    1,
		"-":    1,
		"^":    1,
		"*":    2,
		"/":    2,
		"mod":  2, // % (целочисленное деление)
		"cos":  3,
		"sin":  3,
		"tan":  3,
		"acos": 3,
		"asin": 3,
		"atan": 3,
		"sqrt": 3,
		"ln":   3,
		"log":  3,
		"(":    4,
		"{":    4,
		"[":    4,
	}
	open_brackets := "({["
	close_brackets := ")}]"
	operators := "+-^*/modacosasinatansqrtlnlog"
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
	operators := "+-^*/modacosasinatansqrtlnlog"
	unaryop := "acosasinatansqrtlnlog"
	var stack []float64

	for _, val := range expression {
		var temp float64
		if strings.Contains(operators, val) {
			if len(stack) < 2 {
				if strings.Contains(unaryop, val) {
					lenght := len(stack)
					n1 := stack[lenght-1]
					stack = stack[:lenght-1]
					switch {
					case val == "cos":
						temp = math.Cos(n1)
					case val == "sin":
						temp = math.Sin(n1)
					case val == "tan":
						temp = math.Tan(n1)
					case val == "acos":
						temp = math.Acos(n1)
					case val == "asin":
						temp = math.Asin(n1)
					case val == "atan":
						temp = math.Atan(n1)
					case val == "sqrt":
						temp = math.Sqrt(n1)
					case val == "ln":
						temp = math.Log(n1)
					case val == "log":
						temp = math.Log10(n1)
					}
					stack = append(stack, temp)
					continue
				} else {
					log.Println("Too few arguments")
					break
				}
			}
			lenght := len(stack)
			n1 := stack[lenght-1]
			stack = stack[:lenght-1]
			lenght = len(stack)
			n2 := stack[lenght-1]
			stack = stack[:lenght-1]
			switch {
			case val == "+":
				temp = n2 + n1
			case val == "-":
				temp = n2 - n1
			case val == "^":
				temp = math.Pow(n2, n1)
			case val == "*":
				temp = n2 * n1
			case val == "mod":
				temp = math.Mod(n2, n1)
			case val == "/":
				if n1 == 0 {
					log.Println("Error: division by zero")
					break
				}
				temp = n2 / n1
			case val == "cos":
				stack = append(stack, n2)
				temp = math.Cos(n1)
			case val == "sin":
				stack = append(stack, n2)
				temp = math.Sin(n1)
			case val == "tan":
				stack = append(stack, n2)
				temp = math.Tan(n1)
			case val == "acos":
				stack = append(stack, n2)
				temp = math.Acos(n1)
			case val == "asin":
				stack = append(stack, n2)
				temp = math.Asin(n1)
			case val == "atan":
				stack = append(stack, n2)
				temp = math.Atan(n1)
			case val == "sqrt":
				stack = append(stack, n2)
				temp = math.Sqrt(n1)
			case val == "ln":
				stack = append(stack, n2)
				temp = math.Log(n1)
			case val == "log":
				stack = append(stack, n2)
				temp = math.Log10(n1)
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
	strArr := strings.Fields(str)
	notation := toPolandNotation(strArr)
	rez = calculate(notation)
	return rez
}
