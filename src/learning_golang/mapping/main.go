//Mapping in Golang
package main

import "fmt"

func main() {
	//use make function to create empty map
	m := make(map[string]int) //key and values pair
	m["k1"] = 3
	m["k2"] = 5
	fmt.Println("Map:", m)
	// fmt.Println("map: ", m["k1"])

	v1 := m["k1"]
	fmt.Println("v1", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//check the key and value exists in the map
	// and it prints the boolean if exists true
	//if not exists false
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
