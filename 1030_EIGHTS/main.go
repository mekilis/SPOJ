/* 
    Created by Smart Mek on 3/14/18 10:12 AM
*/

package main

import "fmt"

func main() {
/*	var cube, i int64

	// generate all patterns ending in 8
	for j := int64(2); j < 2000; j += 10 {
		cube = j * j * j
		i++

		if i < 20 {
			continue
		}

		fmt.Printf("%d. %d, j = %d\n", i, cube, j)
	}
	// j = 192, 442, 692, 942.... Differences: 250 250 250...
	// k = 1, j = 192 + 0(250); k = 2, k = 192 + 250 etc
	// j = 192 + (k - 1) * 250*/

	var t, k int64

	fmt.Scanln(&t)

	for tt := int64(0); tt < t; tt++ {
		fmt.Scanln(&k)
		fmt.Println(192 + (k - 1) * 250)
	}
}
