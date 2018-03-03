/*
   Created by Smart Mek on 3/3/18 9:21 PM
*/

// Adding reversed numbers

package main

import (
	"fmt"
)

func main() {

	var N, a, b int

	// get N
	fmt.Scan(&N)

	for n := 0; n < N; n++ {

		fmt.Scan(&a)
		fmt.Scanln(&b)

		a = reverse(a)
		b = reverse(b)

		sum := reverse(a + b)
		fmt.Println(sum)

	}
}

func reverse(x int) int {

	y := 0

	for x > 0 {

		rem := x % 10
		y *= 10
		y += rem
		x /= 10
	}

	return y
}
