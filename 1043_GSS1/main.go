/*
   Created by Smart Mek on May 21, 2018 4:18 PM
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M, n, m, x, i, j int64
	var list = make([]int64, 0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	N = toInt(scanner.Bytes())

	for n = 0; n < N; n++ {
		scanner.Scan()
		x = toInt(scanner.Bytes())
		list = append(list, x)
	}

	scanner.Scan()
	M = toInt(scanner.Bytes())

	for m = 0; m < M; m++ {
		scanner.Scan()
		i = toInt(scanner.Bytes())

		scanner.Scan()
		j = toInt(scanner.Bytes())

		fmt.Println("M", i, j, list)
	}
}

func toInt(buf []byte) int64 {
	var n int64
	var negative bool

	if buf[0] == 45 {
		negative = true
		buf = buf[1:]
	}

	for _, v := range buf {
		n = n*10 + int64(v-'0')
	}

	if negative {
		n = -n
	}

	return n
}