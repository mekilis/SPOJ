/*
   Created by Smart Mek on 6/2/2018, 11:03:30 AM
*/

package main

import (
	"fmt"
)

func main() {
	var T int
	var S string

	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		fmt.Scan(&S)

		l := len(S)
		count := findMinimum(S, 0, l)
		fmt.Println(count)
	}
}

var lookup [][]int

func findMinimum(S string, l, G int) int {

	lookup = make([][]int, 0)
	for g := 0; g < G; g++ {
		lookup = append(lookup, make([]int, G))
	}

	// fill table
	for g := 1; g < G; {
		for l, h := 0, g; h < G; {
			if S[l] == S[h] {
				lookup[l][h] = lookup[l+1][h-1]
			} else {
				lookup[l][h] = min(lookup[l][h-1], lookup[l+1][h]) + 1
			}
			l++
			h++
		}
		g++
	}

	return lookup[0][G-1]
}

func min(x, y int) (z int) {
	z = y
	if x < y {
		z = x
	}
	return
}
