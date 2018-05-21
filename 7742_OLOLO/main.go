/*
   Created by Smart Mek on May 21, 2018 2:20 PM
*/

package main

import "fmt"

func main() {
	var N, Pi, max int64
	var lookup []int64

	fmt.Scan(&N)
	lookup = make([]int64, 1000000001)
	//fmt.Println("created lookup", lookup)

	for n := int64(1); n <= N; n++ {
		fmt.Scan(&Pi)
		lookup[Pi]++

		if Pi > max {
			max = Pi
		}
	}

	//fmt.Println(	"done scanning", lookup)

	for n := int64(1); n <= max; n++ {
		//fmt.Println("scanning")
		if x := lookup[n]; x == 1 {
			fmt.Println(n)
			break
		}
	}
}
