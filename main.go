package main

import "fmt"

func main() {
	// go 数组和切片
	// 定义数组 var 变量名 = [变量长度]变量类型{变量值1,变量值2,...,变量值n}
	var arr1 = [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", arr1)
	fmt.Println(arr1)
	var arr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr2)
	fmt.Println(arr2)
	var arr3 [5]int
	fmt.Println(arr3)
	arr4 := new([5]int)
	arr4[3] = 3
	fmt.Println(arr4)

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	for i, v := range arr1 {
		fmt.Println(i, v)
	}

	// 数组长度len，数组容量cap
	fmt.Println(len(arr1), cap(arr1))

}
