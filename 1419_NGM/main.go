/*
   Created by Smart Mek on May 21, 2018 12:11 PM
*/

package main

import "fmt"

func main() {
	var n uint64
	fmt.Scan(&n)

	move := n % 10
	if move == 0 {
		fmt.Println(2)
	} else {
		fmt.Println(1)
		fmt.Println(move)
	}
}
