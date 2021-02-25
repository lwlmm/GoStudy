package main

import "fmt"

func main() {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2    	//for range 中是传值而非传指针
	}
	fmt.Println(items)

	for i, _ := range items {
		items[i] *= 2
	}
	fmt.Println(items)
}