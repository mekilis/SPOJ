/*
   Created by Smart Mek on 3/8/18 12:26 PM
*/

package main

import (
	"fmt"
	"strconv"
)

var lookup = make(map[string]int64, 100)

func main() {

	var num string
	fmt.Scan(&num)

	//for c := 'A'; c <= 'Z'; c++ {
	//	lookup[string(c)] = int64(c - 64)
	//}

	// use DP
	for num != "0" {
		fmt.Println(compute(num))
		fmt.Scanln(&num)
	}

}

func compute(numStr string) int {
	// using F(N-1) + F(N-2)
	fmt.Println("number is", numStr)

	// create slice based on the length of number
	l := len(numStr)
	f := make([]int, l)
	fmt.Println("array of f", f)

	// NOTE: First digit can be coded only one way
	f[0], f[1] = 1, 1

	// Also get status of first two digits
	x, _ := strconv.Atoi(numStr[:2])
	if x > 9 && x < 27 {
		f[1] = 2
	}

	for i := 2; i < l; i++ {
		// update current index to former
		f[i] = f[i-1]
		// check if previous and current are valid
		x, _ = strconv.Atoi(numStr[i-1:i+1])
		if x > 9 && x < 27 {
			f[i] += f[i-2]
		}
	}

	fmt.Println("final array of f", f)

	return f[l-1]
}
