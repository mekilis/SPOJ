/*
   Created by Smart Mek on May 17, 2018 9:54 AM
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var L float64

	fmt.Scan(&L)
	for L != 0 {

		/*
		circumference = L
		Full circle, C = 2*pi*r = pi * d
		Semicircle = pi * r
		radius = C / pi
		area of semicircle = pi * (r^ 2) / 2
		*/
		r := L / math.Pi

		fmt.Printf("%.2f\n", math.Pi * r * r / 2)

		fmt.Scan(&L)
	}
}
