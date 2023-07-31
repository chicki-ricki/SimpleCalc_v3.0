package main

import (
	"calc/controller"
	"fmt"
	"log"

	"github.com/Knetic/govaluate"
)

func main() {
	// enter := "5*6+(2-9)"
	// enter := "6+3*(1+4*5)*2"
	// enter := "1/2"
	// enter := " 5+(-2+3)"
	// enter := " 5+( -2+3)"
	// enter := " 1.5+1.5"
	// enter := " -1.5+(-1.5)"
	// enter := "0.66E+4 + 2"
	enter := "0.66E+4 + 2e+2 + 300" // don't work!!!
	// enter := "1.1e+10 + 1.1e+10" // don't work!!!
	// enter := "2.4691357802e+10 - 2.4691357802e+10" // don't work!!!
	// enter := "-2 + 3" // don't work!!!

	rez := controller.StartCheck(enter)
	fmt.Println(rez)

	expression, err := govaluate.NewEvaluableExpression(enter)
	if err != nil {
		log.Fatalf("Error with expression: %s\n", err)
	}
	result, err := expression.Evaluate(nil)
	fmt.Println("govaluate.rezult: ", result)

}
