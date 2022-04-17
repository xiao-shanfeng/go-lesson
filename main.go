package main

import "fmt"

func main() {
	// 自增自减
	// go 中只有a++，没有++a
	a := 0
	a++
	//a--
	//++a
	fmt.Println(a)

	// 条件语句
	if a > 10 {
		fmt.Println("这是正确的")
	} else if a == -1 {
		fmt.Println("中间情况")
	} else {
		fmt.Println("这是错误的")
	}

	// 选择语句
	// go中switch语句中，case自带的break，如果想实现别的语言的穿透，使用fallthrough
	switch a {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough
	default:
		fmt.Println("这是默认情况")
	}

	// 循环语句
	for {
		a++
		if a == 5 {
			break
		}
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for a < 10 {
		fmt.Println(a)
		a++
	}

	// 跳转语句

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	//A:
	for {
		if a > 15 {
			//break A
			goto B
		}
		fmt.Println("这是A")
		a++
	}
	//return
B:
	fmt.Println("这是B")
}
