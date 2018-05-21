/*
   Created by Smart Mek on May 21, 2018 1:32 PM
*/

package main

import "fmt"

func main() {
	var N int64

	fmt.Scan(&N)

	for N != -1 {
		// series
		// Beehive Num = 1 + 6(n-1)

		n := N-N + 1

		for i := n; i < 1000000001; i += 6*(n - 1) {
			if i == N {
				fmt.Println("Y")
				break
			}

			if i > N {
				fmt.Println("N")
				break
			}

			n++
		}

		fmt.Scan(&N)
	}
}
