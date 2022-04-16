package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		go语言 四种基本数据类型
		1。整型
		2。浮点型
		3。字符型
		4。布尔型
	*/
	// 整型：unit、unit8、unit16、unit32、int、int8、int16、int32
	var num1 int
	var num2 int
	num1 = 999
	num2 = -999
	fmt.Println(num1)
	fmt.Println(num2)

	// 浮点型：float32、float64
	var f1 float64
	f1 = 3.1415926
	fmt.Println(f1)

	// 字符型
	var str1 string
	str1 = "Hello， go"
	fmt.Println(str1)

	// 变量类型检测
	var number1 int
	var number2 uint
	var string1 string
	number1 = 123
	number2 = 1234
	string1 = "我是一个字符串"
	fmt.Printf("%T %T %T\n", number1, number2, string1)

	// 类型转换
	// 1. string -> int
	var str2 string
	str2 = "123"
	i, _ := strconv.Atoi(str2)
	fmt.Println(i)
	fmt.Printf("%T", i)
	// 2. string -> int64
	var str3 string
	str3 = "1234"
	i2, _ := strconv.ParseInt(str3, 10, 64)
	fmt.Println(i2)
	fmt.Printf("%T\n", i2)
	// 3. int -> string
	var num3 int
	num3 = 234
	str4 := strconv.Itoa(num3)
	fmt.Println(str4)
	fmt.Printf("%T\n", str4)
	// 4.int64 -> string
	var num4 int64
	num4 = 456
	str5 := strconv.FormatInt(num4, 10)
	fmt.Println(str5)
	fmt.Printf("%T\n", str5)
	// 5.string -> float64/float32
	var str6 string
	str6 = "3.1415926"
	f2, _ := strconv.ParseFloat(str6, 32)
	fmt.Println(f2)
	fmt.Printf("%T\n", f2)
	// 6. int -> int64
	var num5 int
	num5 = 123
	num6 := int64(num5)
	fmt.Println(num6)
	fmt.Printf("%T\n", num6)
	// 7. int64 -> int
	var num7 int64
	num7 = 23
	num8 := int(num7)
	fmt.Println(num8)
	fmt.Printf("%T\n", num8)

	// 变量的声明
	name := "Tom"
	age := 26
	flag := true
	score := 59.5
	fmt.Println(name, age, flag, score)
}
