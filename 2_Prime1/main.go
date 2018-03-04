package main

import (
	"fmt"
)

func main() {

	var t, m, n int64

	// get no of test cases
	fmt.Scan(&t)

	for i := int64(0); i < t; i++ {

		fmt.Scan(&m)
		fmt.Scanln(&n)

		for j := m; j <= n; j++ {

			if isPrime(j) {

				fmt.Println(j)
			}
		}

		fmt.Println()
	}
}

func isPrime(n int64) bool {

	// Using AKS
	if n == 2 || n == 3 {

		return true
	}

	if n%2 == 0 || n%3 == 0 || n < 2 {

		return false
	}

	i, w := int64(5), int64(2)

	for i*i <= n {

		if n%i == 0 {

			return false
		}

		i += w
		w = 6 - w
	}

	return true
}
