package main

import "fmt"

func main(){

	//使用type 关键词定义新的类型 Gender
	//go里面没有枚举，一般可以用常量的方式来模拟
	type Gender int8
	const (
		MALE 	Gender = 1
		FEMALE	Gender = 2
	)

	gen := MALE

	switch gen{
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}
	//go的switch 没有break，如果需要继续向下执行，使用fallthrough


}