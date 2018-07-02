/*
   Created by Smart Mek on Jul 02, 2018 10:39 AM
*/

package main

import "fmt"

func main() {
	var T int
	var N int64
	var x int64

	fmt.Scanln(&T)

	for t := 0; t < T; t++ {

		fmt.Scan(&N)

		if N == 0 {
			fmt.Println()
			continue
		}

		fmt.Scan(&x)
		inner, outer, _inner := x, x-x, x-x

		for i := N - N; i < N-1; i++ {
			fmt.Scan(&x)
			_inner = inner
			inner = outer + x
			outer = max(outer, _inner)
		}

		fmt.Println(max(inner, outer))
	}
}

func max(x, y int64) (z int64) {
	z = y
	if x > y {
		z = x
	}

	return
}
