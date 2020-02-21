package main

import "fmt"

func main() {
	nums := []int{2, 30, 4, 90, 23, 3, 034}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//This function works according to ASCII codes
	for i, c := range "Golang" {
		fmt.Println(i, c)
	}
	kywrds := map[string]int{"v1": 1, "v2": 2, "v3": 3}
	for k, v := range kywrds {
		fmt.Printf("%s --> %d\n", k, v)
	}
}
