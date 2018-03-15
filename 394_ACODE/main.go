/*
   Created by Smart Mek on 3/8/18 12:26 PM
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num string
	fmt.Scan(&num)

	// use DP
	for num != "0" {
		fmt.Println(compute(num))
		fmt.Scanln(&num)
	}

}

func compute(numStr string) int64 {
	// using F(N-1) + F(N-2)
	//fmt.Println("number is", numStr)

	// create slice based on the length of number
	l := len(numStr)
	if l == 1 {
		return 1
	}
	f := make([]int64, l)
	//fmt.Println("array of f", f)

	// NOTE: First digit can be coded only one way
	f[0], f[1] = 1, 1

	// Also get status of first two digits
	x, _ := strconv.Atoi(numStr[:2])
	if x > 9 && x < 27 && numStr[1:2] != "0" {
		f[1] = 2
		if l > 2 {
			if numStr[2:3] == "0" {
				//fmt.Println("bad starting number")
				f[1] = f[0]
			}
		}
	} else if x > 20 && x%10 == 0 {
		//fmt.Println("invalid starting number", x)
		return 0
	}

	for i := 2; i < l; i++ {
		// update current index to former
		f[i] = f[i-1]
		// check if previous and current are valid
		x, _ = strconv.Atoi(numStr[i-1 : i+1])
		if x > 20 && x%10 == 0 {
			//fmt.Println("invalid number in range", x)
			return 0
		}
		if x == 0 {
			//fmt.Println("00 is invalid")
			return 0
		}
		if x > 9 && x < 27 && numStr[i:i+1] != "0" {
			temp := f[i]
			f[i] += f[i-2]
			if l > i+1 {
				if numStr[i+1:i+2] == "0" {
					//fmt.Println("not possible to combine. zero exists", f, "now restoring")
					f[i] = temp
				}
			}
		}
	}

	//fmt.Println("final array of f", f)
	return f[l-1]
}
