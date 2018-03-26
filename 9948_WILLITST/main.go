/*
   Created by Smart Mek on 3/26/18 12:42 PM
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var n uint64

	fmt.Scanln(&n)

	for i := uint64(1); i <= math.MaxUint64; i *= 2 {
		//fmt.Println(i)
		if i == n {
			fmt.Println("TAK")
			break
		} else if i > n {
			fmt.Println("NIE")
			break
		}
	}

	//counter := 0
	//
	//for n > 1 && counter < 1000 {
	//
	//	fmt.Println(n)
	//
	//	if n%2 == 0 {
	//		n /= 2
	//	} else {
	//		n = 3 * (n + 3)
	//	}
	//
	//	counter++
	//}
}
