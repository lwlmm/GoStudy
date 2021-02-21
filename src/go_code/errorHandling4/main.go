package main

import(
	"fmt"
)

//使用defer 和 recover来处理panic的情况
func get(index int) (ret int) {
	//使用defer定义了处理异常的函数，在协程退出前，会执行完defer挂载的任务
	defer func() {
		//如果 r ！= nil (即触发了panic),控制权就交给了defer
		if r := recover(); r != nil {	//使用recover使程序恢复正常
			fmt.Println("Some error happened", r)
			ret = -1	//如果不处理返回值，返回值将被默认设置成0
		}
	}()
	arr := [3]int{2, 3, 4}
	ret = arr[index]
	return
}

func main() {
	fmt.Println(get(5))
	fmt.Println("Runtime finished!")
}