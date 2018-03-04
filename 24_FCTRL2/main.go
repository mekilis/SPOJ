/* 
    Created by Smart Mek on 3/4/18 9:57 PM
*/

package main

import "fmt"

func main() {

	var t, n int
	fmt.Scan(&t)

	for tt := 0; tt < t; tt++ {

		fmt.Scanln(&n)
		factorial_large(n)
	}
}

func factorial_large(n int) {

	// create array with maximum number of digits fixed at 200
	res := make([]int, 200)

	res[0] = 1
	m := 1

	for x := 2; x <= n; x++ {

		m = multiply(x, res, m)
	}

	for i := m - 1; i >= 0; i-- {

		fmt.Print(res[i])
	}

	fmt.Println()
}

func multiply(x int, res []int, size_res int) int {

	carry := 0

	for i := 0; i < size_res; i++ {

		prod := res[i] * x + carry
		res[i] = prod % 10
		carry = prod / 10
	}

	for carry != 0 {

		res[size_res] = carry % 10
		carry /= 10
		size_res++
	}

	return size_res
}
