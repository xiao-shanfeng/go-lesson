package _interfacex

import "fmt"

type Animal interface {
	Eat()
	Run()
}

type Cat struct {
	Name string
	Sex  bool
}

type Dog struct {
	Name string
}

func (c *Cat) Eat() {
	fmt.Printf("%v吃东西\n", c.Name)
}

func (c *Cat) Run() {
	fmt.Printf("%v跑了\n", c.Name)
}

func (d *Dog) Eat() {
	fmt.Printf("%v吃东西", d.Name)
}

func _interface() {
	// go 接口
	c := Cat{
		Name: "Tom",
		Sex:  true,
	}
	var a Animal
	a = &c
	fmt.Println(a)
	fmt.Println(c)
	c.Eat()
	c.Run()
	a.Eat()
	a.Run()

	d := Dog{Name: "Mary"}
	//a = d // 因为Dog没有实现Animal中的Run方法，所以不能赋给a
	fmt.Println(d)
	fmt.Println(a)
	//fmt.Println(a.Name) // 报错。因为a是一个接口，实现了结构体的方法，并不能获取结构体的属性

	fmt.Println("======================")
	animal := Cat{
		Name: "Tom",
		Sex:  true,
	}
	fmt.Println(animal)
	animal.Eat()
	animal.Run()

	MyFunc([]int{1, 4, 6, 7})
	MyFunc(3)
}

func MyFunc(a interface{}) {
	fmt.Println(a)
}
