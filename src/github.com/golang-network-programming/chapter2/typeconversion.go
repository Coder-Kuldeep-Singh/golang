package main
import "fmt"

func main() {
	//type casting
	var byt []byte
	byt = []byte("strings")
	var stored string
	stored = string(byt[:])
	//convert string into array
	//var arr = []string
	fmt.Printf("%T",stored)
}
