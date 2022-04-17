package main

import "fmt"

func main() {
	i, s := f1(1, "23")
	fmt.Println(i, s)
	fn := func() {
		fmt.Println("匿名函数")
	}
	fn()

	f2(1, 2, 3, 4, 5, 6, 7, 8)

	(func() {
		fmt.Println("自执行函数")
	})()

	f3(34)()
	// defer 函数延后执行
	defer f4()
	fmt.Println("我后执行")
}

func f4() {
	fmt.Println("我先执行")
}

/**
go 语言中函数不能嵌套函数
*/
func f1(data1 int, data2 string) (ret1 int, ret2 string) {
	if data1 > 3 {
		return data1, data2
	} else {
		return data1, "数据是小的"
	}
}

// 可选参数函数
func f2(data1 int, data2 ...int) {
	fmt.Println(data1, data2)
	for k, v := range data2 {
		fmt.Println(k, v)
	}
}

// 闭包函数
func f3(data int) func() {
	return func() {
		fmt.Println("数据是", data)
	}
}
