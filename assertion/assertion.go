package assertion

import "fmt"

/**
根据类型猜测
*/

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	Class string
	User
}

func (u User) SayHi(name string) {
	fmt.Println("我的名字叫", name)
}

func assertion() {
	u := User{
		Name: "Tom",
		Age:  23,
		Sex:  true,
	}

	s := Student{
		Class: "三年五班",
		User:  User{},
	}

	//fmt.Println(u)
	//u.SayHi("Tom")
	check(u)
	check(s)
}

func check(v interface{}) {
	//fmt.Println(v.(User).Name)
	//v.(User).SayHi(v.(User).Name)
	switch v.(type) {
	case User:
		fmt.Println(v.(User).Name)
		fmt.Println("我是User")
	case Student:
		fmt.Println(v.(Student).Class)
		fmt.Println("我是Student")
	default:
		fmt.Println("没有此类型")
	}
}
