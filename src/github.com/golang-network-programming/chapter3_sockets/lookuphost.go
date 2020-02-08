package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
	}
	name := os.Args[1]
	addrs, err := net.LookupHost(name)
	Error(err)
	fmt.Println(addrs)
	// for _, s := range addrs {
	// 	fmt.Println(s)
	// }
	os.Exit(0)
}

func Error(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
}
