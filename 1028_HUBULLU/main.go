/*
   Created by Smart Mek on Jul 12, 2018 11:41 AM
*/

package main

import "fmt"

func main() {
	var N, T, player int64
	fmt.Scanln(&T)
	for t := int64(0); t < T; t++ {
		fmt.Scan(&N)
		fmt.Scanln(&player)
		// if odd, 1 wins or else 0 wins
		if player == 0 {
			fmt.Println("Airborne wins.")
		} else {
			fmt.Println("Pagfloyd wins.")
		}
	}
}
