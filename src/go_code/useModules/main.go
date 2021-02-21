package main

import (
	"fmt"
	"useModules/calc"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(calc.Add(10, 3))
}