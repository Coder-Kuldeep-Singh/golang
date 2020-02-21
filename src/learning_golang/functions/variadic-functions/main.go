package main

import "fmt"

//Variadic functions can be called
//with any number of trailing arguments.
func Sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	Sum(1, 2, 3, 4, 5, 5, 6)
	Sum(1, 2, 2, 3, 3, 34, 4, 4, 4, 4, 5, 5)
	nums := []int{1, 2, 32, 3, 4, 5, 5, 6, 7, 7}
	Sum(nums...)
}
