package main

import(
	"fmt"
	"time"
)

var ch = make(chan string, 10) //创建大小为 10 的缓冲信道

func download(url string) {
	fmt.Println("Start to download ", url)
	time.Sleep(time.Second)
	ch <- url	//将url发送给信道
}

func main() {
	for i := 0; i < 3; i++ {
		go download("a.com/" + string(i + '0'))
	}
	//在协程并发进行的同时等待ch消息回传
	for i := 0; i < 3; i++ {
		msg := <-ch 	//等待信道返回消息
		fmt.Println("Download finished: ", msg)
	}
	fmt.Println("Done!")
}