package main

import (
	"calculator_v3_0/SmartCalc_v2_0"
	"fmt"
)

func main() {
	stack := "({})"
	rez := SmartCalc_v2_0.StartCalculate(stack)
	fmt.Println("rez in main:", rez)
}
