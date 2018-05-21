/*
   Created by Smart Mek on May 21, 2018 12:55 PM
*/

package main

import "fmt"

func main() {
	var N, T, cards, sum uint64

	fmt.Scan(&T)

	for t := T-T; t < T; t++ {
		fmt.Scan(&N)

		//cards = N % 1000007 // just in case

		// sum = (n/2)[a+l]; a = 1, l = cards - 1; a+l = N
		sum = (N-1)*(N) / 2

		cards = N * (N + 1) + sum
		fmt.Println(cards % 1000007)
	}
}

/*
1 - 2, 0
2 - 6, 1
3 - 12, 3(2, 1)
7 - 56, 21(6, 5, 4, 3, 2, 1)
*/
