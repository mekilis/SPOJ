/*
   Created by Smart Mek on May 02, 2018 12:38 PM
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	//var threeTerm, threeLastTerm, sum int64 = 3, 8, 55
	//var threeTerm, threeLastTerm, sum int64 = 10, 55, 650 // using big.Int instead
	var threeTerm, threeLastTerm, sum = big.NewInt(144115188075856304), big.NewInt(144115188075855890), big.NewInt(7349874591868660947)
	var T = 1

	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		fmt.Scan(threeTerm)
		fmt.Scan(threeLastTerm)
		fmt.Scan(sum)

		var a, d, n, l, opRight, opLeft = new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int)
		// Formula: Sn = (n/2)(a + l) == (n/2)(2a + (n - 1)d)
		// But a + l = 3term + 3lastterm, l = 3t + 3lt - a
		// Therefore, Sn = (n/2)(3t + 3lt)
		opRight.Add(threeTerm, threeLastTerm)
		opLeft.Mul(big.NewInt(2), sum)

		//n = (2 * sum) / (threeTerm + threeLastTerm)
		n.Div(opLeft, opRight)
		fmt.Println(n)

		// 3t = a + 2d -- i, 3lt = a + (n - 1 -2)d = a + (n - 3)d
		// Using elimination ==> d = (3t - 3lt) / (5 - n), Since n >= 7, division by zero can never occur
		// d = (threeTerm - threeLastTerm) / (5 - n)
		opRight.Sub(big.NewInt(5), n)
		opLeft.Sub(threeTerm, threeLastTerm)
		d.Div(opLeft, opRight)

		//a = threeTerm - 2*d
		opRight.Mul(big.NewInt(2), d)
		a.Sub(threeTerm, opRight)

		//l := threeTerm + threeLastTerm - a
		l.Add(threeTerm, threeLastTerm)
		l.Sub(l, a)

		for i := a; i.Cmp(l) <= 0; i.Add(i, d) {
			fmt.Print(i, " ")
		}

		fmt.Println()
	}
}