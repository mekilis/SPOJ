/* 
    Created by Smart Mek on 3/6/18 1:42 PM
*/

package main

import "fmt"

func main() {

	var candy, N int

	fmt.Scan(&N)

	for N != -1 {

		candies := make([]int, 0)

		sum := 0
		for n := 0; n < N; n++ {

			fmt.Scanln(&candy)
			sum += candy
			candies	= append(candies, candy)
		}

		if sum % N == 0 {
			mean := sum / N
			//fmt.Println("possible with", mean)
			count := 0

			// O(n^3)
			for _, v := range candies {
				if v < mean {
					count += (mean - v)
				}
			}

			fmt.Println(count)

		} else {
			//fmt.Println("not possible with,", float64(sum) / float64(N))
			fmt.Println(-1)
		}

		fmt.Scanln(&N)
	}

	// solved in 0.01s 4.0M
}
