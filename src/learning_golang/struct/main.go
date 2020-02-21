package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string) *Person {
	p := Person{name: name}
	p.age = 20
	return &p
}
func main() {
	fmt.Println(Person{"Bob", 19})
	fmt.Println(Person{name: "Jhandu", age: 10})
	fmt.Println(&Person{name: "abcd", age: 123})
	fmt.Println(NewPerson("xyz"))
	s := Person{name: "LMK", age: 234}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)

	sp.age = 8987
	fmt.Println(sp.age)

}
