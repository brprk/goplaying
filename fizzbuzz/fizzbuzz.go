package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 100; i++ {

		switch {
		case i%15 == 0:
			fmt.Printf("%s \n", "fizzbuzz")
		case i%5 == 0:
			fmt.Printf("%s \n", "buzz")
		case i%3 == 0:
			fmt.Printf("%s \n", "fizz")
		default:
			fmt.Printf("%d \n", i)
		}

	}

}
