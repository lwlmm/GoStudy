package main

import "fmt"

func main()  {
	printInductive(1)
}

func printInductive(start int) {
	if start > 10 {
		return
	} else {
		printInductive(start + 1)
		fmt.Println(start)
	}
	return
}