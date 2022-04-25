package channel

import "fmt"

func main() {
	// channel
	// 有缓冲
	/*
		如果定义的是channel的缓冲是1，
		那么向channel里面添加数据大于1的话，程序会死锁
		那么从channel里面取出数据大于1的话，程序也会死锁
	*/
	//ch := make(chan int, 1)
	//ch := make(chan int, 5)
	// 向channel里面推数据
	//go func() {
	//	//ch <- 1
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//	}
	//}()
	// 从channel里面取数据
	//fmt.Println(<-ch)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-ch)
	//}

	//c1 := make(chan int, 5)
	//var readc <-chan int = c1
	//var writec chan<- int = c1
	//c1 <- 1
	//writec <- 2
	//fmt.Println(<-readc)
	//fmt.Println(<-c1)
	//fmt.Println(c1 == readc, c1 == writec)

	//c1 := make(chan int, 5)
	//c1 <- 1
	//c1 <- 2
	//c1 <- 3
	//close(c1) // 当channel里面数据没有了，程序还要从channel里面取数据，那程序会死锁，只有关闭channel也可以避免死锁
	//// channel关闭之后不能在添加数据了
	////c1 <- 4
	////c1 <- 5
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)

	//c1 := make(chan int, 5)
	//c1 <- 1
	//c1 <- 2
	//c1 <- 3
	//c1 <- 4
	//c1 <- 5
	//
	//close(c1)
	/*
		close 常常配合range使用
	*/
	//for v := range c1 {
	//	fmt.Println(v)
	//}

	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 1)
	//ch3 := make(chan int, 1)
	//
	//ch1 <- 1
	//ch2 <- 1
	//ch3 <- 1
	//select {
	//case <-ch1:
	//	fmt.Println("ch1")
	//case <-ch2:
	//	fmt.Println("ch1")
	//case <-ch3:
	//	fmt.Println("ch1")
	//default:
	//	fmt.Println("都没有")
	//}

	c := make(chan int)
	var readc <-chan int = c
	var writec chan<- int = c

	go SetChan(writec)

	GetChan(readc)
}

func SetChan(writec chan<- int) {
	for i := 0; i < 10; i++ {
		//fmt.Println("这是set函数")
		writec <- i
	}
}

func GetChan(readc <-chan int) {
	for i := 0; i < 10; i++ {
		//fmt.Println("这是get函数")
		fmt.Println(<-readc)
	}
}
