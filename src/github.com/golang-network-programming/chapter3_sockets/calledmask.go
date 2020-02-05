package main

import (
	"fmt"
	"os"
	"strconv"
	"net"
	"log"
)

func main() {
	if len(os.Args) != 4 {
		log.Println(os.Stderr, "Usage: %s dotted-ip-addr ones bits\n",os.Args[0])
	}
	dotAddr := os.Args[1]
	//fmt.Println(dotAddr)
	ones, err := strconv.Atoi(os.Args[2]) // changing data type
	Errorhandle(err)
	bits ,err := strconv.Atoi(os.Args[3]) // changing data type
	Errorhandle(err)
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		log.Println("Invalid Address")
		os.Exit(1)
	}
	mask := net.CIDRMask(ones, bits)
	network := addr.Mask(mask)
	fmt.Println("Address is",addr.String())
	fmt.Println("Mask lenght is ",bits)
	fmt.Println("Leading one count is ",ones)
	fmt.Println("Mask is (hex) ",mask.String())
	fmt.Println("Network is ",network.String())
	os.Exit(0)
}

func Errorhandle(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

