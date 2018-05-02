/*
   Created by Smart Mek on May 02, 2018 12:38 PM
*/

package main

import "fmt"

func main() {
	var threeTerm, threeLastTerm, sum uint64 = 3, 8, 55
	var T = 1

	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		fmt.Scan(&threeTerm)
		fmt.Scan(&threeLastTerm)
		fmt.Scan(&sum)

		var i, d uint64
		i = 2
		d = 1
		summer := uint64(0)
		first, second, next := uint64(0), uint64(0), uint64(0)
		var n uint64

		for {
			if i == 2 {
				// sum first two terms
				first = threeTerm - 2 * d
				second = threeTerm - d
				summer += first + second
				i++

				next = second
				continue
			}
			next += d
			summer += next

			if summer == sum && first + (i - 3)*d == threeLastTerm {
				n = i
				break
			}

			if summer > sum {
				// wrong guess for d
				d++
				// reset
				i = 2
				summer = 0
				continue
			}

			i++
		}

		fmt.Println(n)

		for i := uint64(0); i < n; i++ {
			fmt.Print(first + i * d, " ")
		}

		fmt.Println()
	}
}
