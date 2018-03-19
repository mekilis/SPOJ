/*
   Created by Smart Mek on 3/19/18 4:04 PM
*/

package main

import (
	"math/big"
	"fmt"
)

func main() {
	var t int
	var K = new(big.Int)

	fmt.Scanln(&t)
	for tt := 0; tt < t; tt++ {
		fmt.Scanln(K)

		fmt.Println("This was entered:", K)
	}
}
