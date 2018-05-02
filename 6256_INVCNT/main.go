/*
   Created by Smart Mek on Apr 27, 2018 11:30 PM
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var T, x uint64
	var N int

	fmt.Scanln(&T)

	for t := uint64(0); t < T; t++ {
		fmt.Scanln()

		fmt.Scanln(&N)

		A := make([]uint64, 0)

		for n := 0; n < N; n++ {
			fmt.Scanln(&x)
			A = append(A, x)
		}

		counter := mergeSort(A, 0, N-1)
		fmt.Println(counter)
	}
}

func mergeSort(A []uint64, p, r int) uint64 {
	var counter uint64 = 0

	if p >= r {
		return counter
	}

	q := (p + r) / 2
	counter = mergeSort(A, p, q)
	counter += mergeSort(A, q + 1, r)
	counter += merge(A, p, q, r)

	return counter
}

func merge(A []uint64, p, q, r int) uint64 {
	n1, n2 := q - p + 1, r - q
	B := make([]uint64, n1+1)
	C := make([]uint64, n2+1)
	var counter uint64

	for i := 0; i < n1; i++ {
		B[i] = A[p + i]
	}

	for i := 0; i < n2; i++ {
		C[i] = A[q + i + 1]
	}

	inf := uint64(math.MaxUint64)
	B[n1] = inf
	C[n2] = inf

	var i, j int

	for k := p; k <= r; k++ {
		if B[i] <= C[j] {
			A[k] = B[i]
			i++
		} else {
			A[k] = C[j]
			j++

			counter += uint64(n1 - i)
		}
	}

	return counter
}
