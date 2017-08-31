package main

import (
	"fmt"
)

var f float64
var b float64
var fb float64
var fblimit int

func main() {

	for fblimit = 1; fblimit <= 1000; fblimit++ {

		fizzc := 0.0
		buzzc := 0.0
		fizzbuzzc := 0.0

		for i := 0; i <= fblimit; i++ {

			switch {
			case i%15 == 0:
				fizzbuzzc++
			case i%5 == 0:
				buzzc++
			case i%3 == 0:
				fizzc++

			}

		}

		f = float64(fizzc) / float64(fblimit)
		b = float64(buzzc) / float64(fblimit)
		fb = float64(fizzbuzzc) / float64(fblimit)

		fmt.Printf("%.2f \t %.2f \t %.2f \n", f, b, fb)

	}

}
