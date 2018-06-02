/*
   Created by Smart Mek on 6/2/2018, 5:01:13 PM
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// start := time.Now()
	// var N, counter, sum int64
	var ans []string
	// fmt.Scan(&N)
	// for N = 2147483614; N < 2147483647; N++ {
	// 	fmt.Println("=========== N =", N, "=========")
	// 	counter = 0
	// 	sum = N
	// Loop:
	// 	//sum = 0
	// 	ans = strings.Split(strconv.Itoa(int(sum)), "")
	// 	sum = 0
	//
	// 	for i := 0; i < len(ans); i++ {
	// 		x, _ := strconv.Atoi(ans[i])
	// 		j := x * x
	// 		sum += int64(j)
	// 	}
	// 	fmt.Print(" -->", sum)
	//
	// 	for sum != 1 && counter < 20 {
	// 		//ans = strings.Split(strconv.Itoa(sum), "")
	// 		if sum == 4 || sum == 145 {
	// 			break
	// 		}
	// 		counter++
	// 		goto Loop
	// 	}
	//
	// 	fmt.Println()
	// }
	// fmt.Println(time.Since(start))
	var N int
	var sum, counter int
	var Happy = true

	fmt.Scan(&N)
	sum = N
Loop:
	ans = strings.Split(strconv.Itoa(int(sum)), "")
	sum = 0
	for i := 0; i < len(ans); i++ {
		x, _ := strconv.Atoi(ans[i])
		j := x * x
		sum += j
	}

	counter++

	for sum != 1 {
		if sum == 4 || sum == 145 {
			Happy = false
			break
		}
		goto Loop
	}

	if Happy {
		fmt.Println(counter)
	} else {
		fmt.Println(-1)
	}
}
