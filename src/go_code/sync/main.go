package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//模拟下载，耗时2秒
func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	wg.Done()	//为wg减去一个计数
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)	//为wg增加一个计数
		//关键词go 启动新的协程实现并发
		go download("a.com/" + string(i+'0'))
	}
	wg.Wait()	//等待所有协程结束
	fmt.Println("Done!")
}

//执行的时候增加 time 关键词，可看到运行时间