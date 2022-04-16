package main

import (
	"fmt"
	"lesson/dog"
	"lesson/testpackage"
)

func main() {
	// go声明变量形式一
	var a string = "Hello, go"
	// go声明变量形式二
	b := "Hello Go"
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(testpackage.A)
	fmt.Println(testpackage.B)

	fmt.Println(dog.Name)
}
