/* 
    Created by Smart Mek on 3/4/18 9:57 PM
*/

package main

import "fmt"

func main() {

	factorial_large(100)
}

func factorial(n int64) int64 {

	if n == 1 {
		return n * 1
	}

	return n * factorial(n - 1)
}

func factorial_large(n int64) {

	// create array with maximum number of digits fixed at 200
	res := make([]int64, 200)

	res[0] = 1
	m := int64(1)

	for x := int64(2); x <= n; x++ {

		m = multiply(x, res, m)
	}

	for i := m - 1; i >= 0; i-- {

		fmt.Print(res[i])
	}

	fmt.Println()
}

func multiply(x int64, res []int64, size_res int64) int64 {

	carry := int64(0)

	for i := int64(0); i < size_res; i++ {

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
