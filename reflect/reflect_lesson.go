package reflect

import (
	"fmt"
	"reflect"
)

// 反射

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	Class string
	User
}

func (u User) SayName(name string) {
	fmt.Println("我的名字叫", name)
}

func reflectFun() {
	u := User{
		Name: "Tom",
		Age:  23,
		Sex:  true,
	}

	//s := Student{
	//	Class: "五年四班",
	//	User:  u,
	//}

	//check(&s)
	check(u)
	//fmt.Println(s)
}

func check(v interface{}) {
	fmt.Println(v)
	//typeOf := reflect.TypeOf(v)
	valueOf := reflect.ValueOf(v)
	//fmt.Println(typeOf)  //main.User
	//fmt.Println(valueOf) //{Tom 23 true}
	//// 通过反射获取某个结构体的某个字段
	//name := valueOf.FieldByName("Name")
	//fmt.Println(name)
	//// 结构体的属性总数
	//fmt.Println(typeOf.NumField())
	//for i := 0; i < typeOf.NumField(); i++ {
	//	//按照索引取值
	//	fmt.Println(valueOf.Field(i))
	//}
	//
	////按照层级取数据
	//fmt.Println(valueOf.FieldByIndex([]int{0}))    //五年四班
	//fmt.Println(valueOf.FieldByIndex([]int{1}))    //{Tom 23 true}
	//fmt.Println(valueOf.FieldByIndex([]int{1, 0})) //Tom
	//
	//ty := typeOf.Kind()
	//fmt.Println(ty)
	//if ty == reflect.Struct {
	//	fmt.Println("我是struct")
	//}

	//e := valueOf.Elem()
	//e.FieldByName("Class").SetString("六年四班")
	//fmt.Println(v)

	m := valueOf.Method(0)
	m.Call([]reflect.Value{reflect.ValueOf("Jack")})
}
