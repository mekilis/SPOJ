/*
   Created by Smart Mek on May 02, 2018 12:38 PM
*/

package main

import "fmt"

func main() {
	//var threeTerm, threeLastTerm, sum uint64 = 3, 8, 55
	var threeTerm, threeLastTerm, sum uint64 = 10, 55, 650
	var T = 1

	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		fmt.Scan(&threeTerm)
		fmt.Scan(&threeLastTerm)
		fmt.Scan(&sum)

		var d uint64
		d = 1

		first, last := uint64(0), uint64(0)
		var n uint64

		for {
			net := threeTerm + threeLastTerm
			first = threeTerm - 2 * d
			last = net - first
			n = (net - first) / d

			if sum == (n / 2) * (first + last) {
				break
			}

			d++
		}

		fmt.Println(n)

		for i := uint64(0); i < n; i++ {
			fmt.Print(first + i * d, " ")
		}

		fmt.Println()
	}
}
