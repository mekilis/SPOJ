/*
   Created by Smart Mek on Apr 16, 2018 1:00 PM
*/

package main

import (
	"fmt"
	"sort"
)

// would define sort.Interface on this
type MyRange []uint64

func (m MyRange) Len() int           { return len(m) }
func (m MyRange) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MyRange) Less(i, j int) bool { return m[i] < m[j] }

func main() {

	var T uint64
	var N, C int

	fmt.Scan(&T)

	for t := uint64(0); t < T; t++ {

		fmt.Scan(&N)
		fmt.Scan(&C)

		stalls := make(MyRange, 0)
		var x, sum uint64
		for n := 0; n < N; n++ {
			fmt.Scan(&x)
			stalls = append(stalls, x)
			sum += x
		}

		// sort array first
		sort.Sort(stalls)

		x = stalls[N-1] - stalls[0]

		fmt.Println(binarySearch(stalls, N, C, x))
	}
}

func binarySearch(stalls MyRange, N, C int, x uint64) uint64 {
	var p, q, r uint64

	p, r = 1, x

	for p < r {
		q = (p + r) / 2

		cowsPlaced, current := 1, stalls[0]

		for i := 1; i < N; i++ {
			if cowPos := stalls[i] - current; cowPos < q {
				//fmt.Println("not possible to place. current:", current, "next sum:", cowPos)
			} else {
				cowsPlaced++
				current = stalls[i]
				//fmt.Println("cow can be placed. next:", current, "cowsPlaced:", cowsPlaced)

				if cowsPlaced == C {
					break
				}
			}
		}

		if cowsPlaced < C {
			r = q
		} else {
			p = q
		}
	}

	return p
}
