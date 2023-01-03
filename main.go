package main

// struct packing

import (
	"fmt"
)


func main() {
	data := 0b10000001


	maskSex := 0b00000001
	maskAge := 0b11111110

	sex := data & maskSex
	age := data & maskAge >> 1

	fmt.Printf("%08b\n", data) 
	fmt.Printf("%08b\n", sex)
	fmt.Printf("%08b\n", age)
}