/*
 * @Description: 
 * @Author: Matt.Yu
 * @Email: lwlmm_77@163.com
 * @GitHub: github.com/lwlmm
 * @Date: 2021-02-23 01:23:50
 */

//function_calls_function.go
package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}

//my guess : GOG