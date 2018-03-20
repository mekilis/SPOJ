/*
   Created by Smart Mek on 3/19/18 4:04 PM
*/

package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	var t int
	var K = new(big.Int)

	fmt.Scanln(&t)
	for tt := 0; tt < t; tt++ {
		fmt.Scanln(K)

		done := false

		fmt.Println("This was entered:", K)
		KStr := K.String()
		KArray := strings.Split(KStr, "")
		// check if it's all 9s
		l := int64(len(KArray))
		divisor := new(big.Int)
		divisor.Set(big.NewInt(l))
		divisor.Exp(big.NewInt(10), divisor, nil)
		divisor.Sub(divisor, big.NewInt(1))
		fmt.Println(divisor)

		if KStr == divisor.String() {
			fmt.Println("This number is all 9s")
			divisor.Add(divisor, big.NewInt(2))
			fmt.Println("next palindrome is", divisor.String())
			done = true
		}

		if !done {
			// check if its already a palindrome
			mid := l/ 2 + 1
			left := KArray[:mid]
			right := KArray[mid:]
			fmt.Println("This is the breakdown", left, right)


		}
	}
}
