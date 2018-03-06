/* 
    Created by Smart Mek on 3/5/18 11:48 AM
*/

package main

import "fmt"

func main() {

	var N, x, y int
	fmt.Scan(&N)

	for n := 0; n < N; n++ {

		fmt.Scan(&x)
		fmt.Scanln(&y)

		// if x and y are both even
		if x % 2 == 0 && y % 2 == 0 {

			// if x == y
			if x == y || x == y + 2 {

				fmt.Println(x + y)

			} else {

				fmt.Println("No Number")

			}
		} else if x % 2 != 0 && y % 2 != 0 {

			// both are odd
			if x == y || x == y + 2 {

				fmt.Println(x + y - 1)

			} else {

				fmt.Println("No Number")

			}
		} else {

			fmt.Println("No Number")

		}
	}
}
