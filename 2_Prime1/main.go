package main

import "fmt"

func main() {

	var t, m, n int64

	// get no of test cases
	fmt.Scan(&t)

	for i := int64(0); i < t; i++ {

		fmt.Scan(&m)
		fmt.Scanln(&n)

		primesInRange(m, n)

		fmt.Println()
	}
}

func primesInRange(m int64, n int64) {

	// using the sieve of eratosthenes

	//create map of consecutive integers from 2 to n setting marked to false (by default)
	nonPrimes := make(map[int64]bool, n)

	// set p to the first prime number i.e. 2
	for p := int64(2); p * p <= n; p++ {

		// mark off all numbers while counting up in increments of p
		// check if marked is false
		if !nonPrimes[p] {

			// this number hasn't been marked, update the multiples starting from p * 2
			for j := p * 2; j <= n; j += p {

				nonPrimes[j] = true
				// multiple has been marked
			}
		}
	}

	for num := m; num <= n; num++ {

		if num < 2 || nonPrimes[num]{
			continue
		}
		fmt.Println(num)
	}

}
