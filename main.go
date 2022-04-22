package main

import "fmt"

func main() {
	// go 指针
	/*var a string
	a = "abc"
	var b string
	b = a
	b = "bcd"
	fmt.Println(a, b)*/
	// 指针demo
	var a string
	a = "abc"
	var b *string
	b = &a
	*b = "bcd"
	fmt.Println(a, b, *b)
	// 指针传参
	var str string = "我是main内部的"
	point(&str)
	fmt.Println(str)

	// 指针数组
	var arr1 [5]string
	arr1 = [5]string{"1", "2", "3", "4", "5"}
	//arr1[3] = "3"
	var arrP *[5]string
	arrP = &arr1
	arrP[3] = "345"
	fmt.Println(arr1, arrP, *arrP)

	// 数组指针
	var arr2 [5]*string
	var str1 string = "str1"
	var str2 string = "str2"
	var str3 string = "str3"
	var str4 string = "str4"
	var str5 string = "str5"
	arr2 = [5]*string{&str1, &str2, &str3, &str4, &str5}
	*arr2[3] = "123"
	fmt.Println(str4)

	// 指针map
	var m1 map[int]int
	m1 = map[int]int{}
	m1[1] = 1
	m1[2] = 2
	m1[3] = 3
	var m2 *map[int]int
	m2 = &m1
	(*m2)[3] = 789
	fmt.Println(m1, m2)

	// map指针
	var m3 map[int]*int
	m3 = map[int]*int{}
	var i1 int = 1
	var i2 int = 2
	var i3 int = 3
	m3[1] = &i1
	m3[2] = &i2
	m3[3] = &i3
	*(m3[2]) = 456
	fmt.Println(m3, i2)

	// 简化版
	var s string = "我是来测试地址的"
	p := &s
	*p = "123456"
	fmt.Println(s, p, *p)
}

func point(str *string) {
	*str = "我是函数内部的"
}
