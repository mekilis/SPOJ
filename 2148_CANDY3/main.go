/*
   Created by Smart Mek on 3/8/18 11:21 AM
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {

	T := new(big.Int)
	N := new(big.Int)

	fmt.Scanln(T)

	for t := big.NewInt(0); t.Cmp(T) < 0; t.Add(t, big.NewInt(1)) {

		fmt.Scanln()

		fmt.Scanln(N)
		sum := new(big.Int)
		candy := new(big.Int)
		for n := big.NewInt(0); n.Cmp(N) < 0; n.Add(n, big.NewInt(1)) {
			fmt.Scanln(candy)
			sum.Add(sum, candy)
			sum.Mod(sum, N)
		}
		if sum.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} // solved in 0.11s using 6.3M
}
