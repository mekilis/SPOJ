/*
   Created by Smart Mek on Jul 11, 2018 2:50 PM
*/

package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanln(&T)
	for t := 0; t < T; t++ {
		var r, s float64
		fmt.Scan(&r)
		// BC = 2r, AB^2 = max(start)... add AC^2 incrementally
		s = r + r
		max := s * s
		fmt.Printf("Case %d: %.2f\n", t+1, max+.25)
	}
}
