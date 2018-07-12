/*
   Created by Smart Mek on Jul 02, 2018 10:39 AM
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	var N int64
	var x int64
	var zero = int64(0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	T = int(toInt(scanner.Bytes()))

	for t := 0; t < T; t++ {

		scanner.Scan()
		N = toInt(scanner.Bytes())

		if N == 0 {
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanWords)
			fmt.Printf("Case %d: 0\n", t+1)
			continue
		}

		scanner.Scan()
		x = toInt(scanner.Bytes())
		inner, outer, _inner := x, zero, zero

		for i := zero; i < N-1; i++ {
			scanner.Scan()
			x = toInt(scanner.Bytes())
			_inner = inner
			inner = outer + x
			outer = max(outer, _inner)
		}

		fmt.Printf("Case %d: %d\n", t+1, max(inner, outer))
	}
}

func max(x, y int64) (z int64) {
	z = y
	if x > y {
		z = x
	}

	return
}

func toInt(buf []byte) int64 {
	var n int64

	for _, v := range buf {
		n = n*10 + int64(v-'0')
	}

	return n
}

// TLE even with this...