package main

import (
	"calc/controller"
	"fmt"
)

func main() {
	// enter := "5*6+(2-9)="
	// enter := "6+3*(1+4*5)*2="
	enter := "1/2"
	rez := controller.StartCheck(enter)
	fmt.Println(rez)
}
