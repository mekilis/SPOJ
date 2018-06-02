/*
   Created by Smart Mek on 6/2/2018, 11:03:30 AM
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	var S string

	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		fmt.Scan(&S)

		l := len(S)
		count := findMinimum(S, 0, l-1)
		fmt.Println(count)
	}
}

var lookup = make(map[string]int)

func findMinimum(S string, l, h int) int {
	key := S[l : h+1]
	if v, ok := lookup[key]; ok {
		//fmt.Println("solved before", key, v)
		return v
	}
	if l > h {
		return math.MaxInt32
	}

	if l == h {
		return 0
	}

	if l == h-1 {
		if S[l] == S[h] {
			return 0
		}
		return 1
	}

	minimum := 0

	if S[l] == S[h] {
		minimum = findMinimum(S, l+1, h-1)
	} else {
		minimum = min(findMinimum(S, l, h-1), findMinimum(S, l+1, h)) + 1
	}

	lookup[key] = minimum
	return minimum
}

func min(x, y int) (z int) {
	z = y
	if x < y {
		z = x
	}
	return
}
