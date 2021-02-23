package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I'm Matt, Hi"

	fmt.Println("The position of Matt is:",strings.Index(str, "Matt"))

	fmt.Println("The position of the first instance of Hi is:",
				strings.Index(str, "Hi"))
	fmt.Println("The position of the last instance of Hi is:",
				strings.LastIndex(str, "Hi"))
	
	fmt.Println("The position of Null is:", strings.Index(str, "Null"))
}