package calculator

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func strToArr(str string) []string {
	// str += " "
	// str = strings.ToLower(str)
	fmt.Println("enter str in strToArr:", str)
	arr := strings.Fields(str)
	// for _, arrVal := range arr {
	// 	fmt.Println("fields:", arrVal)
	// }
	// lenght := len(str)
	// strArr := make([]string, lenght)
	// num := 0
	// for i := 0; i < lenght; i++ {
	// 	/*
	// 		if str[i] >= '0' && str[i] <= '9' || str[i] == '.' || (strings.Contains(str[i:], "e") && strings.Contains(str[i:], "+")) || (i > 1 && strings.Contains(str[i-1:], "e") && strings.Contains(str[i-1:], "+")) {
	// 			if strings.Contains(strArr[num], "e+") && (str[i] < 48 || str[i] > 57) || strings.Contains(strArr[num], " +") || str[i] == ' ' {
	// 				num++
	// 			}
	// 			strArr[num] += string(str[i])
	// 		} else {
	// 			if str[i] == '=' {
	// 				break
	// 			}
	// 			num++
	// 			strArr[num] += string(str[i])
	// 			num++
	// 		}
	// 	*/
	// 	/*
	// 	fmt.Println("strToArr|string(str[", i, "]):", string(str[i]))
	// 	if str[i] == '(' || str[i] == ')' || str[i] == '{' || str[i] == '}' ||
	// 		str[i] == '[' || str[i] == ']' || str[i] == '*' || str[i] == '/' ||
	// 		// str[i] == '^' || ((str[i] == '+' || str[i] == '-') && str[i+1] == ' ') {
	// 		str[i] == '^' || (str[i] == '+' || str[i] == '-') {
	// 		num++
	// 		strArr[num] += string(str[i])
	// 		num++
	// 	} else if (str[i] >= 'a' && str[i] < 'e') || (str[i] > 'e' && str[i] <= 'z') {
	// 		strArr[num] += string(str[i])
	// 	} else {
	// 		strArr[num] += string(str[i])

	// 	}
	// }
	// j := 0
	// for i, val := range strArr {
	// 	strArr[i] = strings.TrimSpace(val)
	// 	if strArr[i] != "" {
	// 		// arr[j] = strArr[i]
	// 		j++
	// 	}
	// }
	// arr := make([]string, j)
	// j = 0
	// for i, val := range strArr {
	// 	strArr[i] = strings.TrimSpace(val)
	// 	if strArr[i] != "" {
	// 		arr[j] = strArr[i]
	// 		j++
	// 	}
	// }
	// fmt.Println("strToArr|arr:", arr)
	// // return strArr
	// return arr
	// // fmt.Println("strToArr|strArr:", strArr)
	// // return strArr
	// */
	fmt.Println("strToArr|arr:", arr)
	return arr
}

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

	fmt.Println("expression:", expression)
	for _, val := range expression {
		fmt.Println("expression val:", val)
		fmt.Println("strings.Contains(operators, ", val, "):", strings.Contains(operators, val))
		if strings.Contains(operators, val) {
			if len(stack) < 2 {
				if strings.Contains(unaryop, val) {
					lenght := len(stack)
					n1 := stack[lenght-1]
					stack = stack[:lenght-1]
					switch {
					case val == "sin":
						temp := math.Sin(n1)
						fmt.Println("sin(", n1, "):", temp)
						stack = append(stack, temp)
					}
					continue
				} else {
					// log.Println("Too few arguments")
					fmt.Println("calculate|len(stack):", len(stack))
					fmt.Println("stack in if:", stack)
					break
				}
			}
			lenght := len(stack)
			fmt.Println("stack1: ", stack)
			n1 := stack[lenght-1]
			fmt.Println("calculate|n1:", n1)
			// stack[lenght-1] = ""
			stack = stack[:lenght-1]
			lenght = len(stack)
			fmt.Println("stack2: ", stack)
			n2 := stack[lenght-1]
			fmt.Println("calculate|n2:", n2)
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
			// case val == "sin":
			// 	stack = append(stack, n2)
			// 	temp = math.Sin(n1)
			// 	fmt.Println("sin(", n1, "):", temp)
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
				fmt.Println("Num of (", val, "):", num)
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
	for _, val := range strArr {
		fmt.Println("StartCalculate|val: ", val)
	}
	notation := toPolandNotation(strArr)
	fmt.Println("notation:", notation)
	rez = calculate(notation)
	return rez
}
