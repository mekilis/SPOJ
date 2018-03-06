/* 
    Created by Smart Mek on 3/6/18 12:33 PM
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int

	// get first three
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	for a != 0 && b != 0 || b != 0 && c != 0 || a != 0 && c != 0 {

		if b - a == c - b {

			fmt.Println("AP", c + (b - a))

		} else if b / a == c / b {

			r := float64(b) / float64(a)
			l := float64(a) * (math.Pow(r, 3.0))

			fmt.Println("GP", int64(l))

		} else {

			break
		}

		// get next three
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
	}
}
