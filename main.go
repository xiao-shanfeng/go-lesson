package main

import "fmt"

func main() {
	// go 数组和切片
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6}
	//c1 := arr1[2:]
	//c1 := arr1[:]
	c1 := arr1[2:5]
	//c1 := arr1[:5]
	//c1[0] = 7
	fmt.Println(c1)
	fmt.Println(arr1)
	fmt.Println(len(c1), cap(c1))

	// 使用make创建切片
	//c2 := make([]int, 4, 6)
	c2 := make([]int, 4)
	//var c2 []int
	c2[0] = 0
	c2[1] = 1
	c2[2] = 2
	c2[3] = 3
	c2 = append(c2, 4)
	c2 = append(c2, 5)
	c2 = append(c2, 6)
	c2 = append(c2, 7)
	c2 = append(c2, 8)
	fmt.Println(c2, len(c2), cap(c2))

	// copy复制切片
	// arr1 [0,1,2,3,4,5,6]
	c3 := arr1[:]
	c4 := arr1[2:3]
	// copy 把后者覆盖前者
	copy(c3, c4)
	fmt.Println(c3, c4) //2,1,2,3,4,5,6   2
}
