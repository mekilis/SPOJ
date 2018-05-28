/*
   Created by Smart Mek on May 21, 2018 4:18 PM
*/

package main

import (
	"bufio"
	"fmt"
	"math"
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

	// Method 2: Set up tree
	c := math.Ceil(math.Log2(float64(N)))
	size := 2*int(math.Pow(2, c)) - 1

	tree = make([]Node, 0)
	for i := 0; i < size; i++ {
		tree = append(tree, NewNode())
	}

	constructTree(tree, list, 0, n-1, 1)

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
		// 1. Method 1
		//_, _, max := findMaximumSubarray(list, i-1, j-1)

		// 2. Method 2
		max := segmentTreeMax(tree, n, i-1, j-1)

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

// using another O(n log (n)) solution - segment trees
type Node struct {
	maxPrefixSum   int64
	maxSuffixSum   int64
	maxSubArraySum int64
	totalSum       int64
}

var mathMin = int64(math.MinInt16)

func NewNode() Node {
	n := Node{maxPrefixSum: mathMin, maxSuffixSum: mathMin, maxSubArraySum: mathMin, totalSum: 0}
	return n
}

var tree []Node

func segmentTreeMax(tree []Node, n, i, j int64) int64 {
	node := query(tree, 0, n-1, i, j, 1)

	// fmt.Printf("node: %+v\n", node)

	return node.maxSubArraySum
}

func merge(leftChild, rightChild Node) Node {

	var parentNode = NewNode()

	parentNode.maxPrefixSum = max(leftChild.maxPrefixSum, leftChild.totalSum+
		rightChild.maxPrefixSum)

	parentNode.maxSuffixSum = max(rightChild.maxSuffixSum, rightChild.totalSum+
		leftChild.maxSuffixSum)

	parentNode.totalSum = leftChild.totalSum + rightChild.totalSum

	parentNode.maxSubArraySum = max(max(leftChild.maxSubArraySum,
		rightChild.maxSubArraySum), leftChild.maxSuffixSum+rightChild.maxPrefixSum)

	return parentNode
}

func max(x int64, y int64) (z int64) {
	z = x
	if y > x {
		z = y
	}
	return
}

func constructTree(tree []Node, A []int64, start, end, i int64) {
	if start == end {
		tree[i].totalSum = A[start]
		tree[i].maxSuffixSum = A[start]
		tree[i].maxPrefixSum = A[start]
		tree[i].maxSubArraySum = A[start]

		return
	}

	// build children recursively
	mid := (start + end) / 2
	constructTree(tree, A, start, mid, 2*i)
	constructTree(tree, A, mid+1, end, 2*i+1)

	// merge
	tree[i] = merge(tree[2*i], tree[2*i+1])
}

func query(tree []Node, ss, se, qs, qe, i int64) Node {
	if ss > qe || se < qs {
		return NewNode()
	}

	if ss >= qs && se <= qe {
		return tree[i]
	}

	mid := (ss + se) / 2

	left, right := query(tree, ss, mid, qs, qe, 2*i),
		query(tree, mid+1, se, qs, qe, 2*i+1)

	return merge(left, right)
}

/*
// O(n log (n)) solution still timed out
func findMaximumSubarray(A []int64, low, high int64) (int64, int64, int64) {
	if high == low {
		return low, high, A[low]
	} else {

		var leftLow, leftHigh, leftSum int64
		var rightLow, righttHigh, rightSum int64
		var crossLow, crossHigh, crossSum int64

		mid := (low + high) / 2
		leftLow, leftHigh, leftSum = findMaximumSubarray(A, low, mid)
		rightLow, righttHigh, rightSum = findMaximumSubarray(A, mid+1, high)
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
	for j := mid + 1; j <= high; j++ {
		sum += A[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}
*/
