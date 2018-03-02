package main

import "fmt"

func main() {

	var t, m, n int64

	// get no of test cases
	fmt.Scan(&t)

	for i := int64(0); i < t; i++ {

		fmt.Scan(&m)
		fmt.Scanln(&n)

		//fmt.Println(m, n)

		if conditionSatisfied(m, n) {

			primesInRange(m, n)
		}
		fmt.Println()
	}
}

func primesInRange(m int64, n int64) {

	for i := m; i <= n; i++ {

		if i == 2 || i == 3 {

			fmt.Println(i)
			continue
		}

		if i == 1 || i % 2 == 0 {
			continue
		}

		isPrime := true

		for j := int64(2); j < i; j++ {

			if i % j == 0 {

				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Println(i)
		}

	}
}

func conditionSatisfied(m, n int64) bool {

	if m > n {
		return false
	}

	if n > 1000000000 {
		return false
	}

	if n - m > 100000 {
		return false
	}

	return true
}
