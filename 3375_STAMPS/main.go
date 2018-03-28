/*
   Created by Smart Mek on 3/28/18 1:05 PM
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Everyone hates Raymond
	var Scenarios, stamps, friends, min, x int

	fmt.Scan(&Scenarios)

	var f []int

	for scenario := 0; scenario < Scenarios; scenario++ {
		fmt.Scan(&stamps)
		fmt.Scanln(&friends)

		//fmt.Println(stamps, friends)
		f = make([]int, 0)
		min = 0

		for friend := 0; friend < friends; friend++ {
			fmt.Scan(&x)
			f = append(f, x)
			//fmt.Println("done appending x", x, f)
		}

		sort.Ints(f)
		//fmt.Println(f)

		fmt.Printf("Scenario #%d:\n", scenario+1)
		sum := 0
		for i := friends - 1; i >= 0; i-- {
			min++
			sum += f[i]
			if sum >= stamps {
				fmt.Println(min)
				break
			}
		}

		if sum < stamps {
			fmt.Println("impossible")
		}

		if scenario < Scenarios - 1 {
			fmt.Println()
		}
	}
}
/*
150 6
45 12 32 65 48 25
Scenario #1:
3
500 20
65 48 52 19 32 25 65 74 85 12 36 25 99 150 42 46 4 6 7 0
Scenario #2:
6
1 5
45 18 3 2 6
Scenario #3:
1
999 5
125 451 323 169 5
Scenario #4:
4
*/
