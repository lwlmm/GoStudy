package main

import(
	"fmt"
	"os"
)

func main()  {

	//os.Open 返回两个值，第一个是 *File, 第二个是 error
	//如果读取成功，error的值是nil
	_, err := os.Open("88.txt")
	if err != nil {
		fmt.Println(err)
	}

}