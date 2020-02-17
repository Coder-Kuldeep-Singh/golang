package main

import "fmt"

func a1() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i, " ")
	}
}
func a2() {
	for i := 0;i<3;i++ {
		defer func() { fmt.Println(i," ") }()
	}
}

func a3() {
	for i := 0;i<3;i++ {
		defer func(n int) { fmt.Println() }
	}
}