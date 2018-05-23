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
	var b []byte
	b = scanner.Bytes()

	for len(b) == 0 {
		// next scan
		scanner.Scan()
		b = scanner.Bytes()
	}

	N = toInt(b)

	for n = 0; n < N; n++ {

		scanner.Scan()
		b = scanner.Bytes()
		for len(b) == 0 {
			// next scan
			scanner.Scan()
			b = scanner.Bytes()
		}

		x = toInt(b)
		list = append(list, x)
	}

	scanner.Scan()
	b = scanner.Bytes()
	for len(b) == 0 {
		// next scan
		scanner.Scan()
		b = scanner.Bytes()
	}

	M = toInt(b)

	for m = 0; m < M; m++ {
		scanner.Scan()
		b = scanner.Bytes()
		for len(b) == 0 {
			// next scan
			scanner.Scan()
			b = scanner.Bytes()
		}
		i = toInt(b)

		scanner.Scan()
		b = scanner.Bytes()
		for len(b) == 0 {
			// next scan
			scanner.Scan()
			b = scanner.Bytes()
		}
		j = toInt(b)

		//fmt.Println("M", i, j, list)
		_, _, max := findMaximumSubarray(list, i-1, j-1)
		fmt.Println(max)
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

func findMaximumSubarray(A []int64, low, high int64) (int64, int64, int64) {
	if high == low {
		return low, high, A[low]
	} else {

		var leftLow, leftHigh, leftSum int64
		var rightLow, righttHigh, rightSum int64
		var crossLow, crossHigh, crossSum int64

		mid := (low + high) / 2
		leftLow, leftHigh, leftSum = findMaximumSubarray(A, low, mid)
		rightLow, righttHigh, rightSum = findMaximumSubarray(A, mid + 1, high)
		crossLow, crossHigh, crossSum = findMaximumCrossingSubarray(A, low, mid, high)

		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, righttHigh, rightSum
		}

		return crossLow, crossHigh, crossSum
	}
}

func findMaximumCrossingSubarray(A []int64, low, mid, high int64) (int64, int64, int64) {

	var leftSum, rightSum, sum, maxLeft, maxRight int64

	leftSum = -1 << 63
	sum = 0
	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum = -1 << 63
	sum = 0
	for j := mid+1; j <= high; j++ {
		sum += A[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}