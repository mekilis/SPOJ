/* 
    Created by Smart Mek on 3/6/18 3:01 PM
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	var cols int
	var letters string

	fmt.Scan(&cols)

	for cols != 0 {

		// assuming no spaces
		fmt.Scanln(&letters)
		matrix := strings.Split(letters, "")

		//fmt.Println(matrix)
		// using non-zero based row count
		rowCount := 1

		for i := 0; i < len(matrix); i += cols {

			lettersInRange := matrix[i:(rowCount * cols)]

			// if it's not 1st, 3rd, 5th etc row, i.e. an even row based on rowCount, reverse that range
			if rowCount% 2 == 0 {

				// reverse this one first
				for letter := len(lettersInRange)/2 - 1; letter >= 0; letter-- {

					// fetch opposite the swap
					opp := len(lettersInRange) - letter - 1
					//swap
					lettersInRange[letter], lettersInRange[opp] = lettersInRange[opp], lettersInRange[letter]
				}

			}
			rowCount++
		}

		// print out words by grouping columns together
		for c := 0; c < cols; c++ {

			// print using index c to print others in steps of c
			for col := c; col < len(matrix); col += cols {

				fmt.Print(matrix[col])
			}
		}

		fmt.Println()

		fmt.Scanln(&cols)
	}

	//0.00s 3.1M

}
