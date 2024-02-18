package main

import (
	"fmt"
	"introduction/calc"
	"introduction/greeting"
	"introduction/strcon"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(quote.GoV3())
	fmt.Println(calc.Add(1, 2))
	fmt.Println(greeting.Hello("Mike ross"))
	s := strcon.SwapCase("Gopher")
	l := strcon.SwapCase("gopher")
	fmt.Println("Converted string is: ", s)
	fmt.Println("Converted string is: ", l)
}
