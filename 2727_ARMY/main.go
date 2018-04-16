/*
   Created by Smart Mek on Apr 04, 2018 3:29 PM
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var T, NG, NM, x int

	fmt.Scanln(&T)

	for t := 0; t < T; t++ {
		NGArr := make([]int, 0)
		NMArr := make([]int, 0)
		fmt.Scan()

		fmt.Scan(&NG)
		fmt.Scan(&NM)

		for ng := 0; ng < NG; ng++ {
			fmt.Scan(&x)
			NGArr = append(NGArr, x)
		}

		for nm := 0; nm < NM; nm++ {
			fmt.Scan(&x)
			NMArr = append(NMArr, x)
		}

		sort.Ints(NMArr)
		sort.Ints(NGArr)

		// battle begins
		if NMArr[NM - 1] > NGArr[NG - 1] {
			fmt.Println("MechaGodzilla")
		} else {
			fmt.Println("Godzilla")
		}
	}
}
