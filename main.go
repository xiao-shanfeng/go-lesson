package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Sex     bool
	Hobbies []string
	Home
}

type Home struct {
	P string
}

func (p *Person) Song(name string) (ret string) {
	ret = "真的很惊雷"
	fmt.Printf("%v唱了一首%v,%v", p.Name, name, ret)
	return ret
}

func main() {
	// go结构体
	//var p Person
	//p.Name = "Tom"
	//p.Age = 18
	//p.Sex = true
	//p.Hobbies = []string{"吸烟", "烫头"}

	//p := Person{
	//	Name:    "Tom",
	//	Age:     18,
	//	Sex:     true,
	//	Hobbies: []string{"吸烟"},
	//}
	//
	//var pp *Person
	//pp = &p
	//
	//pp.Age = 26

	//fmt.Println(p, pp)

	//p := new(Person)
	//
	//p.Sex = true
	//p.Age = 30
	//p.Name = "Jack"
	//
	//fmt.Println(p)

	p := Person{
		Name:    "Tom",
		Age:     18,
		Sex:     true,
		Hobbies: []string{"吸烟"},
	}
	p.P = "杭州"

	fmt.Println(p.Name)
	SayPerson(p)
	res := p.Song("惊雷")
	fmt.Println("\n", res)
}

func SayPerson(p Person) {
	fmt.Println(p)
}
