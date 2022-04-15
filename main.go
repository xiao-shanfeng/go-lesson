package main

import "fmt"

func main() {
	fmt.Println("Hello, go")
	typeStr := fmt.Sprintf("%T", 123)
	fmt.Println(typeStr)
}
