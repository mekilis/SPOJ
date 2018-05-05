/*
   Created by Smart Mek on May 02, 2018 12:38 PM
*/

package main

import (
	"fmt"
)

func main() {
	//var threeTerm, threeLastTerm, sum int64 = 3, 8, 55
	var threeTerm, threeLastTerm, sum int64 = 10, 55, 650
	var T = 1

	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		fmt.Scan(&threeTerm)
		fmt.Scan(&threeLastTerm)
		fmt.Scan(&sum)

		var a, d, n int64

		// Formula: Sn = (n/2)(a + l) == (n/2)(2a + (n - 1)d)
		// But a + l = 3term + 3lastterm, l = 3t + 3lt - a
		// Therefore, Sn = (n/2)(3t + 3lt)
		n = (2 * sum) / (threeTerm + threeLastTerm)

		fmt.Println(n)

		// 3t = a + 2d -- i
		// 3lt = a + (n - 1 -2)d = a + (n - 3)d
		// Using elimination ==> d = (3t - 3lt) / (5 - n)
		// Since n >= 7, division by zero can never occur
		d = (threeTerm - threeLastTerm) / (5 - n)

		a = threeTerm - 2*d
		//fmt.Println(d)
		// no need to compute l

		for i := int64(0); i < n; i++ {
			fmt.Print(a, " ")
			a += d
		}

		fmt.Println()
	}
}
