/*
 * @Description: 
 * @Author: Matt.Yu
 * @Email: lwlmm_77@163.com
 * @GitHub: github.com/lwlmm
 * @Date: 2021-02-23 01:48:33
 */
 package main
 
 import (
	 "fmt"
	 "math/rand"
	 "time"
 )

 func main() {
	 for i := 0; i < 10; i++ {
		 a := rand.Int()
		 fmt.Println("a[", i, "] = ",a)
	 }
	 for i := 0; i < 5; i++ {
		 r := rand.Intn(8)
		 fmt.Println("r[", i, "] = ",r)
	 }
	 //timens := int64(time.Now().Nanosecond())
	 //rand.Seed(timens)
	 rand.Seed(int64(time.Now().Nanosecond()))
	 for i := 0; i < 10; i++ {
		 fmt.Println("f[", i, "] = ", 100*rand.Float32())
	 }
 }
