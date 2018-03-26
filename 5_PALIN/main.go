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
		l := len(KArray)

		// check if it's all 9s
		allnines := true

		for _, c := range KArray {

			if c != "9" {
				//fmt.Println("not all nines")
				allnines = false
				break
			}
		}

		if allnines {
			KArray = append(KArray, "1")
			KArray[0] = "1"
			for i := 1; i < l; i++ {
				KArray[i] = "0"
			}

 		} else {
			// check if its already a palindrome
			mid := l / 2

			// assuming odd array
			i, j := mid, mid

			// if even
			if l%2 == 0 {
				//right = KArray[mid+1:]
				i = mid - 1
			}

			for i >= 0 && KArray[i] == KArray[j] {
				i -= 1
				j += 1
			}

			//fmt.Println("final value of i and j", i, j)

			var leftNum string
			var bigSum = new(big.Int)

			// has it crossed the boundary?
			if (i < 0 && KArray[i + 1] == KArray[j - 1]) || KArray[i] < KArray[j] {
				// it is a palindrome...
				//fmt.Println("it's a palindrome or something else")
				updateBigNum(leftNum, KStr, KArray, mid, l, bigSum)
			} else {
				//fmt.Println("it's no palindrome")
				// good to go
				for i := i; i >= 0; i-- {
					KArray[l-i-1] = KArray[i]
				}

			}
		}
		fmt.Println(strings.Join(KArray, ""))
	}
}
func updateBigNum(leftNum, KStr string, KArray []string, mid, l int, bigSum *big.Int) {
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

	for i := 0; i < l/2; i++ {
		KArray[i] = string(leftNum[i])
		KArray[l-i-1] = KArray[i]
	}
}
