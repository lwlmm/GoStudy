//练习7.3：fibonacci_array.go: 在第 6.6 节我们看到了一个递归计算 Fibonacci 数值的方法。
//但是通过数组我们可以更快的计算出 Fibonacci 数。
//完成该方法并打印出前 50 个 Fibonacci 数字。
package main

import "fmt"

func main() {
	fmt.Println(Fib(50))
}

func Fib(n int) ([]int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1{
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr
}