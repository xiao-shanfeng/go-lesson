package _goroutine

import (
	"fmt"
	"sync"
)

func FGoroutine() {
	var wg sync.WaitGroup
	wg.Add(1)
	go Run(&wg)
	wg.Wait()
	//time.Sleep(1 * time.Second)
	// 让程序停止1秒钟
	// goroutine 和 channel
	//i := 0
	//for i < 10 {
	//	i++
	//	fmt.Println(i)
	//}
}

func Run(wg *sync.WaitGroup) {
	fmt.Println("run函数跑起来了")
	wg.Done()
}
