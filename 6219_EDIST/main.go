/*
   Created by Smart Mek on May 15, 2018 9:10 AM
*/

package main

import (
	"fmt"
	"math"
)

var lookup = make(map[string]int)

func main() {

	var T int
	var A, B string

	fmt.Scanln(&T)
	for t := 0; t < T; t++ {
		// assuming no spaces
		fmt.Scanln(&A)
		fmt.Scanln(&B)

		// using Levenshtein's algorithm
		var n int = eDist(A, B, len(A), len(B))
		fmt.Println(n)
	}

}

func eDist(A string, B string, i, j int) int {

	key := fmt.Sprintf("%s+%s+%d+%d", A, B, i, j)
	if val, ok := lookup[key]; ok {
		return val
	}

	// d_i0 == delete
	if i == 0 {
		return j
	}

	// d_0j insert
	if j == 0 {
		return i
	}

	// d_ij == depends on whether last char for both are same
	if A[i-1] == B[j-1] {
		return eDist(A, B, i-1, j-1)
	}

	// perform delete, insert and substitution here
	minNum := 1 + min(eDist(A, B, i, j-1), eDist(A, B, i-1, j), eDist(A, B, i-1, j-1))
	lookup[key] = minNum

	return minNum
}

func min(x, y, z int) int {
	a, b, c := float64(x), float64(y), float64(z)
	return int(math.Min(math.Min(a, b), c))
}

