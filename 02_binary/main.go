package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %#x - %o\n", 42, 42, 42, 42)

	//printing utf8
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
