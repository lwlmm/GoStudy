package main

import "fmt"

func Season(month int) string {
	switch month {
	case 1, 2, 12:
		return "winter"
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "fall"
	default:
		return "wrong input"
	}
}
func main() {
	fmt.Println(Season(1))
	fmt.Println(Season(9))
	fmt.Println(Season(200))
}