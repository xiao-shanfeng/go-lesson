package main

import "fmt"

func main() {
	// map的使用
	m1 := make(map[string]string)
	m1["name"] = "Tom"
	m1["age"] = "12"
	fmt.Println(m1)
	fmt.Printf("%T\n", m1)

	// 下面这种声明必须要初始化一下
	var m2 map[int]int
	m2 = map[int]int{}
	m2[1] = 1
	m2[2] = 2
	fmt.Println(m2)

	m3 := map[int]string{1: "tom", 2: "jack"}
	fmt.Println(m3)

	m4 := make(map[int]interface{})
	m4[1] = "tom"
	m4[2] = 12
	m4[3] = true
	m4[4] = [...]int{1, 2, 3, 4}
	fmt.Println(m4, len(m4))

	delete(m4, 4)
	fmt.Println(m4, len(m4))

	for k, v := range m4 {
		fmt.Println(k, v)
	}
}
