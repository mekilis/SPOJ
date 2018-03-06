/* 
    Created by Smart Mek on 3/6/18 2:39 PM
*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	var N, x, y, T int

	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		fmt.Scanln(&N)

		m := make([]int, 0)
		f := make([]int, 0)

		// get male stats
		for n := 0; n < N; n++ {

			fmt.Scan(&x)
			m = append(m, x)
		}

		// sort
		sort.Ints(m)

		// get female stats
		for n := 0; n < N; n++ {

			fmt.Scan(&y)
			f = append(f, y)
		}

		// sort again
		sort.Ints(f)

		// multiply and sum
		score := 0
		for i := 0; i < N; i++ {
			score += m[i] * f[i]
		}

		// print score
		fmt.Println(score)
	}

	// solved 0.12s 7.2M

}
