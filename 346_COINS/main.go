/*
   Created by Smart Mek on 3/7/18 9:04 AM
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX_SIZE = 1000000000

var lookup = make([]int, MAX_SIZE+1)

func main() {
	//fmt.Println(lookup)
	// using less space from now
	var coin int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		coin, _ = strconv.Atoi(scanner.Text())
		exchange := compute(coin)
		//fmt.Println("exchange =", exchange)

		fmt.Println(exchange)
	}
}

func compute(coin int) int {
	// base case
	// max occurs from 12 upwards
	//fmt.Println("got", coin)
	if coin < 12 {
		//fmt.Println("returning", coin)
		return coin
	}
	if lookup[coin] != 0 {
		//fmt.Println("coin", coin, "has already been process to have", lookup[coin])
		return lookup[coin]
	}
	// compute again
	//fmt.Println("solving the first time for", coin, "with", coin / 2, coin / 3, coin /4)
	ex := compute(coin/2) + compute(coin/3) + compute(coin/4)
	max := coin
	//fmt.Println(ex, coin)
	if ex > coin {
		max = ex
	}
	lookup[coin] = max
	return lookup[coin]
}
