package main

import "fmt"

func main() {
	var meters int
	var total int

	fmt.Println("Please enter a number to add 2 ")
	fmt.Scan(&meters)

	total = meters + 2

	fmt.Println("total meters + 2 is ", total)

	fmt.Printf("memory address is in %v \n", &meters)

}
