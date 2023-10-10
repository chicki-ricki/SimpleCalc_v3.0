package main

import (
	"calc/controller"
	"fmt"
)

func main() {
	// enter := "5*6+(2-9)" //23
	// enter := "6+3*(1+4*5)*2" //132
	// enter := "1/2"
	// enter := " 5+(-2+3)"
	// enter := " 5+(+2*3)"
	// enter := " 5+( -2*3)"
	// enter := " 1.5+1.5"
	// enter := "    -1.5 +    (      -1.5)"
	// enter := "0.66E+4 + 2" //6602
	// enter := "0.66 E +     4 + 2e +  2 + 300" //0.66*10^4 + 2*10^2 + 300 = 7100
	// enter := "1.1e+10 + 1.1e+10" //22000000000
	// enter := "2.4e+10 - 2.4e+10"
	// enter := "-2 + 3"
	// enter := "sin(0.35) + 1" //1.34
	// enter := "    tan   (   1     )" //1.56
	// enter := "(5.2e+4 + sin(0.1) * 1000) + (-0.2)" //52099.6
	// enter := "-2"
	// enter := "+5"
	enter := "sqrt   (  +  25)"

	rez, _ := controller.StartCheck(enter)
	fmt.Println("rez expression in main:", rez)
	fmt.Printf("rez expression in main round: %.1f\n", rez)

	// equation := "2*x + 5"
	// equation := "(2*x + 5))"
	equation := "x     -4"
	// equation := "(5.2e+4 + sin(0.1) * x) + (-0.2)"
	start := -100
	end := 100
	pixels := 10 //need 600
	rezGraphic, _ := controller.StartGraphic(equation, start, end, pixels)
	fmt.Println("equation slice of structs for graphic:", rezGraphic)

	rezEq := controller.StartEquation(equation)
	fmt.Println("equation result:", rezEq)

	// graphic.Window()
}
