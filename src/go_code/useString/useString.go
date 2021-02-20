package main

import (
	"fmt"
	"reflect"
)
func main(){
	str1 := "Golang"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str1[1]).Kind())
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) // 可以知道某个变量的类型
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%d %c\n", str2[2], str2[2])
	fmt.Println("len(str2): ",len(str2))

	runeArr := []rune(str2)
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())
	fmt.Println(runeArr[2], string(runeArr[2]))
	fmt.Println("len(runeArr): ", len(runeArr))
}