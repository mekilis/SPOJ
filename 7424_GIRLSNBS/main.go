/*
   Created by Smart Mek on May 18, 2018 10:42 AM
*/

package main

import "fmt"

func main() {
	var G, B int

	fmt.Scan(&G)
	fmt.Scan(&B)

	for G != -1 || B != -1 {

		var greater, lesser int

		if G > B {
			greater, lesser = G, B
		} else {
			greater, lesser = B, G
		}

		// code works well for G or B > 4
		if G == 0 || B == 0 {
			fmt.Println(G+B)
		} else if G == 1 || B == 1 {

			if greater % 2 != 0 {
				// odd
				greater++
			}

			fmt.Println(greater / 2)
		} else if G == B {
			fmt.Println(1)
		} else {
			ans := 2
			// custom mid point
			mid := greater / 2
			if greater % 2 == 0 {
				//even
				mid--
			}
			// answer up to and including mid == 2
			difference := mid - lesser
			if difference > 0 {
				ans += difference
			} else {
				// ans remains 2 except
				if greater - lesser == 1 {
					ans = 1
				}
			}

			fmt.Println(ans)
		}

		fmt.Scan(&G)
		fmt.Scan(&B)
	}
}
