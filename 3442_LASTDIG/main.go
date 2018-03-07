/* 
    Created by Smart Mek on 3/6/18 5:18 PM
*/

package main

import (
	"math"
	"strconv"
	"strings"
	"log"
	"fmt"
)

/*
	----|---------------------------------------------------
		|	0	1	2	3	4	5	6	7	8	9	0	...
	2	|	1	2	4	8	6	2	4	8	6	2	4
	3	|	1	3	9	7	1	3	9	7	1	3	9
	4	|	1	4	6	4	6	4	6	4	6	4	6
	5	|	1	5	5	5	5	5	5	5	5	5	5
	6	|	1	6	6	6	6	6	6	6	6	6	6
	7	|	1	7	9	3	1	7	9	3	1	7	9
	8	|	1	8	4	2	6	8	4	2	6	8	4
	9	|	1	9	1	9	1	9	1	9	1	9	1
	----|--------------------------------------------------

*/

func main() {

	powerMatrix := make([][]int, 21)
	for i := 0.0; i <= 20.0; i++ {
		for j := 0.0; j <= 4.0; j++ {
			pow := math.Pow(i, j)
			powString := strings.Split(strconv.Itoa(int(pow)), "")
			lastDigit, err := strconv.Atoi(powString[len(powString) - 1])
			if err != nil {
				log.Fatal(err)
			}
			powerMatrix[int(i)] = append(powerMatrix[int(i)], lastDigit)
		}
	}

	//fmt.Println(powerMatrix)
	var t, a, b int

	fmt.Scan(&t)

	for tt := 0; tt < t; tt++ {
		fmt.Scan(&a)
		fmt.Scan(&b)

		key := b

		if key != 0 {
			key = b % 4
			if key == 0 {
				key = 4
			}
		}

		//fmt.Println(a, key)
		fmt.Println(powerMatrix[a][key])
	}

} // apparently solution was too long. might be the chart. solved in C++ 4.3.2 though in 0.0s with 2.8M used.
