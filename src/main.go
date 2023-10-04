package main

import (
	"calc/controller"
	"fmt"
)

func main() {
	// enter := "5*6+(2-9)"
	// enter := "6+3*(1+4*5)*2" //132
	// enter := "1/2"
	// enter := " 5+(-2+3)"
	// enter := " 5+(+2+3)"
	// enter := " 5+( -2*3)"
	// enter := " 1.5+1.5"
	// enter := "    -1.5 +    (      -1.5)"
	// enter := "0.66E+4 + 2"
	// enter := "0.66 E +     4 + 2e +  2 + 300" //0.66*10^4 + 2*10^2 + 300
	// enter := "1.1e+10 + 1.1e+10"
	// enter := "2.4e+10 - 2.4e+10"
	// enter := "-2 + 3"
	enter := "sin(0.35) + 1"
	// enter := "(5.2e+4 + sin(0.1) * 1000) + (-0.2)"

	rez := controller.StartCheck(enter)
	fmt.Println("rez in main:", rez)
	fmt.Printf("rez in main round: %.1f\n", rez)

	// expression, err := govaluate.NewEvaluableExpression(enter)
	// if err != nil {
	// 	log.Fatalf("Error with expression: %s\n", err)
	// }
	// result, err := expression.Evaluate(nil)
	// fmt.Println("govaluate.rezult: ", result)

	// // // fmt.Println(0.52e+4)
	// // rez := "last"
	// // fmt.Println(rez[len(rez)-1:])
	// // fmt.Println(rez[len(rez)-2 : len(rez)-1])
}
