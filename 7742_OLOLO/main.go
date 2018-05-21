/*
   Created by Smart Mek on May 21, 2018 2:20 PM
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, Pi, buffer int64

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		N = toInt(scanner.Bytes())
	}

	for n := int64(1); n <= N; n++ {
		scanner.Scan()
		Pi = toInt(scanner.Bytes())
		buffer ^= Pi
	}

	fmt.Println(buffer)
}

func toInt(buf []byte) int64 {
	var n int64

	for _, v := range buf {
		n = n*10 + int64(v-'0')
	}

	return n
}
// 7888
