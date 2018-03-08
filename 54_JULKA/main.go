/*
   Created by Smart Mek on 3/8/18 8:36 AM
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// operations needed for big numbers
	// 10 - 2 = 8. 8 / 2 = 4. 4 + 2 = 6
	// +, -, /
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(105), nil)

	klaudia, natalia := new(big.Int), new(big.Int)

	for t := 0; t < 10; t++ {

		bigTotal := new(big.Int)
		bigDifference := new(big.Int)

		fmt.Scanln(bigTotal)
		fmt.Scanln(bigDifference)

		if bigTotal.Cmp(&limit) < 0 {
			klaudia, natalia = big.NewInt(0), big.NewInt(0)

			// compute total - difference
			bigTotal.Sub(bigTotal, bigDifference)

			// compute and print higher share
			natalia.Div(bigTotal, big.NewInt(2))
			klaudia.Add(natalia, bigDifference)

			fmt.Println(klaudia)
			fmt.Println(natalia)
		}
	}
}