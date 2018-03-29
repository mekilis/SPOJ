/*
   Created by Smart Mek on 3/28/18 2:24 PM
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	var n, d, m, top, bot int64

	fmt.Scanln(&T)
	for t := 0; t < T; t++ {
		fmt.Scan(&n)
		// n = d(d + 1) / 2
		// d = (-1 +/- floor(sqrt(1 + 8n))) / 2 ; use only the +ve part
		d = int64(-1 + math.Floor(math.Sqrt(1 + (8 * float64(n))))) / 2
		m = n - d * (d + 1) / 2

		// top / bot
		if m < 2 {
			top = d + m
			bot = 1
		} else {
			top = d - m + 2
			bot = m
		}

		if d % 2 != 0 {
			top, bot = bot, top
		}

		fmt.Printf("TERM %d IS %d/%d\n", n, top, bot )
	}
}
