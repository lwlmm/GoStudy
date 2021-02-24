package main

import "fmt"

func main() {
	i := 0
	for {
		if i >= 3 { break }
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement after the loop")
}