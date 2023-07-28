package main

import (
	"calc/controller"
	"fmt"
)

func main() {
	enter := "654+3*(1+4*5)*2="
	rez := controller.StartCheck(enter)
	fmt.Println("rez in main:", rez)
}
