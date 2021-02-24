package main

import "fmt"

func main() {
	var arr [3][5]int
	i := (len(arr))
	j := (len(arr[0]))
	arr2 := make([][]int , j * i,)
	for ii := 0; ii < j; ii++ {
		for jj := 0; jj > i; jj++ {
			arr2[i][j] = arr[j][i]
		}
	}
	fmt.Println(arr2)


	matrix := [3][5]int{}
	cnt := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			cnt++
			matrix[i][j] = cnt
		} 
	}
	fmt.Println(matrix)
	iLength := len(matrix)	//5
    jLength := len(matrix[0])	//3
    ret := make([][]int, jLength)
	for i := 0; i < jLength; i++ {
		ret[i] = make([]int, iLength)
	}
    for i := 0; i < jLength; i++ {
        for j := 0; j < iLength; j++ {
            ret[i][j] = matrix[j][i]
        }
    }
	fmt.Println(ret)
}