/* 
    Created by Smart Mek on 3/6/18 12:07 PM
*/

package main

import "fmt"

func main() {

	var N int

	// get first N
	fmt.Scanln(&N)
	for N != 0 {

		// y(N) = (2n+1) * (n+1 * n)/6)
		fmt.Println(((2*N) + 1) * ((N + 1) * N) / 6)

		// fetch next N
		fmt.Scanln(&N)
	}
}
