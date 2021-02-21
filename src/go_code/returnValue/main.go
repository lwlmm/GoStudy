package main

import "fmt"

func add(num1 int, num2 int) int{
	return num1 + num2
}

func div(num1 int, num2 int) (int, int){
	return num1 / num2, num1 % num2
}

//可以给返回值命名，以简化return
func mul(num1 float32, num2 float32) (ans float32){
	ans = num1 * num2
	return
}

func main(){
	quo, rem := div(100, 17)
	fmt.Println(quo, rem)
	fmt.Println(add(100, 17))
	newAns := mul(8.8, 7.7)
	fmt.Println(newAns)
}

