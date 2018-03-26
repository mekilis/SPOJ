/*
   Created by Smart Mek on 3/26/18 11:53 AM
*/

package main

import "fmt"

func main() {
	c := 0.00

	fmt.Scanln(&c)

	for c != 0.0 {

		n := 0.0
		sumC := 0.0

		for sumC < c {
			n += 1.0
			sumC += 1 / (n + 1)
		}

		fmt.Println(n, "card(s)")
		fmt.Scanln(&c)
	}

}
