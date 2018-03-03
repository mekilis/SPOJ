/* 
    Created by Smart Mek on 3/3/18 9:51 PM
*/

package main

import (
	"math"
	"fmt"
)

func main() {

	// f(n) = (k) E (i=1) [n / (5^i) ]
	// k = log(base 5) n

	var T int
	var N float64

	fmt.Scanln(&T)

	for t := 0; t < T; t++ {

		fmt.Scanln(&N)

		// using change of base formula
		k := math.Log(N) / math.Log(5.0)

		sum := 0.0
		for kk := math.Floor(k); kk > 0; kk-- {

			fact := math.Pow(5.0, kk)

			s := N / fact

			sum += math.Floor(s)
		}

		fmt.Println(int64(sum))
	}


}
