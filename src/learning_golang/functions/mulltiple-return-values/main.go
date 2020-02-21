package main

import "fmt"

func values() (int, int) {
	return 2, 4
}
func main() {
	v1, v2 := values()
	fmt.Println(v1)
	fmt.Println(v2)
	_, c := values()
	fmt.Println(c)

}
