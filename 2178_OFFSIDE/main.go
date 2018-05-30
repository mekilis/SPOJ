/*
   Created by Smart Mek on May 30, 2018 12:39 PM
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var A, B, C, D int

	var attackers = make([]int, 0)
	var defenders = make([]int, 0)

	var offside bool

	fmt.Scan(&A)
	fmt.Scan(&D)

	for A != 0 && D != 0 {

		offside = false

		for a := 0; a < A; a++ {
			fmt.Scan(&B)
			attackers = append(attackers, B)
		}

		for d := 0; d < D; d++ {
			fmt.Scan(&C)
			defenders = append(defenders, C)
		}

		sort.Ints(attackers)
		sort.Ints(defenders)

		if attackers[0] == defenders[0] {
			if attackers[0] != defenders[1] {
				offside = true
			}
		} else if attackers[0] < defenders[0] {
			offside = true
		}

		if offside {
			fmt.Println("Y")
		} else {
			fmt.Println("N")
		}

		fmt.Scan(&A)
		fmt.Scanln(&D)
	}
}

// WA