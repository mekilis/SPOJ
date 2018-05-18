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

		if G == 0 || B == 0 {
			fmt.Println(G+B)
		} else {

			ans := greater / (lesser + 1)
			if greater % (lesser + 1) > 0 {
				ans++
			}

			fmt.Println(ans)
		}

		fmt.Scan(&G)
		fmt.Scan(&B)
	}
}
