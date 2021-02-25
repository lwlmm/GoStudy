/*练习 7.7 sum_array.go

a) 写一个 Sum 函数，传入参数为一个 32 位 float 数组成的数组 arrF，
返回该数组的所有数字和。如果把数组修改为切片的话代码要做怎样的修改？
如果用切片形式方法实现不同长度数组的的和呢？

b) 写一个 SumAndAverage 方法，返回两个 int 和 float32 类型
的未命名变量的和与平均值。*/

package main

import "fmt"

func main() {
	arr := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(Sum(arr))
}

func Sum(arr []float32) float32 {
	var sum float32
	for _, f := range arr {
		sum += f
	}
	return sum
}