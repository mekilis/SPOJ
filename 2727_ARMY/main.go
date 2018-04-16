/*
   Created by Smart Mek on Apr 04, 2018 3:29 PM
*/

package main

import "fmt"

func main() {
	var T, NG, NM, x int

	fmt.Scanln(&T)

	for t := 0; t < T; t++ {
		var NGMAx, NMMax int
		fmt.Scan()

		fmt.Scan(&NG)
		fmt.Scan(&NM)

		if NG == 0 && NM == 0 {
			fmt.Println("uncertain")
			continue
		}

		for ng := 0; ng < NG; ng++ {
			fmt.Scan(&x)
			if x > NGMAx {
				NGMAx = x
			}
		}

		for nm := 0; nm < NM; nm++ {
			fmt.Scan(&x)
			if x > NMMax {
				NMMax = x
			}
		}

		// battle begins
		if NMMax > NGMAx {
			fmt.Println("MechaGodzilla")
		} else {
			fmt.Println("Godzilla")
		}
	}
}
