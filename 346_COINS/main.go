/*
   Created by Smart Mek on 3/7/18 9:04 AM
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	// using less space from now
	var coin int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		coin,_  = strconv.Atoi(scanner.Text())

		exchange := coin/2 + coin/3 + coin/4
		fmt.Println("exchange =", exchange)
		if exchange > coin {
			coin = exchange
		}

		fmt.Println(int(coin))
	}
}
