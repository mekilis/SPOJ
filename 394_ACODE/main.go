/*
   Created by Smart Mek on 3/8/18 12:26 PM
*/

package main

import (
	"fmt"
)

var lookup = make([]string, 10000)

func main() {

	var n string
	fmt.Scan(&n)

	for c := 'A'; c <= 'Z'; c++ {
		lookup[c-64] = string(c)
	}

	// use DP

	for n != "0" {
		// do n
		fmt.Scanln(&n)
	}

}

func compute(s string, step int) string {
	return 0
}
