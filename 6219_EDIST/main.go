/*
   Created by Smart Mek on May 15, 2018 9:10 AM
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

var lookup []int

func main() {

	var T int
	var A, B string

	fmt.Scanln(&T)
	for t := 0; t < T; t++ {
		// assuming no spaces
		fmt.Scanln(&A)
		fmt.Scanln(&B)
		lookup = make([]int, len(B)+1)

		// using Levenshtein's algorithm
		var n = eDist(A, B)
		fmt.Println(n)
	}

}

// using a Non recursive (faster) solution
func eDist(A, B string) int {

	var minNum int
	lookup = make([]int, utf8.RuneCountInString(B)+1)

	for i := range lookup {
		lookup[i] = i
	}

	// d_i*
	for _, valA := range A {
		j, temp := 1, lookup[0]
		lookup[0]++

		for _, valB := range B {
			// delete, insert
			minNum = min(lookup[j]+1, lookup[j-1]+1)

			if valA != valB {
				// substitute
				minNum = min(minNum, temp+1)
			} else {
				// update
				minNum = min(minNum, temp)
			}

			temp, lookup[j] = lookup[j], minNum
			j++
		}
	}

	return lookup[len(lookup)-1]
}

func min(x, y int) int {

	if x <= y {
		return x
	}

	return y
}
