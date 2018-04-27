/*
   Created by Smart Mek on Apr 27, 2018 5:54 PM
*/

package main

import "fmt"

func main() {
	var T, n int

	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		fmt.Scan(&n)
		f := n * (n + 2) * (2 * n + 1)
		if n % 2 != 0 {
			f--
		}
		f /= 8
		fmt.Println(f)
	}
}
