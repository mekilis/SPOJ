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
	var s string
	var K = new(big.Int)

	fmt.Scanln(&t)
	for tt := 0; tt < t; tt++ {
		fmt.Scanln(&s)
		var ok bool
		K, ok = K.SetString(s, 10)
		if !ok {
			panic("wrong string representation of number")
		}
		done := false

		//fmt.Println("This was entered:", K)
		KStr := K.String()

		if len(KStr) == 1 {
			K.Add(K, big.NewInt(1))
			KStr = K.String()
			if KStr != "10" {
				fmt.Println(KStr)
			} else {
				fmt.Println(11)
			}
			continue
		}

		KArray := strings.Split(KStr, "")
		// check if it's all 9s
		l := int64(len(KArray))
		divisor := new(big.Int)
		divisor.Set(big.NewInt(l))
		divisor.Exp(big.NewInt(10), divisor, nil)
		divisor.Sub(divisor, big.NewInt(1))
		//fmt.Println(divisor)

		if KStr == divisor.String() {
			//fmt.Println("This number is all 9s")
			divisor.Add(divisor, big.NewInt(2))
			fmt.Println(divisor.String())
			done = true
		}

		if !done {
			// check if its already a palindrome
			mid := l / 2
			//left := KArray[:mid]
			//right := KArray[mid:]

			// assuming even array
			i, j := mid-1, mid

			// if odd
			if l%2 != 0 {
				//right = KArray[mid+1:]
				j = mid + 1
			}
			//fmt.Println("This is the breakdown", left, right)

			// ignore parts of K where K[i] and K[j] are the same
			k := int64(0)
			for k = i; k >= 0; k-- {

				ii, jj := k, l-k-1
				if KArray[ii] != KArray[jj] {
					break
				}
				// update i and j now
				i, j = ii, jj

			}

			if k != i && i != 0 {
				// update i and j to the next offset
				i--
				j++
			}

			//fmt.Println("final value of i and j", i, j)

			var leftNum string
			var bigSum = new(big.Int)

			// has it crossed the boundary?
			if i == 0 && KArray[i] == KArray[j] {
				// it is a palindrome...
				//fmt.Println("it's a palindrome")
				updateBigNum(leftNum, KStr, KArray, mid, l, bigSum)
			} else {
				//fmt.Println("it's no palindrome")
				// apparently K[i] != k[j]
				if KArray[i] > KArray[j] {

					// good to go
					for i := i; i >= 0; i-- {
						KArray[l-i-1] = KArray[i]
					}
				} else {
					// little modification
					updateBigNum(leftNum, KStr, KArray, mid, l, bigSum)
				}

			}

			fmt.Println(strings.Join(KArray, ""))

		}
	}
}
func updateBigNum(leftNum, KStr string, KArray []string, mid, l int64, bigSum *big.Int) {
	if l%2 != 0 {
		leftNum = KStr[:mid+1]
	} else {
		leftNum = KStr[:mid]
	}

	//fmt.Println("left num =", leftNum)
	bigSum.SetString(leftNum, 10)
	bigSum.Add(bigSum, big.NewInt(1))
	leftNum = bigSum.String()
	//fmt.Println("left sum is now", leftNum)

	if l%2 != 0 {
		KArray[mid] = string(leftNum[mid])
	} else {
		KArray[mid] = string(leftNum[mid-1])
	}

	for i := int64(0); i < l/2; i++ {
		KArray[i] = string(leftNum[i])
		KArray[l-i-1] = KArray[i]
	}
}
